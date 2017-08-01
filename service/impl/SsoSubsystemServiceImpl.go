package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service"
)

type SsoSubsystemServiceImpl struct {
	ssoSubsystemDao dao.SsoSubsystemDao
}

func NewSsoSubsystemService() service.SsoSubsystemService {
	return &SsoSubsystemServiceImpl{
		ssoSubsystemDao: impl.NewSsoSubsystemDao(),
	}
}

func (this *SsoSubsystemServiceImpl) Get() ([]entity.SsoSubsystemEntity, error) {
	return this.ssoSubsystemDao.Get()
}

func (this *SsoSubsystemServiceImpl) Post(data entity.SsoSubsystemEntity) error {
	return this.ssoSubsystemDao.Post(data)
}

func (this *SsoSubsystemServiceImpl) Put(data entity.SsoSubsystemEntity) error {
	return this.ssoSubsystemDao.Put(data)
}

func (this *SsoSubsystemServiceImpl) Delete(data []entity.SsoSubsystemEntity) error {
	return this.ssoSubsystemDao.Delete(data)
}
