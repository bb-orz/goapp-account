package verify

import (
	"errors"
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XMail"
	"goapp/common"
	"goapp/dtos"
	"strconv"
)

/*Verify 校验领域层：实现邮件验证码校验、手机短信验证码校验等相关的校验业务逻辑*/
type VerifyDomain struct {
	cache *verifyCache
}

func NewVerifyDomain() *VerifyDomain {
	domain := new(VerifyDomain)
	domain.cache = NewMailCache()
	return domain
}

func (domain *VerifyDomain) DomainName() string {
	return DomainName
}

// 构造验证邮箱邮件
func (domain *VerifyDomain) emailVerifyCodeCheckEmailAddressMsg(address string, code string) error {
	subject := "Verify Email Code From " + goinfras.XApp().Sctx.Global().GetAppName()
	body := fmt.Sprintf("Verify Code: %s", code)
	err := XMail.XCommonMail().SendSimpleMail(subject, body, XMail.BodyTypePlain, "", []string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}
	return nil
}

// 构造验证邮箱邮件
func (domain *VerifyDomain) emailVerifyCodeResetPasswordMsg(address string, code string) error {
	subject := "[" + goinfras.XApp().Sctx.Global().GetAppName() + "] " + "Reset Password Code From "
	// 设置重置密码的链接
	url := goinfras.XApp().Sctx.Global().GetHost() + "?code=" + code
	body := fmt.Sprintf("Click This link To Reset Your Password: %s", url)

	// 发送邮件
	err := XMail.XCommonMail().SendSimpleMail(subject, body, "text/plain", "", []string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}

	return nil
}

// TODO 构造注册验证码短信
func (domain *VerifyDomain) verifyCodeForPhoneRegisterMsg(phone string, code string) error {

	return nil
}

// TODO 构造手机绑定验证码短信
func (domain *VerifyDomain) verifyCodeForPhoneBindingMsg(phone string, code string) error {

	return nil
}

// TODO 构造登录验证码短信
func (domain *VerifyDomain) verifyCodeForPhoneLoginMsg(phone string, code string) error {

	return nil
}

// 发送验证码到邮箱
func (domain *VerifyDomain) SendEmailVerifyCode(dto dtos.SendEmailVerifyCodeDTO) error {
	var err error
	var code string

	email := dto.Email

	// 生成6位随机字符串
	code = common.RandomString(6)

	// 保存到缓存
	err = domain.cache.SetEmailVerifyCode(dto.VcType, dto.Email, code)
	if err != nil {
		return common.DomainInnerErrorOnCacheSet(err, "SetUserVerifyEmailCode")
	}

	// 发送邮件
	switch dto.VcType {
	case EmailVCTypeCheck:
		return domain.emailVerifyCodeCheckEmailAddressMsg(email, code)
	case EmailVCTypeForgetPassword:
		return domain.emailVerifyCodeResetPasswordMsg(email, code)
	default:
		return common.DomainInnerErrorOnParameter(err, "VcType")
	}
}

// 验证邮箱
func (domain *VerifyDomain) VerifyEmailAddress(email string, vcode string) (bool, error) {
	var code string
	var isExist bool
	// 缓存取出
	if code, isExist = domain.cache.GetEmailVerifyCode(EmailVCTypeCheck, email); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetEmailVerifyCode")
	}
	// 校验
	if vcode == code {
		return true, nil
	}
	return false, nil
}

// 忘记密码时验证重置码
func (domain *VerifyDomain) VerifyResetPasswordCode(email string, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetEmailVerifyCode(EmailVCTypeForgetPassword, email); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetEmailVerifyCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 手机号注册发送验证码
func (domain *VerifyDomain) SendPhoneSmsVerifyCodeMsg(dto dtos.SendPhoneVerifyCodeDTO) error {
	var err error
	var code string

	phone := strconv.Itoa(int(dto.Phone))

	// 生成4位随机数字
	code, err = common.RandomNumber(4)
	if err != nil {
		return common.DomainInnerErrorOnAlgorithm(err, "RandomNumber(4)")
	}

	// 保存到缓存
	err = domain.cache.SetUserPhoneVerifyCode(dto.VcType, phone, code)
	if err != nil {
		return common.DomainInnerErrorOnCacheSet(err, "SetUserPhoneVerifyCode")
	}

	switch dto.VcType {
	case PhoneVCTypeRegister:
		return domain.verifyCodeForPhoneRegisterMsg(phone, code)
	case PhoneVCTypeLogin:
		return domain.verifyCodeForPhoneLoginMsg(phone, code)
	case PhoneVCTypeBinding:
		return domain.verifyCodeForPhoneBindingMsg(phone, code)
	default:
		return common.DomainInnerErrorOnParameter(err, "VcType")
	}
}

// 注册时手机短信验证码校测
func (domain *VerifyDomain) VerifyPhoneForRegister(phone string, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetUserPhoneVerifyCode(PhoneVCTypeRegister, phone); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserPhoneVerifyCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	} else {
		return false, common.DomainInnerErrorOnDecodeData(errors.New("code verify fail"), fmt.Sprintf("[validating code]:%s | [cache code]:%s", vcode, code))
	}
}

// 登录时手机短信验证码校测
func (domain *VerifyDomain) VerifyPhoneForLogin(phone string, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetUserPhoneVerifyCode(PhoneVCTypeLogin, phone); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserPhoneVerifyCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	} else {
		return false, common.DomainInnerErrorOnDecodeData(errors.New("code verify fail"), fmt.Sprintf("[validating code]:%s | [cache code]:%s", vcode, code))
	}
}

// 绑定手机时手机短信验证码校测
func (domain *VerifyDomain) VerifyPhoneForBinding(phone string, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetUserPhoneVerifyCode(PhoneVCTypeBinding, phone); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserPhoneVerifyCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	} else {
		return false, common.DomainInnerErrorOnDecodeData(errors.New("code verify fail"), fmt.Sprintf("[validating code]:%s | [cache code]:%s", vcode, code))
	}
}
