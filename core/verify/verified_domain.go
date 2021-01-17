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
	return "VerifyDomain"
}

// 生成邮箱验证码
func (domain *VerifyDomain) genEmailVerifyCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = common.RandomString(6)

	// 保存到缓存
	err = domain.cache.SetUserVerifyEmailCode(uid, code)
	if err != nil {
		return "", common.DomainInnerErrorOnCacheSet(err, "SetUserVerifyEmailCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifyDomain) sendValidateEmail(address string, code string) error {
	from := "no-reply@" + goinfras.XApp().Sctx.Global().GetHost()
	subject := "Verify Email Code From " + goinfras.XApp().Sctx.Global().GetAppName()
	body := fmt.Sprintf("Verify Code: %s", code)
	err := XMail.XCommonMail().SendSimpleMail(from, "", "", subject, body, "text/plain", "", []string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}
	return nil
}

// 发送验证码到邮箱
func (domain *VerifyDomain) SendValidateEmail(dto dtos.SendEmailForVerifyDTO) error {
	var err error
	var code string

	uid := dto.Id
	email := dto.Email

	code, err = domain.genEmailVerifyCode(uid)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendValidateEmail(email, code)
}

// 注册时验证邮箱
func (domain *VerifyDomain) VerifyEmail(uid uint, vcode string) (bool, error) {
	var code string
	var isExist bool
	// 缓存取出
	if code, isExist = domain.cache.GetUserVerifyEmailCode(uid); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserVerifyEmailCode")
	}
	// 校验
	if vcode == code {
		return true, nil
	}
	return false, nil
}

// 生成邮箱验证码
func (domain *VerifyDomain) genResetPasswordCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = common.RandomString(40)

	// 保存到缓存
	err = domain.cache.SetForgetPasswordVerifyCode(uid, code)
	if err != nil {
		return "", common.DomainInnerErrorOnCacheSet(err, "SetForgetPasswordVerifyCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifyDomain) sendResetPasswordCodeEmail(address string, code string) error {
	from := "no-reply@" + goinfras.XApp().Sctx.Global().GetHost()
	subject := "Reset Password Code From " + goinfras.XApp().Sctx.Global().GetAppName()
	// 设置重置密码的链接
	url := goinfras.XApp().Sctx.Global().GetHost() + "?code=" + code
	body := fmt.Sprintf("Click This link To Reset Your Password: %s", url)

	// 发送邮件
	err := XMail.XCommonMail().SendSimpleMail(from, "", "", subject, body, "text/plain", "", []string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}

	return nil
}

// 发送验证码到邮箱
func (domain *VerifyDomain) SendResetPasswordCodeEmail(dto dtos.SendEmailForgetPasswordDTO) error {

	code, err := domain.genResetPasswordCode(dto.Id)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendResetPasswordCodeEmail(dto.Email, code)
}

// 忘记密码时验证重置码
func (domain *VerifyDomain) VerifyResetPasswordCode(uid uint, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetForgetPasswordVerifyCode(uid); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetForgetPasswordVerifyCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成手机短信验证码
func (domain *VerifyDomain) genPhoneVerifyCode(uid uint) (string, error) {
	var err error
	var code string

	// 生成4位随机数字
	code, err = common.RandomNumber(4)
	if err != nil {
		return "", common.DomainInnerErrorOnAlgorithm(err, "XGlobal.RandomNumber(4)")
	}

	// 保存到缓存
	err = domain.cache.SetUserVerifyPhoneCode(uid, code)
	if err != nil {
		return "", common.DomainInnerErrorOnCacheSet(err, "SetUserVerifyPhoneCode")
	}

	return code, nil
}

// 构造短信
func (domain *VerifyDomain) sendValidatePhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *VerifyDomain) SendValidatePhoneMsg(dto dtos.SendPhoneVerifyCodeDTO) error {
	var err error
	var code string

	uid := dto.Id
	phone := strconv.Itoa(int(dto.Phone))

	code, err = domain.genPhoneVerifyCode(uid)
	if err != nil {
		return err
	}

	return domain.sendValidatePhoneMsg(phone, code)
}

// 注册时验证手机短信
func (domain *VerifyDomain) VerifyPhone(uid uint, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetUserVerifyPhoneCode(uid); !isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserVerifyPhoneCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	} else {
		return false, common.DomainInnerErrorOnDecodeData(errors.New("code verify fail"), fmt.Sprintf("[validating code]:%s | [cache code]:%s", vcode, code))
	}
}
