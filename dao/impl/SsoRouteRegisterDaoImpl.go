package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/hzwy23/dbobj"
)

type SsoRouteRegisterDaoImpl struct {
}

var ssoSql004 = `select register_url,service_cd,remote_url from  sso_route_register where register_url = ? and service_cd = ?`
var ssoSql006 = `select register_url,service_cd,remote_url from  sso_route_register where register_url = ?`

func NewSsoRouteRegisterDao() dao.SsoRouteRegisterDao {
	return &SsoRouteRegisterDaoImpl{}
}
func (this *SsoRouteRegisterDaoImpl) Get(registerUrl string, serviceCd string) (entity.SsoRouteRegister, error) {
	var ret entity.SsoRouteRegister
	var err error
	if len(serviceCd) == 0 {
		err = dbobj.QueryForStruct(ssoSql006, &ret, registerUrl)
	} else {
		err = dbobj.QueryForStruct(ssoSql004, &ret, registerUrl, serviceCd)
	}
	return ret, err
}
