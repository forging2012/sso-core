package dao

import "github.com/asofdate/sso-core/entity"

type SsoRemoteServiceDao interface {
	Get(serviceCd string) (entity.SsoRemoteService, error)
}
