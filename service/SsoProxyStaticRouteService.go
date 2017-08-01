package service

import "github.com/asofdate/sso-core/entity"

type SsoProxyStaticRouteService interface {
	GetDetails(registerUrl string, serviceCd ...string) (entity.SsoProxyStaticRoute, error)
	Get() ([]entity.SsoProxyStaticRoute, error)
	Post(data entity.SsoProxyStaticRoute) error
	Put(data entity.SsoProxyStaticRoute) error
	Delete(data []entity.SsoProxyStaticRoute) error
}
