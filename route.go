package sso_core

import (
	"github.com/asofdate/sso-core/controller"
	"github.com/asofdate/sso-core/filter"
	"github.com/asofdate/sso-core/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func Register() {
	// 请求授权
	beego.Any("/v1/sso/auth", controller.AuthorizationCtl.Auth)
	beego.Any("/v1/sso/identify", controller.AuthorizationCtl.Identify)
	beego.Any("/v1/sso/logout", controller.AuthorizationCtl.Logout)
	beego.Get("/v1/sso", controller.AuthorizationCtl.AuthPage)
	beego.Get("/", controller.AuthorizationCtl.AuthPage)

	beego.Get("/v1/sso/local/static/page", controller.StaticPageObj.GetLocalStaticResource)
	beego.Get("/v1/sso/subsystem/page", controller.StaticPageObj.GetSubsystemConfigPage)
	beego.Get("/v1/sso/subsystem/api/page", controller.StaticPageObj.GetRemoteApiPage)
	beego.Get("/v1/sso/proxy/static/page", controller.StaticPageObj.GetProxyStaticPage)

	beego.Router("/v1/sso/local/static", &controller.LocalStaticController{})
	beego.Router("/v1/sso/subsystem", &controller.SubsystemController{})
	beego.Router("/v1/sso/proxy/static", &controller.ProxyStaticController{})
	beego.Router("/v1/sso/proxy/api", &controller.ProxySubsystemApiController{})

	filter.AddMatchRoute("/", true)
	filter.AddMatchRoute("/v1/sso", true)
	filter.AddMatchRoute("/v1/sso/auth", true)
	filter.AddMatchRoute("/v1/sso/identify", true)

	// 读取本地静态资源配置表，注册静态资源
	utils.RegisterStaticDir()

	// 拦截未匹配的路由
	beego.InsertFilter("/*", beego.AfterStatic, func(ctx *context.Context) {
		filter.SsoReverProxy(ctx)
	}, false)
}
