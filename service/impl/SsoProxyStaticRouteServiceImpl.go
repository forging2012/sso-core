package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service"
)

type SsoProxyStaticRouteServiceImpl struct {
	ssoProxyStatic dao.SsoProxyStaticRouteDao
}

func NewSsoProxyStaticRouteService() service.SsoProxyStaticRouteService {
	return &SsoProxyStaticRouteServiceImpl{
		ssoProxyStatic: impl.NewSsoProxyStaticRouteDao(),
	}
}

func (this *SsoProxyStaticRouteServiceImpl) GetDetails(registerUrl string, serviceCd ...string) (entity.SsoProxyStaticRoute, error) {
	return this.ssoProxyStatic.GetDetails(registerUrl, serviceCd...)
}

func (this *SsoProxyStaticRouteServiceImpl) Get() ([]entity.SsoProxyStaticRoute, error) {
	return this.ssoProxyStatic.Get()
}

func (this *SsoProxyStaticRouteServiceImpl) Post(data entity.SsoProxyStaticRoute) error {
	return this.ssoProxyStatic.Post(data)
}

func (this *SsoProxyStaticRouteServiceImpl) Put(data entity.SsoProxyStaticRoute) error {
	return this.ssoProxyStatic.Put(data)
}

func (this *SsoProxyStaticRouteServiceImpl) Delete(data []entity.SsoProxyStaticRoute) error {
	return this.ssoProxyStatic.Delete(data)
}
