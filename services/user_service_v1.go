package services

import (
	"github.com/bb-orz/goinfras/XValidate"
	"goapp/common"
	"goapp/core/user"
	"goapp/core/verify"
	"goapp/dtos"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑
// 接收领域层和dao层的错误包装处理

var _ IUserService = new(UserServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		userService := new(UserServiceV1)
		SetUserService(userService)
	})
}

// 用户服务实例 V1
type UserServiceV1 struct{}

func (service *UserServiceV1) IsEmailAccountExist(dto dtos.IsEmailAccountExistDTO) (bool, error) {
	var err error
	var isExist bool
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 验证用户邮箱是否存在
	if isExist, err = userDomain.IsEmailExist(dto.Email); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if isExist {
		return true, nil
	}

	return false, nil
}

func (service *UserServiceV1) IsPhoneAccountExist(dto dtos.IsPhoneAccountExistDTO) (bool, error) {
	var err error
	var isExist bool
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 验证用户邮箱是否存在
	if isExist, err = userDomain.IsPhoneExist(dto.Phone); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if isExist {
		return true, nil
	}

	return false, nil
}

// 邮箱创建用户账号
func (service *UserServiceV1) CreateUserWithEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UserInfoDTO, error) {
	var err error
	var isExist bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ErrorOnValidate(err)
	}

	// 验证用户邮箱是否存在
	if isExist, err = userDomain.IsEmailExist(dto.Email); err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if isExist {
		return nil, common.ErrorOnVerify("Email Account Exist!")
	}

	userDTO, err = userDomain.CreateUserForEmail(dto)
	if err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	return userDTO.TransToUserInfoDTO(), nil
}

// 手机号码创建用户账号
func (service *UserServiceV1) CreateUserWithPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UserInfoDTO, error) {
	var err error
	var verifyCodeOk bool
	var isExist bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	var verifyDomain *verify.VerifyDomain

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ErrorOnValidate(err)
	}

	// 校验验证码
	verifyDomain = verify.NewVerifyDomain()
	if verifyCodeOk, err = verifyDomain.VerifyPhoneForRegister(dto.Phone, dto.VerifyCode); !verifyCodeOk || err != nil {
		return nil, common.ErrorOnVerify("Phone SMS Verify Code Fail")
	}

	// 验证用户手机号码是否存在
	userDomain = user.NewUserDomain()
	if isExist, err = userDomain.IsPhoneExist(dto.Phone); err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if isExist {
		return nil, common.ErrorOnVerify("Phone Account Exist!")
	}

	userDTO, err = userDomain.CreateUserForPhone(dto)
	if err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	return userDTO.TransToUserInfoDTO(), nil
}

// 邮箱账号登录鉴权
func (service *UserServiceV1) EmailAuth(dto dtos.AuthWithEmailPasswordDTO) (string, error) {
	var err error
	var token string
	var isPass bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	userDomain = user.NewUserDomain()
	if userDTO, isPass, err = userDomain.VerifyPasswordForEmail(dto.Email, dto.Password); isPass && userDTO != nil && err == nil {
		// JWT token
		token, err = userDomain.GenToken(userDTO.Id, userDTO.No, userDTO.Name, userDTO.Avatar)
		if err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	} else if err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if userDTO == nil {
		return "", common.ErrorOnVerify("Email Account Not Exist!")
	} else if !isPass {
		// 校验密码失败
		return "", common.ErrorOnVerify("Password Error!")
	}

	return token, nil
}

// 手机账号短信验证码登录鉴权
func (service *UserServiceV1) PhoneAuth(dto dtos.AuthWithPhonePasswordDTO) (string, error) {
	var err error
	var token string
	var verifyCodeOk bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	var verifyDomain *verify.VerifyDomain

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// 校验验证码
	verifyDomain = verify.NewVerifyDomain()
	if verifyCodeOk, err = verifyDomain.VerifyPhoneForLogin(dto.Phone, dto.VerifyCode); err != nil {
		return "", common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if !verifyCodeOk {
		return "", common.ErrorOnVerify("Phone SMS Verify Code Fail")
	}

	userDomain = user.NewUserDomain()
	if userDTO, err = userDomain.GetUserByPhone(dto.Phone); userDTO != nil && err == nil {
		token, err = userDomain.GenToken(userDTO.Id, userDTO.No, userDTO.Name, userDTO.Avatar)
		if err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	} else if err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if userDTO == nil {
		return "", common.ErrorOnVerify("Phone Account Not Exist!")
	}

	return token, nil
}

// 移除登录令牌缓存信息
func (service *UserServiceV1) RemoveToken(dto dtos.RemoveTokenDTO) error {
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ErrorOnValidate(err)
	}

	err = userDomain.RemoveTokenCache(dto.Token)
	if err != nil {
		return common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	return nil
}

// 获取用户信息
func (service *UserServiceV1) GetUserInfo(dto dtos.GetUserInfoDTO) (*dtos.UserInfoDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ErrorOnValidate(err)
	}

	// 查找用户信息
	userDTO, err = userDomain.GetUser(dto.Id)
	if err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return userDTO.TransToUserInfoDTO(), nil
}

// 注册后验证用户邮箱
func (service *UserServiceV1) EmailValidate(dto dtos.EmailValidateDTO) (bool, error) {
	var err error
	var pass bool
	var emailBinding bool
	var verifyDomain *verify.VerifyDomain
	var userDomain *user.UserDomain
	verifyDomain = verify.NewVerifyDomain()
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 从cache拿出保存的邮箱验证码
	if pass, err = verifyDomain.VerifyEmailAddress(dto.Email, dto.VerifyCode); err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if pass {
		if emailBinding, err = userDomain.IsEmailBinding(dto.Id, dto.Email); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}
		if emailBinding {
			// 设置email_verify字段
			if err = userDomain.SetEmailVerify(dto.Id); err != nil {
				return false, common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			// 设置用户已校验状态
			if err = userDomain.SetUserStatusNormal(dto.Id); err != nil {
				return false, common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return true, nil
		} else {
			return false, common.ErrorOnVerify("Email Not Binding!")
		}
	} else {
		return false, common.ErrorOnVerify("Email Verify Code Fail!")
	}
}

// 注册后验证用户邮箱
func (service *UserServiceV1) EmailBinding(dto dtos.EmailValidateDTO) (bool, error) {
	var err error
	var pass bool
	var emailExist bool
	var verifyDomain *verify.VerifyDomain
	var userDomain *user.UserDomain
	verifyDomain = verify.NewVerifyDomain()
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 从cache拿出保存的邮箱验证码
	if pass, err = verifyDomain.VerifyEmailAddress(dto.Email, dto.VerifyCode); err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if pass {
		if emailExist, err = userDomain.IsEmailExist(dto.Email); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if !emailExist {
			if err = userDomain.SetEmail(dto.Id, dto.Email); err != nil {
				return false, common.ErrorOnServerInner(err, userDomain.DomainName())
			}
		} else {
			return false, common.ErrorOnVerify("Email Account Exist!")
		}

		// 设置email_verify字段
		if err = userDomain.SetEmailVerify(dto.Id); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}
		// 设置用户已校验状态
		if err = userDomain.SetUserStatusNormal(dto.Id); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}
		return true, nil

	} else {
		return false, common.ErrorOnVerify("Email Verify Code Fail!")
	}
}

// 验证手机号码
func (service *UserServiceV1) PhoneBinding(dto dtos.PhoneValidateDTO) (bool, error) {
	var err error
	var pass bool
	var phoneExist bool
	var verifyDomain *verify.VerifyDomain
	var userDomain *user.UserDomain
	verifyDomain = verify.NewVerifyDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 从cache拿出保存的短信验证码
	if pass, err = verifyDomain.VerifyPhoneForBinding(dto.Phone, dto.VerifyCode); err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if pass {
		if phoneExist, err = userDomain.IsPhoneExist(dto.Phone); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if !phoneExist {
			if err = userDomain.SetPhone(dto.Id, dto.Phone); err != nil {
				return false, common.ErrorOnServerInner(err, userDomain.DomainName())
			}
		} else {
			return false, common.ErrorOnVerify("Phone Account Exist!")
		}

		// 设置phone_verify字段
		if err = userDomain.SetPhoneVerify(dto.Id); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		// 设置用户已校验状态
		if err = userDomain.SetUserStatusNormal(dto.Id); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		return true, nil

	} else {
		return false, common.ErrorOnVerify("Sms Verify Code Error!")
	}
}

// 修改用户密码
func (service *UserServiceV1) ModifiedPassword(dto dtos.ModifiedPasswordDTO) (bool, error) {
	var err error
	var isPass bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 校验密码
	if userDTO, isPass, err = userDomain.VerifyPassword(dto.Id, dto.Old); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if userDTO == nil {
		return false, common.ErrorOnVerify("Account Not Exist!")
	} else if !isPass {
		// 校验密码失败
		return false, common.ErrorOnVerify("Password Error!")
	}

	// 设置新密码
	if err = userDomain.ReSetPasswordById(dto.Id, dto.New); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return true, nil
}

// 忘记密码重设
func (service *UserServiceV1) ResetForgetPassword(dto dtos.ResetForgetPasswordDTO) (bool, error) {
	var err error
	var isExist bool
	var isVerify bool
	var userDomain *user.UserDomain
	var verifyDomain *verify.VerifyDomain
	userDomain = user.NewUserDomain()
	verifyDomain = verify.NewVerifyDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 查找账号是否存在
	isExist, err = userDomain.IsEmailExist(dto.Email)
	if err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	if !isExist {
		return false, common.ErrorOnVerify("Account Not Exist!")
	}

	// 校验Code
	isVerify, err = verifyDomain.VerifyResetPasswordCode(dto.Email, dto.VerifyCode)
	if err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if !isVerify {
		return false, common.ErrorOnVerify("Reset Password Fail,Please Retry!")
	}

	// 重设密码
	if err = userDomain.ReSetPasswordByEmail(dto.Email, dto.New); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return true, nil
}

// 设置用户头像链接
func (service *UserServiceV1) SetAvatarUri(dto dtos.SetAvatarUriDTO) (bool, error) {

	// userDomain.SetUserInfo()
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	err = userDomain.SetAvatar(dto.Id, dto.Avatar)
	if err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return true, nil
}

// 更新用户信息多个字段
func (service *UserServiceV1) SetUserInfos(dto dtos.SetUserInfoDTO) (bool, error) {

	// userDomain.SetUserInfo()
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	err = userDomain.UpdateUsers(dto)
	if err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return true, nil
}
