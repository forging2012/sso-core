package dao

import "github.com/asofdate/sso-core/entity"

type SysUserInfoDao interface {
	Get(userId string) (entity.SysUserInfo, error)
}
