package controller

import (
	"encoding/json"

	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/i18n"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service/impl"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ProxySubsystemApiController struct {
	beego.Controller
}

var ssoSubsystemApi = impl.NewSsoRouteRegisterApiService()

func (this *ProxySubsystemApiController) Get() {
	rst, err := ssoSubsystemApi.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

func (this *ProxySubsystemApiController) Post() {
	remoteSysApi, err := this.parseHttpReq()
	if err != nil {
		return
	}
	err = ssoSubsystemApi.Post(remoteSysApi)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxySubsystemApiController) Put() {
	remoteSysApi, err := this.parseHttpReq()
	if err != nil {
		return
	}
	err = ssoSubsystemApi.Put(remoteSysApi)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxySubsystemApiController) Delete() {
	this.Ctx.Request.ParseForm()
	jsStr := this.Ctx.Request.FormValue("JSON")
	var subsystemApi []entity.SsoRouteRegisterApi
	err := json.Unmarshal([]byte(jsStr), &subsystemApi)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "请求参数格式不正确，请联系管理员")
		return
	}
	err = ssoSubsystemApi.Delete(subsystemApi)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *ProxySubsystemApiController) parseHttpReq() (entity.SsoRouteRegisterApi, error) {
	this.Ctx.Request.ParseForm()
	form := this.Ctx.Request.Form
	jclaim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 403, i18n.Disconnect(this.Ctx.Request))
		return entity.SsoRouteRegisterApi{}, err
	}

	remoteSysApi := entity.SsoRouteRegisterApi{
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
