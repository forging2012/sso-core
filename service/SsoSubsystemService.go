package service

import "github.com/asofdate/sso-core/entity"

type SsoSubsystemService interface {
	Get() ([]entity.SsoSubsystemEntity, error)
	Post(data entity.SsoSubsystemEntity) error
	Put(data entity.SsoSubsystemEntity) error
	Delete(data []entity.SsoSubsystemEntity) error
}
