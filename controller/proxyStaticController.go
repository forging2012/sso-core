package controller

import (
	"encoding/json"

	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service/impl"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/i18n"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ProxyStaticController struct {
	beego.Controller
}

var proxyStaticSrv = impl.NewSsoProxyStaticRouteService()

func (this *ProxyStaticController) Get() {
	rst, err := proxyStaticSrv.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

func (this *ProxyStaticController) Post() {
	remoteSysApi, err := this.parseHttpReq()
	if err != nil {
		return
	}
	err = proxyStaticSrv.Post(remoteSysApi)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxyStaticController) Delete() {
	this.Ctx.Request.ParseForm()
	jsStr := this.Ctx.Request.FormValue("JSON")
	var ssoProxyStatic []entity.SsoProxyStaticRoute
	err := json.Unmarshal([]byte(jsStr), &ssoProxyStatic)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "请求参数格式不正确，请联系管理员")
		return
	}
	err = proxyStaticSrv.Delete(ssoProxyStatic)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxyStaticController) Put() {
	ssoProxyStatic, err := this.parseHttpReq()
	if err != nil {
		return
	}
	err = proxyStaticSrv.Put(ssoProxyStatic)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxyStaticController) parseHttpReq() (entity.SsoProxyStaticRoute, error) {
	this.Ctx.Request.ParseForm()
	form := this.Ctx.Request.Form
	jclaim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 403, i18n.Disconnect(this.Ctx.Request))
		return entity.SsoProxyStaticRoute{}, err
	}

	remoteSysApi := entity.SsoProxyStaticRoute{
		ServiceCd:   form.Get("serviceCd"),
		RegisterUrl: form.Get("registerUrl"),
		RouteDesc:   form.Get("routeDesc"),
		RemoteUrl:   form.Get("remoteUrl"),
		CreateUser:  jclaim.UserId,
		ModifyUser:  jclaim.UserId,
		Uuid:        form.Get("uuid"),
	}
	return remoteSysApi, nil
}
