package controller

import (
	"io/ioutil"

	"github.com/astaxie/beego/context"
)

type StaticPage struct {
}

var StaticPageObj = &StaticPage{}

func (this *StaticPage) GetLocalStaticResource(ctx *context.Context) {
	hz, _ := ioutil.ReadFile("./views/sso/localStatic.tpl")
	ctx.ResponseWriter.Write(hz)
}

func (this *StaticPage) GetSubsystemConfigPage(ctx *context.Context) {
	hz, _ := ioutil.ReadFile("./views/sso/subsystem.tpl")
	ctx.ResponseWriter.Write(hz)
}

func (this *StaticPage) GetProxyStaticPage(ctx *context.Context) {
	hz, _ := ioutil.ReadFile("./views/sso/proxyStatic.tpl")
	ctx.ResponseWriter.Write(hz)
}

func (this *StaticPage) GetRemoteApiPage(ctx *context.Context) {
	hz, _ := ioutil.ReadFile("./views/sso/remoteApi.tpl")
	ctx.ResponseWriter.Write(hz)
}
