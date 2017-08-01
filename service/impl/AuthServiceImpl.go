package impl

import (
	"github.com/asofdate/sso-jwt-auth/utils/crypto/haes"
	"github.com/asofdate/sso-core/dao"
	"github.com/asofdate/sso-core/dao/impl"
	"github.com/asofdate/sso-core/dto"
	"github.com/asofdate/sso-core/service"
	"github.com/astaxie/beego/logs"
)

type AuthServiceImpl struct {
	sysSec dao.SysSecUserDao
}

func NewAuthService() service.AuthService {
	return &AuthServiceImpl{
		sysSec: impl.NewSysSecUserDao(),
	}
}

func (this *AuthServiceImpl) Auth(cdto dto.AuthDto) dto.AuthDto {
	// 查询用户密码
	sysSecUser, err := this.sysSec.Get(cdto.Username)
	if err != nil {
		logs.Error(err)
		cdto.RetCode = 433
		cdto.RetMsg = "获取用户密码信息失败"
		return cdto
	}

	if sysSecUser.UserId == "" {
		logs.Error("用户不存在")
		cdto.RetCode = 441
		cdto.RetMsg = "用户不存在"
		return cdto
	}

	if sysSecUser.StatusId != 0 {
		logs.Error("用户已经被锁定")
		cdto.RetCode = 435
		cdto.RetMsg = "用户已经被锁定"
		return cdto
	}

	if sysSecUser.ContinueErrorCnt > 7 {
		logs.Error("连续错误登陆7次，账号已被锁定")
		cdto.RetCode = 436
		cdto.RetMsg = "连续错误登陆7次，账号已被锁定"
		return cdto
	}

	// 加密用户信息
	password, err := haes.Encrypt(cdto.Password)
	if err != nil {
		logs.Error(err)
		cdto.RetCode = 434
		cdto.RetMsg = "加密用户信息失败"
		return cdto
	}

	// 对比用户密码
	if sysSecUser.UserPasswd == password {
		cdto.RetCode = 200
		cdto.RetMsg = "success"
		this.updateStatus(0, 0, cdto.Username)
		return cdto
	}
	logs.Error("密码不正确")
	cdto.RetCode = 435
	cdto.RetMsg = "用户密码不正确"
	if sysSecUser.ContinueErrorCnt > 6 {
		sysSecUser.StatusId = 1
	}
	this.updateStatus(sysSecUser.StatusId, sysSecUser.ContinueErrorCnt+1, cdto.Username)
	return cdto
}

func (this *AuthServiceImpl) updateStatus(status int, continueErrCnt int, userId string) {
	this.sysSec.UpdateStatus(status, continueErrCnt, userId)
}
