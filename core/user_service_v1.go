package core

import (
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XValidate"
	"goinfras-sample-account/common"
	"goinfras-sample-account/core/oauth2"
	"goinfras-sample-account/core/user"
	"goinfras-sample-account/core/verified"
	"goinfras-sample-account/dtos"
	"goinfras-sample-account/services"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑
// 接收领域层和dao层的错误包装处理

var _ services.IUserService = new(UserServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		userService := new(UserServiceV1)
		services.SetUserService(userService)
	})
}

// 用户服务实例 V1
type UserServiceV1 struct {}

// 邮箱创建用户账号
func (service *UserServiceV1) CreateUserWithEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UserDTO, error) {
	var err error
	var isExist bool
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ClientErrorOnValidateParameters(err)
	}

	// 验证用户邮箱是否存在
	if isExist, err = userDomain.IsEmailExist(dto.Email); err != nil {
		return nil, common.ServerInnerError(err,userDomain.DomainName())
	} else if isExist {
		return nil, common.ClientErrorOnCheckInformation(err,  "Email Account Exist!")
	}

	userDTO, err = userDomain.CreateUserForEmail(dto)
	if err != nil {
		return nil, common.ServerInnerError(err, userDomain.DomainName())
	}
	return userDTO, nil
}

// 手机号码创建用户账号
func (service *UserServiceV1) CreateUserWithPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UserDTO, error) {
	var err error
	var isExist bool
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ClientErrorOnValidateParameters(err)
	}

	// 验证用户手机号码是否存在
	if isExist, err = userDomain.IsPhoneExist(dto.Phone); err != nil {
		return nil, common.ServerInnerError(err,userDomain.DomainName())
	} else if isExist {
		return nil, common.ClientErrorOnCheckInformation(err, "Phone Account Exist!")
	}

	userDTO, err = userDomain.CreateUserForPhone(dto)
	if err != nil {
		return nil, common.ServerInnerError(err, userDomain.DomainName())
	}
	return userDTO, nil
}

// 邮箱账号登录鉴权
func (service *UserServiceV1) EmailAuth(dto dtos.AuthWithEmailPasswordDTO) (string, error) {
	var err error
	var token string
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ClientErrorOnValidateParameters(err)
	}

	// 查找邮件账号是否存在
	if userDTO, err = userDomain.GetUserInfoByEmail(dto.Email); err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}
	if userDTO == nil {
		return "", common.ClientErrorOnCheckInformation(err, "Email Account Not Exist!")
	} else if !XGlobal.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码失败
		return "", common.ClientErrorOnCheckInformation(err, "Password Error!")
	}

	// JWT token
	token, err = userDomain.GenToken(userDTO.No, userDTO.Name, userDTO.Avatar)
	if err != nil {
		return "", common.ServerInnerError(err,userDomain.DomainName())
	}

	return token, nil
}

// 手机账号登录鉴权
func (service *UserServiceV1) PhoneAuth(dto dtos.AuthWithPhonePasswordDTO) (string, error) {
	var err error
	var token string
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()


	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ClientErrorOnValidateParameters(err)
	}

	// 查找手机账号是否存在
	userDTO, err = userDomain.GetUserInfoByPhone(dto.Phone)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	if userDTO == nil {
		return "", common.ClientErrorOnCheckInformation(err, "Phone Account Not Exist!")
	} else if !XGlobal.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码失败
		return "", common.ClientErrorOnCheckInformation(err, "Password Error!")
	}

	// JWT token
	token, err = userDomain.GenToken(userDTO.No, userDTO.Name, userDTO.Avatar)
	if err != nil {
		return "", common.ServerInnerError(err, user.DomainName)
	}
	return token, nil
}

// 获取用户信息
func (service *UserServiceV1) GetUserInfo(dto dtos.GetUserInfoDTO) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return nil, common.ClientErrorOnValidateParameters(err)
	}

	// 查找用户信息
	userDTO, err = userDomain.GetUserInfo(dto.ID)
	if err != nil {
		return nil, common.ServerInnerError(err, userDomain.DomainName())
	}

	return userDTO, nil
}

// 批量设置用户信息
func (service *UserServiceV1) SetUserInfos(dto dtos.SetUserInfoDTO) error {
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ClientErrorOnValidateParameters(err)
	}

	uid := dto.ID
	err = userDomain.SetUserInfos(uid, dto)
	if err != nil {
		return common.ServerInnerError(err, userDomain.DomainName())
	}

	return nil
}

// 验证用户邮箱
func (service *UserServiceV1) ValidateEmail(dto dtos.ValidateEmailDTO) (bool, error) {
	var err error
	var pass bool
	var verifiedDomain *verified.VerifiedDomain
	verifiedDomain = verified.NewVerifiedDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ClientErrorOnValidateParameters(err)
	}

	// 从cache拿出保存的邮箱验证码
	pass, err = verifiedDomain.VerifiedEmail(dto.ID, dto.VerifiedCode)
	if err != nil {
		return false, common.ServerInnerError(err, verifiedDomain.DomainName())
	}

	if pass {
		return true, nil
	}else {
		return false, common.ClientErrorOnCheckInformation(err,"Email Verified Code Error!")
	}
}

// 验证手机号码
func (service *UserServiceV1) ValidatePhone(dto dtos.ValidatePhoneDTO) (bool, error) {
	var err error
	var pass bool
	var verifiedDomain *verified.VerifiedDomain
	verifiedDomain = verified.NewVerifiedDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return false, common.ClientErrorOnValidateParameters(err)
	}

	// 从cache拿出保存的短信验证码
	pass, err = verifiedDomain.VerifiedPhone(dto.ID, dto.VerifiedCode)
	if err != nil {
		return false, common.ServerInnerError(err, verifiedDomain.DomainName())
	}

	if pass {
		return true, nil
	}else {
		return false, common.ClientErrorOnCheckInformation(err,"Sms Verified Code Error!")
	}
}

// 设置用户账号状态
func (service *UserServiceV1) SetStatus(dto dtos.SetStatusDTO) (int, error) {
	var err error
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return -1, common.ClientErrorOnValidateParameters(err)
	}

	err = userDomain.SetStatus(dto.ID, dto.Status)
	if err != nil {
		return -1, common.ServerInnerError(err, userDomain.DomainName())
	}

	return 0, nil
}

// 修改用户密码
func (service *UserServiceV1) ChangePassword(dto dtos.ChangePasswordDTO) error {
	var err error
	var userDTO *dtos.UserDTO
	var userDomain *user.UserDomain
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ClientErrorOnValidateParameters(err)
	}

	// 查找账号是否存在
	userDTO, err = userDomain.GetUserInfo(dto.ID)
	if err != nil {
		return common.ServerInnerError(err, userDomain.DomainName())
	}

	// 校验旧密码
	if userDTO == nil {
		return common.ClientErrorOnCheckInformation(err, "Account Not Exist!")
	} else if !XGlobal.ValidatePassword(dto.Old, userDTO.Salt, userDTO.Password) {
		// 校验旧密码失败
		return common.ClientErrorOnCheckInformation(err, "Old Password Is Wrong!")
	}

	// 设置新密码
	if err = userDomain.ReSetPassword(dto.ID, dto.New); err != nil {
		return common.ServerInnerError(err, userDomain.DomainName())
	}

	return nil
}

// 忘记密码重设
func (service *UserServiceV1) ForgetPassword(dto dtos.ForgetPasswordDTO) error {
	var err error
	var isExist bool
	var isVerified bool
	var userDomain *user.UserDomain
	var verifiedDomain *verified.VerifiedDomain
	userDomain = user.NewUserDomain()
	verifiedDomain = verified.NewVerifiedDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return common.ClientErrorOnValidateParameters(err)
	}

	// 查找账号是否存在
	isExist, err = userDomain.IsUserExist(dto.ID)
	if err != nil {
		return common.ServerInnerError(err, userDomain.DomainName())
	}
	if !isExist {
		return common.ClientErrorOnCheckInformation(err, "Account Not Exist!")
	}

	// 校验Code
	isVerified, err = verifiedDomain.VerifiedResetPasswordCode(dto.ID, dto.Code)
	if err != nil {
		return common.ServerInnerError(err, verifiedDomain.DomainName())
	}

	if !isVerified {
		return common.ClientErrorOnCheckInformation(nil, "Reset Password Fail,Please Retry!")
	}

	return nil

}

// 上传用户头像
func (service *UserServiceV1) UploadAvatar() error {

	return nil
}

// qq oauth 鉴权
func (service *UserServiceV1) QQOAuth(dto dtos.QQLoginDTO) (string, error) {
	var err error
	var token string
	var qqOauthAccountInfo *XOAuth.OAuthAccountInfo // qq账号鉴权信息
	var findUserBindingDTO *dtos.UserOAuthsDTO  // 查找绑定用户
	var userOAuthsInfo *dtos.UserOAuthsDTO      // 创建用户后的信息

	var oauthDomain *oauth2.OauthDomain
	var userDomain *user.UserDomain
	oauthDomain = oauth2.NewOauthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ClientErrorOnValidateParameters(err)
	}

	// oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	qqOauthAccountInfo, err = oauthDomain.GetQQOauthUserInfo(dto.AccessCode)
	if err != nil {
		return "", common.ServerInnerError(err, oauthDomain.DomainName())
	}

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在
	findUserBindingDTO, err = userDomain.GetUserOauths(user.QQOauthPlatform, qqOauthAccountInfo.OpenId, qqOauthAccountInfo.UnionId)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	// 如不存在进入创建用户流程,否则进登录流程
	if findUserBindingDTO == nil {
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.QQOauthPlatform, qqOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.User.No,
				userOAuthsInfo.User.Name,
				userOAuthsInfo.User.Avatar)
			if err != nil {
				return "", common.ServerInnerError(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ServerInnerError(err,userDomain.DomainName())
		}
	}

	// 跳过创建，直接返回token，登录成功
	token, err = userDomain.GenToken(
		findUserBindingDTO.User.No,
		findUserBindingDTO.User.Name,
		findUserBindingDTO.User.Avatar)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	return token, nil
}

// wechat Oauth 鉴权
func (service *UserServiceV1) WechatOAuth(dto dtos.WechatLoginDTO) (string, error) {
	var err error
	var token string
	var wechatOauthAccountInfo *XOAuth.OAuthAccountInfo // 微信账号鉴权信息
	var findUserBindingDTO *dtos.UserOAuthsDTO      // 查找绑定用户
	var userOAuthsInfo *dtos.UserOAuthsDTO          // 创建用户后的信息

	var oauthDomain *oauth2.OauthDomain
	var userDomain *user.UserDomain
	oauthDomain = oauth2.NewOauthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ClientErrorOnValidateParameters(err)
	}

	// oauth domain：使用wechat回调授权码code开始鉴权流程并获取微信用户信息
	wechatOauthAccountInfo, err = oauthDomain.GetQQOauthUserInfo(dto.AccessCode)
	if err != nil {
		return "", common.ServerInnerError(err, oauthDomain.DomainName())
	}

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在
	findUserBindingDTO, err = userDomain.GetUserOauths(user.WechatOauthPlatform, wechatOauthAccountInfo.OpenId, wechatOauthAccountInfo.UnionId)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	// 如不存在进入创建用户流程,否则进登录流程
	if findUserBindingDTO == nil {
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.WechatOauthPlatform, wechatOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.User.No,
				userOAuthsInfo.User.Name,
				userOAuthsInfo.User.Avatar)
			if err != nil {
				return "", common.ServerInnerError(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ServerInnerError(err, userDomain.DomainName())
		}
	}

	// 跳过创建，直接返回token，登录成功
	token, err = userDomain.GenToken(
		findUserBindingDTO.User.No,
		findUserBindingDTO.User.Name,
		findUserBindingDTO.User.Avatar)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	return token, nil
}

// 微博 Oauth 鉴权
func (service *UserServiceV1) WeiboOAuth(dto dtos.WeiboLoginDTO) (string, error) {
	var err error
	var token string
	var weiboOauthAccountInfo *XOAuth.OAuthAccountInfo // 微博账号鉴权信息
	var findUserBindingDTO *dtos.UserOAuthsDTO     // 查找绑定用户
	var userOAuthsInfo *dtos.UserOAuthsDTO         // 创建用户后的信息

	var oauthDomain *oauth2.OauthDomain
	var userDomain *user.UserDomain
	oauthDomain = oauth2.NewOauthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ClientErrorOnValidateParameters(err)
	}

	// oauth domain：使用wechat回调授权码code开始鉴权流程并获取微信用户信息
	weiboOauthAccountInfo, err = oauthDomain.GetQQOauthUserInfo(dto.AccessCode)
	if err != nil {
		return "", common.ServerInnerError(err, oauthDomain.DomainName())
	}

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在
	findUserBindingDTO, err = userDomain.GetUserOauths(user.WeiboOauthPlatform, weiboOauthAccountInfo.OpenId, weiboOauthAccountInfo.UnionId)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	// 如不存在进入创建用户流程,否则进登录流程
	if findUserBindingDTO == nil {
		userOAuthsInfo, err = userDomain.CreateUserOAuthBinding(user.WeiboOauthPlatform, weiboOauthAccountInfo)
		// JWT token
		if userOAuthsInfo != nil {
			token, err = userDomain.GenToken(
				userOAuthsInfo.User.No,
				userOAuthsInfo.User.Name,
				userOAuthsInfo.User.Avatar)
			if err != nil {
				return "", common.ServerInnerError(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ServerInnerError(err, userDomain.DomainName())
		}
	}

	// 跳过创建，直接返回token，登录成功
	token, err = userDomain.GenToken(
		findUserBindingDTO.User.No,
		findUserBindingDTO.User.Name,
		findUserBindingDTO.User.Avatar)
	if err != nil {
		return "", common.ServerInnerError(err, userDomain.DomainName())
	}

	return token, nil
}
