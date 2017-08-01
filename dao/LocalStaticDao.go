package dao

import "github.com/asofdate/sso-core/entity"

type LocalStaticDao interface {
	Get() ([]entity.LocalStaticEntity, error)
	Post(data entity.LocalStaticEntity) error
	Delete(data []entity.LocalStaticEntity) error
	Put(data entity.LocalStaticEntity) error
}
