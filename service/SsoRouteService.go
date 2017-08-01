package service

import (
	"github.com/asofdate/sso-core/dto"
)

type SsoRouteService interface {
	Get(registerUrl string, serviceCd string) (dto.SsoServiceDto, error)
	GetProxyStatic(registerUrl string, serviceCd ...string) (dto.SsoServiceDto, error)
}
