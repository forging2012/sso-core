package dao

import "github.com/asofdate/sso-core/entity"

type SsoRouteRegisterDao interface {
	Get(registerUrl string, serviceCd string) (entity.SsoRouteRegister, error)
}
