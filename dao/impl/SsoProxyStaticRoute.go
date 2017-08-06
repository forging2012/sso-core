package impl

import (
	"errors"

	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/dbobj"
)

type SsoProxyStaticRoute struct {
}

var ssoSql020 = `select register_url,route_desc,create_time,remote_url,create_user,modify_user,modify_time, uuid,service_cd from sso_proxy_static_route`
var ssoSql021 = `insert into sso_proxy_static_route(register_url,route_desc,create_time,remote_url,create_user,modify_user,modify_time,uuid,service_cd) values(?,?,now(),?,?,?,now(),uuid(),?)`
var ssoSql022 = `delete from sso_proxy_static_route where uuid = ?`
var ssoSql023 = `update sso_proxy_static_route set service_cd = ?, register_url = ?,remote_url = ?,route_desc = ?, modify_time = now(), modify_user = ? where uuid = ?`
var ssoSql025 = `select register_url,route_desc,create_time,remote_url,create_user,modify_user,modify_time, uuid,service_cd from sso_proxy_static_route where register_url = ?`
var ssoSql026 = `select register_url,route_desc,create_time,remote_url,create_user,modify_user,modify_time, uuid,service_cd from sso_proxy_static_route where register_url = ? and service_cd = ?`

func NewSsoProxyStaticRouteDao() dao.SsoProxyStaticRouteDao {
	return &SsoProxyStaticRoute{}
}

func (this *SsoProxyStaticRoute) GetDetails(registerUrl string, serviceCd ...string) (entity.SsoProxyStaticRoute, error) {
	var ret entity.SsoProxyStaticRoute
	if len(serviceCd) == 1 && len(serviceCd[0]) != 0 {
		err := dbobj.QueryForStruct(ssoSql026, &ret, registerUrl, serviceCd[0])
		return ret, err
	}
	err := dbobj.QueryForStruct(ssoSql025, &ret, registerUrl)
	return ret, err
}

func (this *SsoProxyStaticRoute) Get() ([]entity.SsoProxyStaticRoute, error) {
	var rst []entity.SsoProxyStaticRoute
	err := dbobj.QueryForSlice(ssoSql020, &rst)
	return rst, err
}

func (this *SsoProxyStaticRoute) Post(data entity.SsoProxyStaticRoute) error {
	if validator.IsEmpty(data.RegisterUrl) {
		return errors.New("路由地址不能为空")
	}
	_, err := dbobj.Exec(ssoSql021, data.RegisterUrl, data.RouteDesc, data.RemoteUrl, data.CreateUser, data.ModifyUser, data.ServiceCd)
	return err
}

func (this *SsoProxyStaticRoute) Put(data entity.SsoProxyStaticRoute) error {
	if validator.IsEmpty(data.RegisterUrl) {
		return errors.New("路由地址不能为空")
	}
	_, err := dbobj.Exec(ssoSql023, data.ServiceCd, data.RegisterUrl, data.RemoteUrl, data.RouteDesc, data.ModifyUser, data.Uuid)
	return err
}

func (this *SsoProxyStaticRoute) Delete(data []entity.SsoProxyStaticRoute) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range data {
		_, err := tx.Exec(ssoSql022, val.Uuid)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
