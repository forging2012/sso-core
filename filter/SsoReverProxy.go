package filter

import (
	"crypto/tls"
	"html/template"
	"net/http"
	"net/http/httputil"
	"path"

	"github.com/asofdate/sso-core/service/impl"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var ssoRouteService = impl.NewSsoRouteService()

func SsoReverProxy(ctx *context.Context) {
	// 未匹配路由信息，交给反向代理去请求子系统路由
	if _, yes := beego.BeeApp.Handlers.FindRouter(ctx); !yes {
		ctx.Request.ParseForm()
		// 反向代理可以接受一个参数，如果serviceCd为空，则忽略此参数
		// 如果注册了多个相同的路由，则取查询到的第一个子系统路由信息
		serviceCd := ctx.Request.FormValue("serviceCd")

		ssoEntity, err := ssoRouteService.Get(ctx.Request.URL.Path, serviceCd)
		if err != nil || len(ssoEntity.RemoteUrl) == 0 {
			ssoEntity, err = ssoRouteService.GetProxyStatic(ctx.Request.URL.Path, serviceCd)
			if err != nil || len(ssoEntity.RemoteUrl) == 0 {
				logger.Error("没有被注册的路由", ctx.Request.URL)
				hret.Error(ctx.ResponseWriter, 404, "没有被注册的路由")
				return
			}
		}

		director := func(req *http.Request) {
			req = ctx.Request
			req.URL.Path = path.Join(ssoEntity.PrefixUrl, ssoEntity.RemoteUrl)
			req.URL.Scheme = ssoEntity.RemoteScheme
			req.URL.Host = ssoEntity.RemoteHost + ":" + ssoEntity.RemotePort
			logger.Debug("proxy route is:", req.URL)
		}

		proxy := &httputil.ReverseProxy{
			Director: director,
			Transport: &http.Transport{
				TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
				DisableCompression: true,
			},
			ModifyResponse: func(response *http.Response) error {
				location, err := response.Location()
				if err == nil {
					// TODO
					// 反向代理路发生重定向，
					// 后台暂时没有处理重定向请求，需要前端自行处理
					logger.Debug("redirect", location, err)
				}
				return nil
			}}

		proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)
		// 匹配成功，退出beego路由处理程序
		ctx.ResponseWriter.Started = true
	}

	// 匹配系统默认过滤路由
	if FilterMatchRoute(ctx.Request.URL.Path) {
		return
	}

	// 系统内部路由连接校验
	if !jwt.CheckToken(ctx.Request) {
		hz, _ := template.ParseFiles("./views/sso/disconnect.tpl")
		hz.Execute(ctx.ResponseWriter, nil)
	}
}
