package impl

import (
	"github.com/asofdate/sso-core/entity"
	"github.com/hzwy23/dbobj"
)

type SysUserInfoDaoImpl struct {
}

var ssoSql001 = `"select user_id,user_name,user_email,user_phone,org_unit_id from sys_user_info where user_id = ?"`

func (this *SysUserInfoDaoImpl) Get(userId string) (entity.SysUserInfo, error) {
	var ret entity.SysUserInfo
	err := dbobj.QueryForStruct(ssoSql001, &ret, userId)
	return ret, err
}
