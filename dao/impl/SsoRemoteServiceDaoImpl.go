package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/hzwy23/dbobj"
)

type SsoRemoteServiceDaoImpl struct {
}

var ssoSql005 = `select service_cd,remote_scheme,remote_host,remote_port,prefix_url from sso_remote_service where service_cd = ?`

func NewSsoRemoteServiceDao() dao.SsoRemoteServiceDao {
	return &SsoRemoteServiceDaoImpl{}
}

func (this *SsoRemoteServiceDaoImpl) Get(serviceCd string) (entity.SsoRemoteService, error) {
	var ret entity.SsoRemoteService
	err := dbobj.QueryForStruct(ssoSql005, &ret, serviceCd)
	return ret, err
}
