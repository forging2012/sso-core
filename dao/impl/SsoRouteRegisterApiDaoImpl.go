package impl

import (
	"errors"

	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/dbobj"
)

type SsoRouteRegisterApiDaoImpl struct {
}

var sslSql016 = `select service_cd, register_url, route_desc, create_time,create_user,modify_user,modify_time,uuid,remote_url from sso_route_register`
var ssoSql017 = `delete from sso_route_register where uuid = ?`
var ssoSql018 = `insert into sso_route_register(service_cd, register_url, route_desc, create_time, create_user, modify_user, modify_time,uuid,remote_url) values(?,?,?,now(),?,?,now(),uuid(),?)`
var ssoSql019 = `update sso_route_register set remote_url = ?, service_cd = ?, register_url = ?, route_desc = ?, modify_user = ?, modify_time = now() where uuid = ?`

func NewSsoRouteRegisterApiDao() dao.SsoRouteRegisterApiDao {
	return &SsoRouteRegisterApiDaoImpl{}
}

func (this *SsoRouteRegisterApiDaoImpl) Get() ([]entity.SsoRouteRegisterApi, error) {
	var rst []entity.SsoRouteRegisterApi
	err := dbobj.QueryForSlice(sslSql016, &rst)
	return rst, err
}

func (this *SsoRouteRegisterApiDaoImpl) Post(data entity.SsoRouteRegisterApi) error {
	if validator.IsEmpty(data.ServiceCd) {
		return errors.New("子系统编码不能为空")
	}
	_, err := dbobj.Exec(ssoSql018, data.ServiceCd, data.RegisterUrl, data.RouteDesc, data.CreateUser, data.ModifyUser, data.RemoteUrl)
	return err
}

func (this *SsoRouteRegisterApiDaoImpl) Put(data entity.SsoRouteRegisterApi) error {
	if validator.IsEmpty(data.ServiceCd) {
		return errors.New("子系统编码不能为空")
	}
	_, err := dbobj.Exec(ssoSql019, data.RemoteUrl, data.ServiceCd, data.RegisterUrl, data.RouteDesc, data.ModifyUser, data.Uuid)
	return err
}

func (this *SsoRouteRegisterApiDaoImpl) Delete(data []entity.SsoRouteRegisterApi) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range data {
		_, err := tx.Exec(ssoSql017, val.Uuid)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
