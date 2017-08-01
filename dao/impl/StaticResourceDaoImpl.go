package impl

import (
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/hzwy23/dbobj"
)

type StaticResourceDaoImpl struct {
}

var ssoSql007 = `select url,path from sso_local_static_route`

func NewStaticResourceDao() dao.StaticResourceDao {
	return &StaticResourceDaoImpl{}
}

func (this *StaticResourceDaoImpl) Get() ([]entity.StaticResource, error) {
	var ret []entity.StaticResource
	err := dbobj.QueryForSlice(ssoSql007, &ret)
	return ret, err
}
