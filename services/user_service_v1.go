package services

import (
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XValidate"
	"goapp/common"
	"goapp/core/third"
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

// 邮箱创建用户账号
func (service *UserServiceV1) CreateUserWithEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UsersDTO, error) {
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
	return userDTO, nil
}

// 手机号码创建用户账号
func (service *UserServiceV1) CreateUserWithPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UsersDTO, error) {
	var err error
	var isExist bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ErrorOnValidate(err)
	}

	// 验证用户手机号码是否存在
	if isExist, err = userDomain.IsPhoneExist(dto.Phone); err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if isExist {
		return nil, common.ErrorOnVerify("Phone Account Exist!")
	}

	userDTO, err = userDomain.CreateUserForPhone(dto)
	if err != nil {
		return nil, common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	return userDTO, nil
}

// 邮箱账号登录鉴权
func (service *UserServiceV1) EmailAuth(dto dtos.AuthWithEmailPasswordDTO) (string, error) {
	var err error
	var token string
	var isPass bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

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

// 手机账号登录鉴权
func (service *UserServiceV1) PhoneAuth(dto dtos.AuthWithPhonePasswordDTO) (string, error) {
	var err error
	var token string
	var isPass bool
	var userDTO *dtos.UsersDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	if userDTO, isPass, err = userDomain.VerifyPasswordForPhone(dto.Phone, dto.Password); isPass && userDTO != nil && err == nil {
		// JWT token
		token, err = userDomain.GenToken(userDTO.Id, userDTO.No, userDTO.Name, userDTO.Avatar)
		if err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	} else if err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if userDTO == nil {
		return "", common.ErrorOnVerify("Phone Account Not Exist!")
	} else if !isPass {
		// 校验密码失败
		return "", common.ErrorOnVerify("Password Error!")
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

// 批量设置用户信息
func (service *UserServiceV1) SetUserInfos(dto dtos.SetUserInfoDTO) error {
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ErrorOnValidate(err)
	}

	uid := dto.Id
	err = userDomain.UpdateUsers(uid, dto)
	if err != nil {
		return common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return nil
}

// 验证用户邮箱
func (service *UserServiceV1) ValidateEmail(dto dtos.ValidateEmailDTO) (bool, error) {
	var err error
	var pass bool
	var verifyDomain *verify.VerifyDomain
	var userDomain *user.UserDomain
	verifyDomain = verify.NewVerifyDomain()
	userDomain = user.NewUserDomain()
	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 从cache拿出保存的邮箱验证码
	pass, err = verifyDomain.VerifyEmail(dto.Id, dto.VerifyCode)
	if err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if pass {
		// 设置email_verify字段
		if err = userDomain.SetEmailVerify(dto.Id, 1); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}
		return true, nil
	} else {
		return false, common.ErrorOnVerify("Email Verify Code Error!")
	}
}

// 验证手机号码
func (service *UserServiceV1) ValidatePhone(dto dtos.ValidatePhoneDTO) (bool, error) {
	var err error
	var pass bool
	var verifyDomain *verify.VerifyDomain
	var userDomain *user.UserDomain
	verifyDomain = verify.NewVerifyDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	// 从cache拿出保存的短信验证码
	if pass, err = verifyDomain.VerifyPhone(dto.Id, dto.VerifyCode); err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if pass {
		// 设置phone_verify字段
		if err = userDomain.SetPhoneVerify(dto.Id, 1); err != nil {
			return false, common.ErrorOnServerInner(err, userDomain.DomainName())
		}
		return true, nil
	} else {
		return false, common.ErrorOnVerify("Sms Verify Code Error!")
	}
}

// 设置用户账号状态
func (service *UserServiceV1) SetStatus(dto dtos.SetStatusDTO) (int, error) {
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return -1, common.ErrorOnValidate(err)
	}

	if err = userDomain.SetStatus(dto.Id, dto.Status); err != nil {
		return -1, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return 0, nil
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
	if userDTO, isPass, err = userDomain.VerifyPassword(dto.Id, dto.Old); isPass && userDTO != nil && err == nil {
		return true, nil
	} else if err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	} else if userDTO == nil {
		return false, common.ErrorOnVerify("Account Not Exist!")
	} else if !isPass {
		// 校验密码失败
		return false, common.ErrorOnVerify("Password Error!")
	}

	// 设置新密码
	if err = userDomain.ReSetPassword(dto.Id, dto.New); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return false, nil
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
	isExist, err = userDomain.IsUserExist(dto.Id)
	if err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}
	if !isExist {
		return false, common.ErrorOnVerify("Account Not Exist!")
	}

	// 校验Code
	isVerify, err = verifyDomain.VerifyResetPasswordCode(dto.Id, dto.Code)
	if err != nil {
		return false, common.ErrorOnServerInner(err, verifyDomain.DomainName())
	}

	if !isVerify {
		return false, common.ErrorOnVerify("Reset Password Fail,Please Retry!")
	}

	// 重设密码
	if err = userDomain.ReSetPassword(dto.Id, dto.New); err != nil {
		return false, common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	return true, nil
}

// TODO 上传用户头像
func (service *UserServiceV1) UploadAvatar() error {

	return nil
}

// qq oauth 鉴权
func (service *UserServiceV1) QQOAuthLogin(dto dtos.QQLoginDTO) (string, error) {
	var err error
	var token string
	var isQQBinding bool
	var qqOauthAccountInfo *XOAuth.OAuthAccountInfo // qq账号鉴权信息
	var userOAuthsInfo *dtos.UserOAuthInfoDTO       // 创建用户后的信息

	var thirdOauthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOauthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	if qqOauthAccountInfo, err = thirdOauthDomain.GetQQOauthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOauthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isQQBinding, err = userDomain.IsQQAccountBinding(qqOauthAccountInfo.OpenId, qqOauthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isQQBinding {
		// 未绑定，进入创建用户流程
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.QQOauthPlatform, qqOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.Id,
				userOAuthsInfo.No,
				userOAuthsInfo.Name,
				userOAuthsInfo.Avatar)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthsInfo, err = userDomain.GetUserOauths(user.QQOauthPlatform, qqOauthAccountInfo.OpenId, qqOauthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthsInfo.Id, userOAuthsInfo.No, userOAuthsInfo.Name, userOAuthsInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}

// wechat Oauth 鉴权
func (service *UserServiceV1) WechatOAuthLogin(dto dtos.WechatLoginDTO) (string, error) {
	var err error
	var token string
	var isWechatBinding bool
	var wechatOauthAccountInfo *XOAuth.OAuthAccountInfo // 微信账号鉴权信息
	var userOAuthsInfo *dtos.UserOAuthInfoDTO           // 创建用户后的信息

	var thirdOauthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOauthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用wechat回调授权码code开始鉴权流程并获取微信用户信息
	if wechatOauthAccountInfo, err = thirdOauthDomain.GetWechatOauthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOauthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isWechatBinding, err = userDomain.IsWechatAccountBinding(wechatOauthAccountInfo.OpenId, wechatOauthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isWechatBinding {
		// 未绑定，进入创建用户流程
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.WechatOauthPlatform, wechatOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.Id,
				userOAuthsInfo.No,
				userOAuthsInfo.Name,
				userOAuthsInfo.Avatar)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthsInfo, err = userDomain.GetUserOauths(user.WechatOauthPlatform, wechatOauthAccountInfo.OpenId, wechatOauthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthsInfo.Id, userOAuthsInfo.No, userOAuthsInfo.Name, userOAuthsInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}

// 微博 Oauth 鉴权
func (service *UserServiceV1) WeiboOAuthLogin(dto dtos.WeiboLoginDTO) (string, error) {
	var err error
	var token string
	var isWeiboBinding bool
	var weiboOauthAccountInfo *XOAuth.OAuthAccountInfo // 微博账号鉴权信息
	var userOAuthsInfo *dtos.UserOAuthInfoDTO          // 创建用户后的信息

	var thirdOauthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOauthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用weibo回调授权码code开始鉴权流程并获取微博用户信息
	if weiboOauthAccountInfo, err = thirdOauthDomain.GetQQOauthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOauthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isWeiboBinding, err = userDomain.IsQQAccountBinding(weiboOauthAccountInfo.OpenId, weiboOauthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isWeiboBinding {
		// 未绑定，进入创建用户流程
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.WeiboOauthPlatform, weiboOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.Id,
				userOAuthsInfo.No,
				userOAuthsInfo.Name,
				userOAuthsInfo.Avatar)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthsInfo, err = userDomain.GetUserOauths(user.WeiboOauthPlatform, weiboOauthAccountInfo.OpenId, weiboOauthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthsInfo.Id, userOAuthsInfo.No, userOAuthsInfo.Name, userOAuthsInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}
