package service

import "github.com/asofdate/sso-core/dto"

type AuthService interface {
	Auth(cdto dto.AuthDto) dto.AuthDto
}
