package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/hzwy23/dbobj"
)

type SysSecUserDaoImpl struct {
}

var ssoSql002 = `select user_id,user_passwd,status_id,continue_error_cnt from sys_sec_user where user_id = ?`
var ssoSql003 = `update sys_sec_user set status_id = ?, continue_error_cnt = ? where user_id = ?`

func NewSysSecUserDao() dao.SysSecUserDao {
	return &SysSecUserDaoImpl{}
}

func (this *SysSecUserDaoImpl) Get(userId string) (entity.SysSecUser, error) {
	var ret entity.SysSecUser
	err := dbobj.QueryForStruct(ssoSql002, &ret, userId)
	return ret, err
}

func (this *SysSecUserDaoImpl) UpdateStatus(status int, continueErrorCnt int, userId string) {
	dbobj.Exec(ssoSql003, status, continueErrorCnt, userId)
}
