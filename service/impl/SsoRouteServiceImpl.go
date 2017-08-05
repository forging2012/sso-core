package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/dto"
	"github.com/asofdate/sso-core/service"
	"github.com/astaxie/beego/logs"
)

type SsoRouteServiceImpl struct {
	ssoRoute       dao.SsoRouteRegisterDao
	ssoRemote      dao.SsoRemoteServiceDao
	ssoProxyStatic dao.SsoProxyStaticRouteDao
}

func NewSsoRouteService() service.SsoRouteService {
	return &SsoRouteServiceImpl{
		ssoRoute:       impl.NewSsoRouteRegisterDao(),
		ssoRemote:      impl.NewSsoRemoteServiceDao(),
		ssoProxyStatic: impl.NewSsoProxyStaticRouteDao(),
	}
}

func (this *SsoRouteServiceImpl) Get(registerUrl string, serviceCd string) (dto.SsoServiceDto, error) {
	var ssoService dto.SsoServiceDto

	reg, err := this.ssoRoute.Get(registerUrl, serviceCd)
	if err != nil {
		logs.Error(err)
		return ssoService, err
	}

	srv, err := this.ssoRemote.Get(reg.ServiceCd)
	if err != nil {
		logs.Error(err)
		return ssoService, err
	}

	ssoService.RegisterUrl = reg.RegisterUrl
	ssoService.ServiceCd = reg.ServiceCd
	ssoService.RemoteUrl = reg.RemoteUrl
	ssoService.RemoteScheme = srv.RemoteScheme
	ssoService.RemoteHost = srv.RemoteHost
	ssoService.RemotePort = srv.RemotePort
	ssoService.PrefixUrl = srv.PrefixUrl
	return ssoService, nil
}

func (this *SsoRouteServiceImpl) GetProxyStatic(registerUrl string, serviceCd ...string) (dto.SsoServiceDto, error) {
	var ssoService dto.SsoServiceDto

	reg, err := this.ssoProxyStatic.GetDetails(registerUrl, serviceCd...)
	if err != nil {
		logs.Error(err)
		return ssoService, err
	}

	srv, err := this.ssoRemote.Get(reg.ServiceCd)
	if err != nil {
		logs.Error(err)
		return ssoService, err
	}

	ssoService.RegisterUrl = reg.RegisterUrl
	ssoService.ServiceCd = reg.ServiceCd
	ssoService.RemoteUrl = reg.RemoteUrl
	ssoService.RemoteScheme = srv.RemoteScheme
	ssoService.RemoteHost = srv.RemoteHost
	ssoService.RemotePort = srv.RemotePort
	ssoService.PrefixUrl = srv.PrefixUrl
	return ssoService, nil
}
