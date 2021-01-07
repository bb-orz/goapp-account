package services

import (
	"github.com/bb-orz/goinfras/XValidate"
	"goapp/common"
	"goapp/core/verify"
	"goapp/dtos"
	"sync"
)

// 实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ IMailService = new(MailServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		mailService := new(MailServiceV1)
		SetMailService(mailService)
	})
}

// 邮件服务实例V1
type MailServiceV1 struct {
	verifyDomain *verify.VerifyDomain
}

// 发送绑定邮箱验证码到指定邮箱
func (service *MailServiceV1) SendEmailForVerify(dto dtos.SendEmailForVerifyDTO) error {
	var err error
	var verifyDomain *verify.VerifyDomain
	verifyDomain = verify.NewVerifyDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ErrorOnValidate(err)
	}

	if err = verifyDomain.SendValidateEmail(dto); err != nil {
		return common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	return nil
}

// 发送忘记密码链接到邮箱
func (service *MailServiceV1) SendEmailForgetPassword(dto dtos.SendEmailForgetPasswordDTO) error {
	var err error
	var verifyDomain *verify.VerifyDomain
	verifyDomain = verify.NewVerifyDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ErrorOnValidate(err)
	}

	if err = verifyDomain.SendResetPasswordCodeEmail(dto); err != nil {
		return common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	return nil
}

// TODO 其他邮件相关服务...
