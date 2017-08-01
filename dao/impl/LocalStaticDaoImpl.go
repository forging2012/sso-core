package impl

import (
	"errors"

	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/entity"
	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/dbobj"
)

type LocalStaticDaoImpl struct {
}

var ssoSql008 = `select url,path,create_time,create_user,modify_time,modify_user,uuid from sso_local_static_route`
var ssoSql009 = `insert into sso_local_static_route(url,path,create_time,create_user,modify_time,modify_user,uuid) values(?,?,now(),?,now(),?,uuid())`
var ssoSql010 = `delete from sso_local_static_route where uuid = ?`
var ssoSql011 = `update sso_local_static_route set url = ?, path = ?, modify_user = ?, modify_time = now() where uuid = ?`

func NewLocalStaticDao() dao.LocalStaticDao {
	return &LocalStaticDaoImpl{}
}

func (this *LocalStaticDaoImpl) Put(data entity.LocalStaticEntity) error {
	if validator.IsEmpty(data.Url) {
		return errors.New("路由信息格式不正确")
	}

	if validator.IsEmpty(data.Path) {
		return errors.New("本地路径不正确")
	}

	_, err := dbobj.Exec(ssoSql011, data.Url, data.Path, data.ModifyUser, data.Uuid)
	if err != nil {
		logger.Error(err)
		return errors.New("配置本地静态资源失败，失败原因是：" + err.Error())
	}
	return nil
}

func (this *LocalStaticDaoImpl) Get() ([]entity.LocalStaticEntity, error) {
	var ret []entity.LocalStaticEntity
	err := dbobj.QueryForSlice(ssoSql008, &ret)
	return ret, err
}

func (this *LocalStaticDaoImpl) Delete(data []entity.LocalStaticEntity) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range data {
		_, err := tx.Exec(ssoSql010, val.Uuid)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (this *LocalStaticDaoImpl) Post(data entity.LocalStaticEntity) error {

	if validator.IsEmpty(data.Url) {
		return errors.New("路由信息格式不正确")
	}

	if validator.IsEmpty(data.Path) {
		return errors.New("本地路径不正确")
	}

	_, err := dbobj.Exec(ssoSql009, data.Url, data.Path, data.CreateUser, data.ModifyUser)
	if err != nil {
		logger.Error(err)
		return errors.New("配置本地静态资源失败，失败原因是：" + err.Error())
	}
	return nil
}
