package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-core/service"
)

type StaticResourceServiceImpl struct {
	staticDao      dao.StaticResourceDao
	localStaticDao dao.LocalStaticDao
}

func NewStaticResourceService() service.StaticResourceService {
	return &StaticResourceServiceImpl{
		staticDao:      impl.NewStaticResourceDao(),
		localStaticDao: impl.NewLocalStaticDao(),
	}
}

func (this *StaticResourceServiceImpl) Put(data entity.LocalStaticEntity) error {
	return this.localStaticDao.Put(data)
}

func (this *StaticResourceServiceImpl) Get() ([]entity.StaticResource, error) {
	return this.staticDao.Get()
}

func (this *StaticResourceServiceImpl) GetList() ([]entity.LocalStaticEntity, error) {
	return this.localStaticDao.Get()
}

func (this *StaticResourceServiceImpl) Post(data entity.LocalStaticEntity) error {
	return this.localStaticDao.Post(data)
}

func (this *StaticResourceServiceImpl) Delete(data []entity.LocalStaticEntity) error {
	return this.localStaticDao.Delete(data)
}
