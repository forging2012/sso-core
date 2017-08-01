package utils

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"

	"fmt"

	"github.com/asofdate/sso-jwt-auth/models"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/i18n"
	"github.com/asofdate/sso-core/service/impl"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var resourceModel = new(models.ResourceModel)
var ssoRouteService = impl.NewSsoRouteService()

// 代理子系统页面
func SysIndexReverProxy(ctx *context.Context) {

	// 从用户请求中获取资源id信息
	id := ctx.Request.FormValue("Id")

	// 系统中找不到这个路由信息，使用反向代理，查询路由对应的内部系统
	serviceCd, err := resourceModel.GetServiceCd(id)
	if err != nil || len(serviceCd) == 0 {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 404, "查询菜单资源属性失败")
		return
	}

	reqUrl := ctx.Request.URL.Path

	ssoEntity, err := ssoRouteService.Get(reqUrl, serviceCd)
	if err != nil || len(ssoEntity.RemoteUrl) == 0 {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}

	director := func(req *http.Request) {
		req = ctx.Request
		req.URL.Path = ssoEntity.RemoteUrl
		req.URL.Scheme = ssoEntity.RemoteScheme
		req.URL.Host = ssoEntity.RemoteHost + ":" + ssoEntity.RemotePort
	}

	proxy := &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: func(response *http.Response) error {
			// 查看系统是否被重定向，
			// 如果系统被重定向，则继续获取重定向的内容
			location, err := response.Location()
			if err == nil {
				// TODO
				// 重定向追踪
				fmt.Println("重定向，location:", location, "err is:", err)
			}

			// 修改html中内容，允许iframe中打开页面。
			return nil
		},
	}

	if ssoEntity.RemoteScheme == "https" {
		proxy.Transport = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		}
	}

	proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return
}

// 代理子系统菜单构建过程
// 按照sso生成菜单的过程，生成子系统菜单
func SsoSubsystemMenuReverProxy(ctx *context.Context) {
	Id := ctx.Request.FormValue("Id")
	serviceCd, err := resourceModel.GetServiceCd(Id)
	if err != nil || len(serviceCd) == 0 {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 404, "查询菜单资源所属系统失败")
		return
	}
	reqUrl := ctx.Request.URL.Path

	// 外部系统，启用反向代理，获取信息
	ssoEntity, err := ssoRouteService.Get(reqUrl, serviceCd)

	if err != nil || len(ssoEntity.RemoteUrl) == 0 {
		logs.Error("没有查找到路由信息：", reqUrl)
		hret.Error(ctx.ResponseWriter, 404, "API没有注册")
		return
	}

	director := func(req *http.Request) {
		req = ctx.Request
		req.URL.Path = ssoEntity.RemoteUrl
		req.URL.Scheme = ssoEntity.RemoteScheme
		req.URL.Host = ssoEntity.RemoteHost + ":" + ssoEntity.RemotePort
	}

	proxy := &httputil.ReverseProxy{Director: director}
	if ssoEntity.RemoteScheme == "https" {
		proxy.Transport = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		}
	}

	proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	ctx.ResponseWriter.Started = true
	return
}
