package service

import "github.com/asofdate/sso-core/entity"

type SsoRouteRegisterApiService interface {
	Get() ([]entity.SsoRouteRegisterApi, error)
	Post(data entity.SsoRouteRegisterApi) error
	Put(data entity.SsoRouteRegisterApi) error
	Delete(data []entity.SsoRouteRegisterApi) error
}
