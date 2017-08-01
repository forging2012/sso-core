package dao

import "github.com/asofdate/sso-core/entity"

type StaticResourceDao interface {
	Get() ([]entity.StaticResource, error)
}
