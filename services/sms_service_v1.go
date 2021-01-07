package services

import (
	"github.com/bb-orz/goinfras/XValidate"
	"goapp/common"
	"goapp/core/verify"
	"goapp/dtos"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ ISmsService = new(SmsServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		smsService := new(SmsServiceV1)
		SetSmsService(smsService)
	})
}

// 短信服务实例V1
type SmsServiceV1 struct{}

// 发送绑定手机短信验证码
func (service *SmsServiceV1) SendPhoneVerifyCode(dto dtos.SendPhoneVerifyCodeDTO) error {
	var err error
	var verifyDomain *verify.VerifyDomain
	verifyDomain = verify.NewVerifyDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ErrorOnValidate(err)
	}

	if err = verifyDomain.SendValidatePhoneMsg(dto); err != nil {
		return common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	return nil

}

// TODO 其他短信相关服务...

// 发送忘记密码短信验证码
