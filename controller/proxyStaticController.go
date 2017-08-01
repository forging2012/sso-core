package controller

import (
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-core/service/impl"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ProxyStaticController struct {
	beego.Controller
}

var proxyStaticService = impl.NewSsoProxyStaticRouteService()

func (this *ProxyStaticController) Get() {
	rst, err := proxyStaticService.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}
