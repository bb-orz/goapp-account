package verified

import (
	"errors"
	"fmt"
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XMail"
	"goinfras-sample-account/common"
	"goinfras-sample-account/services"
	"strconv"
)

/*Verified 校验领域层：实现邮件绑定、手机短信绑定相关的校验业务逻辑*/
type VerifiedDomain struct {
	cache *verifiedCache
}

func NewVerifiedDomain() *VerifiedDomain {
	domain := new(VerifiedDomain)
	domain.cache = NewMailCache()
	return domain
}

func  (domain *VerifiedDomain)DomainName() string {
	return "VerifiedDomain"
}

// 生成邮箱验证码
func (domain *VerifiedDomain) genEmailVerifiedCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = XGlobal.RandomString(6)

	// 保存到缓存
	err = domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "",common.DomainInnerErrorOnCacheSet(err, "SetUserVerifiedEmailCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifiedDomain) sendValidateEmail(address string, code string) error {
	from := "no-reply@" + XGlobal.GetHost()
	subject := "Verified Email Code From " + XGlobal.GetAppName()
	body := fmt.Sprintf("Verified Code: %s", code)
	err := XMail.XCommonMail().SendSimpleMail(from,"","", subject, body,"text/plain","",[]string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}
	return nil
}

// 发送验证码到邮箱
func (domain *VerifiedDomain) SendValidateEmail(dto services.SendEmailForVerifiedDTO) error {
	var err error
	var code string

	uid := dto.ID
	email := dto.Email

	code, err = domain.genEmailVerifiedCode(uid)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendValidateEmail(email, code)
}

// 注册时验证邮箱
func (domain *VerifiedDomain) VerifiedEmail(uid uint, vcode string) (bool, error) {
	var code string
	var isExist bool
	// 缓存取出
	if code, isExist = domain.cache.GetUserVerifiedEmailCode(uid);!isExist {
		return false,common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetUserVerifiedEmailCode")
	}
	// 校验
	if vcode == code {
		return true, nil
	}
	return false, nil
}

// 生成邮箱验证码
func (domain *VerifiedDomain) genResetPasswordCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = XGlobal.RandomString(40)

	// 保存到缓存
	err = domain.cache.SetForgetPasswordVerifiedCode(uid, code)
	if err != nil {
		return "",common.DomainInnerErrorOnCacheSet(err, "SetForgetPasswordVerifiedCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifiedDomain) sendResetPasswordCodeEmail(address string, code string) error {
	from := "no-reply@" + XGlobal.GetHost()
	subject := "Reset Password Code From " + XGlobal.GetAppName()
	// 设置重置密码的链接
	url := XGlobal.GetHost() + "?code=" + code
	body := fmt.Sprintf("Click This link To Reset Your Password: %s", url)

	// 发送邮件
	err := XMail.XCommonMail().SendSimpleMail(from,"","", subject, body,"text/plain","",[]string{address})
	if err != nil {
		return common.DomainInnerErrorOnNetRequest(err, "SendSimpleMail")
	}

	return nil
}

// 发送验证码到邮箱
func (domain *VerifiedDomain) SendResetPasswordCodeEmail(dto services.SendEmailForgetPasswordDTO) error {

	code, err := domain.genResetPasswordCode(dto.ID)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendResetPasswordCodeEmail(dto.Email, code)
}

// 忘记密码时验证重置码
func (domain *VerifiedDomain) VerifiedResetPasswordCode(uid uint, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetForgetPasswordVerifiedCode(uid);!isExist{
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"), "GetForgetPasswordVerifiedCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成手机短信验证码
func (domain *VerifiedDomain) genPhoneVerifiedCode(uid uint) (string, error) {
	var err error
	var code string

	// 生成4位随机数字
	code, err = XGlobal.RandomNumber(4)
	if err != nil {
		return "", common.DomainInnerErrorOnAlgorithm(err,  "XGlobal.RandomNumber(4)")
	}

	// 保存到缓存
	err = domain.cache.SetUserVerifiedPhoneCode(uid, code)
	if err != nil {
		return "", common.DomainInnerErrorOnCacheSet(err, "SetUserVerifiedPhoneCode")
	}

	return code, nil
}

// 构造短信
func (domain *VerifiedDomain) sendValidatePhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *VerifiedDomain) SendValidatePhoneMsg(dto services.SendPhoneVerifiedCodeDTO) error {
	var err error
	var code string

	uid := dto.ID
	phone := strconv.Itoa(int(dto.Phone))

	code, err = domain.genPhoneVerifiedCode(uid)
	if err != nil {
		return err
	}

	return domain.sendValidatePhoneMsg(phone, code)
}

// 注册时验证手机短信
func (domain *VerifiedDomain) VerifiedPhone(uid uint, vcode string) (bool, error) {
	var isExist bool
	var code string

	// 缓存取出
	if code, isExist = domain.cache.GetUserVerifiedPhoneCode(uid);!isExist {
		return false, common.DomainInnerErrorOnCacheGet(errors.New("cache code not exist"),"GetUserVerifiedPhoneCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}else {
		return false,common.DomainInnerErrorOnDecodeData(errors.New("code verified fail"),fmt.Sprintf("[validating code]:%s | [cache code]:%s",vcode,code))
	}
}
