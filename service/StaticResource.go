package service

import "github.com/asofdate/sso-core/entity"

type StaticResourceService interface {
	Get() ([]entity.StaticResource, error)
	GetList() ([]entity.LocalStaticEntity, error)
	Post(data entity.LocalStaticEntity) error
	Delete(data []entity.LocalStaticEntity) error
	Put(data entity.LocalStaticEntity) error
}
