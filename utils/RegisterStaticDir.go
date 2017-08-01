package utils

import (
	"github.com/asofdate/sso-core/service/impl"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func RegisterStaticDir() {
	staticSrv := impl.NewStaticResourceService()
	staticResource, err := staticSrv.Get()
	if err != nil {
		logs.Error("获取静态资源失败", err)
		return
	}

	for _, val := range staticResource {
		beego.SetStaticPath(val.Url, val.Path)
	}
	logs.Info("register static resource successfully.")
}