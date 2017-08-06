package utils

import (
	"github.com/asofdate/sso-core/service/impl"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/astaxie/beego"
)

func RegisterStaticDir() {
	staticSrv := impl.NewStaticResourceService()
	staticResource, err := staticSrv.Get()
	if err != nil {
		logger.Error("获取静态资源失败", err)
		return
	}

	for _, val := range staticResource {
		logger.Debug("register static route, url is:", val.Url, ",local path is:", val.Path)
		beego.SetStaticPath(val.Url, val.Path)
	}
	logger.Info("register static resource successfully.")
}
