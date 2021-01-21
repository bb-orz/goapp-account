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
	var qqOauthAccountInfo *XOAuth.OAuthAccountInfo // qq账号鉴权信息
	var userOAuthsInfo *dtos.UserOAuthInfoDTO

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
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.QQOauthPlatform, qqOauthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				qqOauthAccountInfo.NickName,
				qqOauthAccountInfo.AvatarUrl,
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
func (service *ThirdOAuthServiceV1) WechatOAuthLogin(dto dtos.WechatLoginDTO) (string, error) {
	var err error
	var token string
	var isWechatBinding bool
	var insertId int64
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
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.WechatOauthPlatform, wechatOauthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				wechatOauthAccountInfo.NickName,
				wechatOauthAccountInfo.AvatarUrl,
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
func (service *ThirdOAuthServiceV1) WeiboOAuthLogin(dto dtos.WeiboLoginDTO) (string, error) {
	var err error
	var token string
	var isWeiboBinding bool
	var insertId int64

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
		insertId, err = userDomain.CreateUserWithOAuthBinding(user.WeiboOauthPlatform, weiboOauthAccountInfo)
		// JWT token
		if insertId >= 0 {
			token, err = userDomain.GenToken(
				uint(insertId),
				"",
				weiboOauthAccountInfo.NickName,
				weiboOauthAccountInfo.AvatarUrl,
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
		if userOAuthsInfo, err = userDomain.GetUserOauths(user.WeiboOauthPlatform, weiboOauthAccountInfo.OpenId, weiboOauthAccountInfo.UnionId); err != nil {
			return "", common.ErrorOnServerInner(err, userDomain.DomainName())
		}

		if token, err = userDomain.GenToken(userOAuthsInfo.Id, userOAuthsInfo.No, userOAuthsInfo.Name, userOAuthsInfo.Avatar); err != nil {
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
	if oAuthAccountInfo, err = thirdOAuthDomain.GetQQOauthUserInfo(qqBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.QQOauthPlatform, oAuthAccountInfo); err != nil {
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
	if oAuthAccountInfo, err = thirdOAuthDomain.GetWechatOauthUserInfo(wechatBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.WechatOauthPlatform, oAuthAccountInfo); err != nil {
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
	if oAuthAccountInfo, err = thirdOAuthDomain.GetWeiboOauthUserInfo(weiboBindingDTO.AccessCode); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	userDomain := user.NewUserDomain()
	if insertId, err = userDomain.CreateOAuthBinding(user.WeiboOauthPlatform, oAuthAccountInfo); err != nil {
		return false, common.ErrorOnServerInner(err, thirdOAuthDomain.DomainName())
	}

	if insertId >= 0 {
		return true, nil
	}

	return false, nil
}
