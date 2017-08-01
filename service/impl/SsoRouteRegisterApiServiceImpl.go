package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service"
)

type SsoRouteRegisterApiServiceImpl struct {
	ssoRouteRegisterApiDao dao.SsoRouteRegisterApiDao
}

func NewSsoRouteRegisterApiService() service.SsoRouteRegisterApiService {
	return &SsoRouteRegisterApiServiceImpl{
		ssoRouteRegisterApiDao: impl.NewSsoRouteRegisterApiDao(),
	}
}

func (this *SsoRouteRegisterApiServiceImpl) Get() ([]entity.SsoRouteRegisterApi, error) {
	return this.ssoRouteRegisterApiDao.Get()
}

func (this *SsoRouteRegisterApiServiceImpl) Post(data entity.SsoRouteRegisterApi) error {
	return this.ssoRouteRegisterApiDao.Post(data)
}
func (this *SsoRouteRegisterApiServiceImpl) Put(data entity.SsoRouteRegisterApi) error {
	return this.ssoRouteRegisterApiDao.Put(data)
}
func (this *SsoRouteRegisterApiServiceImpl) Delete(data []entity.SsoRouteRegisterApi) error {
	return this.ssoRouteRegisterApiDao.Delete(data)
}
