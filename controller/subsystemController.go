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

type SubsystemController struct {
	beego.Controller
}

var subsystemService = impl.NewSsoSubsystemService()

func (this *SubsystemController) Get() {
	rst, err := subsystemService.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

func (this *SubsystemController) Post() {
	subsystemEntity, err := this.parseHttpRequest()
	if err != nil {
		return
	}
	err = subsystemService.Post(subsystemEntity)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SubsystemController) Put() {
	subsystemEntity, err := this.parseHttpRequest()
	if err != nil {
		return
	}
	err = subsystemService.Put(subsystemEntity)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SubsystemController) Delete() {
	this.Ctx.Request.ParseForm()
	jsStr := this.Ctx.Request.FormValue("JSON")
	var args []entity.SsoSubsystemEntity
	err := json.Unmarshal([]byte(jsStr), &args)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "参数格式不正确，请联系管理员")
		return
	}
	err = subsystemService.Delete(args)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 422, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SubsystemController) parseHttpRequest() (entity.SsoSubsystemEntity, error) {
	this.Ctx.Request.ParseForm()
	form := this.Ctx.Request.Form
	var subsystemEntity entity.SsoSubsystemEntity
	subsystemEntity.ServiceCd = form.Get("serviceCd")
	subsystemEntity.ServiceName = form.Get("serviceName")
	subsystemEntity.RemotePort = form.Get("remotePort")
	subsystemEntity.RemoteScheme = form.Get("remoteScheme")
	subsystemEntity.RemoteHost = form.Get("remoteHost")
	subsystemEntity.PrefixUrl = form.Get("prefixUrl")
	jclaim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 403, i18n.Disconnect(this.Ctx.Request))
		return subsystemEntity, err
	}
	subsystemEntity.CreateUser = jclaim.UserId
	subsystemEntity.ModifyUser = jclaim.UserId
	return subsystemEntity, nil
}
