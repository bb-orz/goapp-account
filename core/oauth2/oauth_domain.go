package oauth2

import (
	"github.com/bb-orz/goinfras/XOAuth"
	"goinfras-sample-account/common"
)

/*
Oauth 领域层：实现第三方平台鉴权相关具体业务逻辑，主要为通过accessCode获取用户在第三方平台账号的信息
*/
type OauthDomain struct{}

func NewOauthDomain() *OauthDomain {
	domain := new(OauthDomain)
	return domain
}

func (domain *OauthDomain)DomainName() string {
	return "OauthDomain"

}

// 通过accessCode获取qq user info
func (domain *OauthDomain) GetQQOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XQQOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.WrapError(oAuthResult.Error, common.ErrorFormatDomainThirdPart, "QQ.Authorize")
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取wechat user info
func (domain *OauthDomain) GetWechatOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XWechatOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.WrapError(oAuthResult.Error, common.ErrorFormatDomainThirdPart, "Wechat.Authorize")
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取weibo user info
func (domain *OauthDomain) GetWeiboOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XWeiboOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.WrapError(oAuthResult.Error, common.ErrorFormatDomainThirdPart, "Weibo.Authorize")
	}

	return oAuthResult.UserInfo, nil
}
