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

type LocalStaticController struct {
	beego.Controller
}

var localStaticSrv = impl.NewStaticResourceService()

func (this *LocalStaticController) Get() {
	rst, err := localStaticSrv.GetList()
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "查询本地静态资源失败")
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

func (this *LocalStaticController) Post() {
	this.Ctx.Request.ParseForm()
	form := this.Ctx.Request.Form
	url := form.Get("url")
	path := form.Get("path")
	jclaim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 403, "获取用户连接信息失败")
		return
	}
	data := entity.LocalStaticEntity{
		Url:        url,
		Path:       path,
		CreateUser: jclaim.UserId,
		ModifyUser: jclaim.UserId,
	}
	err = localStaticSrv.Post(data)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *LocalStaticController) Put() {
	this.Ctx.Request.ParseForm()
	form := this.Ctx.Request.Form
	url := form.Get("url")
	path := form.Get("path")
	uuid := form.Get("uuid")
	jclaim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 403, "获取用户连接信息失败")
		return
	}
	data := entity.LocalStaticEntity{
		Url:        url,
		Path:       path,
		CreateUser: jclaim.UserId,
		ModifyUser: jclaim.UserId,
		Uuid:       uuid,
	}
	err = localStaticSrv.Put(data)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *LocalStaticController) Delete() {
	this.Ctx.Request.ParseForm()
	jsStr := this.Ctx.Request.FormValue("JSON")
	var args []entity.LocalStaticEntity
	err := json.Unmarshal([]byte(jsStr), &args)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 432, err.Error())
		return
	}
	err = localStaticSrv.Delete(args)
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}
