package impl

import (
	"errors"

	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/dbobj"
)

type SsoSubsystemDaoImpl struct {
}

var ssoSql012 = `select service_cd, remote_scheme, remote_host, remote_port, create_time, modify_time, create_user, modify_user, service_name, prefix_url from sso_remote_service`
var ssoSql013 = `insert into sso_remote_service(service_cd, remote_scheme, remote_host,remote_port,create_time,modify_time,create_user,modify_user,service_name,prefix_url) values(?,?,?,?,now(),now(),?,?,?,?)`
var ssoSql014 = `delete from sso_remote_service where service_cd = ?`
var ssoSql015 = `update sso_remote_service set prefix_url = ?, remote_scheme = ?, remote_host = ?, remote_port = ?, modify_time = now(), modify_user = ?, service_name = ? where service_cd = ?`

func NewSsoSubsystemDao() dao.SsoSubsystemDao {
	return &SsoSubsystemDaoImpl{}
}

func (this *SsoSubsystemDaoImpl) Get() ([]entity.SsoSubsystemEntity, error) {
	var rst []entity.SsoSubsystemEntity
	err := dbobj.QueryForSlice(ssoSql012, &rst)
	return rst, err
}

func (this *SsoSubsystemDaoImpl) Post(data entity.SsoSubsystemEntity) error {

	if validator.IsEmpty(data.ServiceCd) {
		return errors.New("子系统编码不能为空")
	}
	_, err := dbobj.Exec(ssoSql013, data.ServiceCd, data.RemoteScheme, data.RemoteHost, data.RemotePort, data.CreateUser, data.ModifyUser, data.ServiceName, data.PrefixUrl)
	return err
}

func (this *SsoSubsystemDaoImpl) Put(data entity.SsoSubsystemEntity) error {
	if validator.IsEmpty(data.ServiceCd) {
		return errors.New("子系统编码不能为空")
	}
	_, err := dbobj.Exec(ssoSql015, data.PrefixUrl, data.RemoteScheme, data.RemoteHost, data.RemotePort, data.ModifyUser, data.ServiceName, data.ServiceCd)
	return err
}

func (this *SsoSubsystemDaoImpl) Delete(data []entity.SsoSubsystemEntity) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range data {
		_, err = tx.Exec(ssoSql014, val.ServiceCd)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
