package services

import (
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XValidate"
	"goapp/common"
	"goapp/core/third"
	"goapp/core/user"
	"goapp/dtos"
	"sync"
)

var _ IThirdOAuthService = new(ThirdOAuthServiceV1)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		thirdOAuthService := new(ThirdOAuthServiceV1)
		SetThirdOAuthService(thirdOAuthService)
	})
}

// 用户服务实例 V1
type ThirdOAuthServiceV1 struct{}

// qq oauth 鉴权
func (service *ThirdOAuthServiceV1) QQOAuthLogin(dto dtos.QQLoginDTO) (string, error) {
	var err error
	var token string
	var isQQBinding bool
	var insertId int64
	var qqOAuthAccountInfo *XOAuth.OAuthAccountInfo // qq账号鉴权信息
	var userOAuthInfo *dtos.UserOAuthInfoDTO

	var thirdOAuthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOAuthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	if qqOAuthAccountInfo, err = thirdOAuthDomain.GetQQOAuthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isQQBinding, err = userDomain.IsQQAccountBinding(qqOAuthAccountInfo.OpenId, qqOAuthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isQQBinding {
		// 未绑定，进入创建用户流程
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.QQOAuthPlatform, qqOAuthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				qqOAuthAccountInfo.NickName,
				qqOAuthAccountInfo.AvatarUrl,
			)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthInfo, err = userDomain.GetUserOAuth(user.QQOAuthPlatform, qqOAuthAccountInfo.OpenId, qqOAuthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthInfo.Id, userOAuthInfo.No, userOAuthInfo.Name, userOAuthInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}

// wechat OAuth 鉴权
func (service *ThirdOAuthServiceV1) WechatOAuthLogin(dto dtos.WechatLoginDTO) (string, error) {
	var err error
	var token string
	var isWechatBinding bool
	var insertId int64
	var wechatOAuthAccountInfo *XOAuth.OAuthAccountInfo // 微信账号鉴权信息
	var userOAuthInfo *dtos.UserOAuthInfoDTO            // 创建用户后的信息

	var thirdOAuthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOAuthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用wechat回调授权码code开始鉴权流程并获取微信用户信息
	if wechatOAuthAccountInfo, err = thirdOAuthDomain.GetWechatOAuthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isWechatBinding, err = userDomain.IsWechatAccountBinding(wechatOAuthAccountInfo.OpenId, wechatOAuthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isWechatBinding {
		// 未绑定，进入创建用户流程
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.WechatOAuthPlatform, wechatOAuthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				wechatOAuthAccountInfo.NickName,
				wechatOAuthAccountInfo.AvatarUrl,
			)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthInfo, err = userDomain.GetUserOAuth(user.WechatOAuthPlatform, wechatOAuthAccountInfo.OpenId, wechatOAuthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthInfo.Id, userOAuthInfo.No, userOAuthInfo.Name, userOAuthInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}

// 微博 OAuth 鉴权
func (service *ThirdOAuthServiceV1) WeiboOAuthLogin(dto dtos.WeiboLoginDTO) (string, error) {
	var err error
	var token string
	var isWeiboBinding bool
	var insertId int64

	var weiboOAuthAccountInfo *XOAuth.OAuthAccountInfo // 微博账号鉴权信息
	var userOAuthInfo *dtos.UserOAuthInfoDTO           // 创建用户后的信息

	var thirdOAuthDomain *third.ThirdOAuthDomain
	var userDomain *user.UserDomain
	thirdOAuthDomain = third.NewThirdOAuthDomain()
	userDomain = user.NewUserDomain()

	// 校验传输参数
	if err = XValidate.V(dto); err != nil {
		return "", common.ErrorOnValidate(err)
	}

	// third oauth domain：使用weibo回调授权码code开始鉴权流程并获取微博用户信息
	if weiboOAuthAccountInfo, err = thirdOAuthDomain.GetQQOAuthUserInfo(dto.AccessCode); err != nil {
		return "", common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	// user domain: 使用OpenId UnionId查找 oauth表查看用户是否存在
	if isWeiboBinding, err = userDomain.IsQQAccountBinding(weiboOAuthAccountInfo.OpenId, weiboOAuthAccountInfo.UnionId); err != nil {
		return "", common.ErrorOnServerInner(err, userDomain.DomainName())
	}

	if !isWeiboBinding {
		// 未绑定，进入创建用户流程
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.WeiboOAuthPlatform, weiboOAuthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				weiboOAuthAccountInfo.NickName,
				weiboOAuthAccountInfo.AvatarUrl,
			)
			if err != nil {
				return "", common.ErrorOnServerInner(err, userDomain.DomainName())
			}
			return token, nil
		} else {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

	} else {
		// 已绑定，获取用户信息，进入登录流程
		if userOAuthInfo, err = userDomain.GetUserOAuth(user.WeiboOAuthPlatform, weiboOAuthAccountInfo.OpenId, weiboOAuthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthInfo.Id, userOAuthInfo.No, userOAuthInfo.Name, userOAuthInfo.Avatar); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}
	}

	return token, nil
}

// QQ账号绑定
func (service *ThirdOAuthServiceV1) QQOAuthBinding(dto dtos.QQBindingDTO) (bool, error) {
	var err error
	var insertId int64
	var qqBindingDTO dtos.QQBindingDTO
	var oAuthAccountInfo *XOAuth.OAuthAccountInfo

	// 校验传输参数
	if err = XValidate.V(qqBindingDTO); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	thirdOAuthDomain := third.NewThirdOAuthDomain()
	if oAuthAccountInfo, err = thirdOAuthDomain.GetQQOAuthUserInfo(qqBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.QQOAuthPlatform, oAuthAccountInfo); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	if insertId >= 0 {
		return true, nil
	}

	return false, nil
}

// 微信账号绑定
func (service *ThirdOAuthServiceV1) WechatOAuthBinding(dto dtos.WechatBindingDTO) (bool, error) {
	var err error
	var insertId int64
	var wechatBindingDTO dtos.WechatBindingDTO
	var oAuthAccountInfo *XOAuth.OAuthAccountInfo

	// 校验传输参数
	if err = XValidate.V(wechatBindingDTO); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	thirdOAuthDomain := third.NewThirdOAuthDomain()
	if oAuthAccountInfo, err = thirdOAuthDomain.GetWechatOAuthUserInfo(wechatBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.WechatOAuthPlatform, oAuthAccountInfo); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	if insertId >= 0 {
		return true, nil
	}

	return false, nil
}

// 微博账户绑定
func (service *ThirdOAuthServiceV1) WeiboOAuthBinding(dto dtos.WeiboBindingDTO) (bool, error) {
	var err error
	var insertId int64
	var weiboBindingDTO dtos.WeiboBindingDTO
	var oAuthAccountInfo *XOAuth.OAuthAccountInfo

	// 校验传输参数
	if err = XValidate.V(weiboBindingDTO); err != nil {
		return false, common.ErrorOnValidate(err)
	}

	thirdOAuthDomain := third.NewThirdOAuthDomain()
	if oAuthAccountInfo, err = thirdOAuthDomain.GetWeiboOAuthUserInfo(weiboBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.WeiboOAuthPlatform, oAuthAccountInfo); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	if insertId >= 0 {
		return true, nil
	}

	return false, nil
}
