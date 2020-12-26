package core

import (
	"github.com/bb-orz/goinfras/XValidate"
	"goinfras-sample-account/common"
	"goinfras-sample-account/core/verified"
	"goinfras-sample-account/dtos"
	"goinfras-sample-account/services"
	"sync"
)

// 实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IMailService = new(MailServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		mailService := new(MailServiceV1)
		services.SetMailService(mailService)
	})
}

// 邮件服务实例V1
type MailServiceV1 struct {
	verifiedDomain *verified.VerifiedDomain
}

// 发送绑定邮箱验证码到指定邮箱
func (service *MailServiceV1) SendEmailForVerified(dto dtos.SendEmailForVerifiedDTO) error {
	var err error
	var verifiedDomain *verified.VerifiedDomain
	verifiedDomain = verified.NewVerifiedDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return  common.ClientErrorOnValidateParameters(err)
	}

	if err = verifiedDomain.SendValidateEmail(dto); err != nil {
		return common.ServerInnerError(err, verifiedDomain.DomainName())
	}

	return nil
}

// 发送忘记密码链接到邮箱
func (service *MailServiceV1) SendEmailForgetPassword(dto dtos.SendEmailForgetPasswordDTO) error {
	var err error
	var verifiedDomain *verified.VerifiedDomain
	verifiedDomain = verified.NewVerifiedDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return  common.ClientErrorOnValidateParameters(err)
	}

	if err = verifiedDomain.SendResetPasswordCodeEmail(dto); err != nil {
		return common.ServerInnerError(err, verifiedDomain.DomainName())
	}

	return nil
}

// TODO 其他邮件相关服务...