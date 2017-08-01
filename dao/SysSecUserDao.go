package dao

import "github.com/asofdate/sso-core/entity"

type SysSecUserDao interface {
	Get(userId string) (entity.SysSecUser, error)
	UpdateStatus(status int, continueErrorCnt int, userId string)
}
