package controller

import (
	"net/http"

	"strconv"

	"encoding/json"

	"io/ioutil"

	"github.com/asofdate/sso-jwt-auth/hrpc"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/asofdate/sso-core/dto"
	"github.com/asofdate/sso-core/service"
	"github.com/asofdate/sso-core/service/impl"
	"github.com/astaxie/beego/context"
)

type AuthorizationController struct {
	authSrv service.AuthService
}

var AuthorizationCtl = &AuthorizationController{
	authSrv: impl.NewAuthService(),
}

func (this *AuthorizationController) AuthPage(ctx *context.Context) {
	fd, _ := ioutil.ReadFile("./views/sso/login.tpl")
	ctx.ResponseWriter.Write(fd)
}

func (this *AuthorizationController) Logout(ctx *context.Context) {
	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		rdto := dto.AuthDto{
			Username: "no connected",
			RetCode:  434,
			RetMsg:   "用户没有登录系统",
		}
		this.result(ctx.ResponseWriter, rdto)
		return
	}
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(ctx.ResponseWriter, &cookie)
	logoutdto := dto.AuthDto{
		Username: jclaim.UserId,
		RetCode:  200,
		RetMsg:   "安全退出系统",
	}
	this.result(ctx.ResponseWriter, logoutdto)
	return
}

func (this *AuthorizationController) Identify(ctx *context.Context) {
	cookie, err := ctx.Request.Cookie("Authorization")
	if err != nil {
		token := ctx.Request.Header.Get("Authorization")
		if validator.IsEmpty(token) || !jwt.CheckToken(token) {
			ctx.Redirect(http.StatusTemporaryRedirect, "/v1/sso")
			return
		}
		this.result(ctx.ResponseWriter, dto.AuthDto{
			Username: "",
			RetCode:  200,
			RetMsg:   "success",
		})
		return
	}

	if jwt.CheckToken(cookie.Value) {
		this.result(ctx.ResponseWriter, dto.AuthDto{
			Username: "",
			RetCode:  200,
			RetMsg:   "success",
		})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/v1/sso")
	return
}

func (this *AuthorizationController) Auth(ctx *context.Context) {
	ctx.Request.ParseForm()
	form := ctx.Request.Form
	username := form.Get("username")
	password := form.Get("password")
	targetSystem := form.Get("targetSystem")
	effectiveTime := form.Get("effectiveTime")

	if validator.IsEmpty(username) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  422,
			RetMsg:   "鉴权失败，账号为空",
		}
		this.result(ctx.ResponseWriter, rdto)
		return
	}

	if validator.IsEmpty(password) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  423,
			RetMsg:   "鉴权失败，密码为空",
		}
		this.result(ctx.ResponseWriter, rdto)
		return
	}

	if validator.IsEmpty(targetSystem) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  424,
			RetMsg:   "鉴权失败，目标系统为空",
		}
		this.result(ctx.ResponseWriter, rdto)
		return
	}

	retDto := this.authSrv.Auth(dto.AuthDto{
		Username:      username,
		Password:      password,
		TargetSystem:  targetSystem,
		EffectiveTime: effectiveTime,
		RetCode:       0,
		RetMsg:        "",
	})

	if retDto.RetCode == 200 {
		domainId, err := hrpc.GetDomainId(username)
		if err != nil {
			logger.Error(username, " 用户没有指定的域", err)
			retDto.RetCode = 426
			retDto.RetMsg = "获取用户域信息失败"
			this.result(ctx.ResponseWriter, retDto)
			return
		}

		orgid, err := hrpc.GetOrgUnitId(username)
		if err != nil {
			logger.Error(username, " 用户没有指定机构", err)
			retDto.RetCode = 427
			retDto.RetMsg = "获取用户所在机构失败"
			this.result(ctx.ResponseWriter, retDto)
			return
		}

		et, err := strconv.ParseInt(effectiveTime, 10, 64)
		if err != nil || validator.IsEmpty(effectiveTime) {
			et = 17280
		}

		token := jwt.GenToken(username, domainId, orgid, et)
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: int(et)}
		http.SetCookie(ctx.ResponseWriter, &cookie)
		retDto.RetMsg = token
		this.result(ctx.ResponseWriter, retDto)
		return
	}
	this.result(ctx.ResponseWriter, retDto)
	return
}

func (this *AuthorizationController) result(respone http.ResponseWriter, cdto dto.AuthDto) {
	token, err := json.Marshal(cdto)
	if err != nil {
		respone.WriteHeader(http.StatusExpectationFailed)
		respone.Write([]byte(`{username:` + cdto.Username + `,RetCode:"431",retMsg:"format json type info failed."}`))
		return
	}
	respone.WriteHeader(cdto.RetCode)
	respone.Write([]byte(token))
}
