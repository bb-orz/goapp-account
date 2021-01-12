package third

import (
	"fmt"
	"github.com/bb-orz/goinfras/XOAuth"
	"goapp/common"
)

/*
Oauth 领域层：实现第三方平台鉴权相关具体业务逻辑，主要为通过accessCode获取用户在第三方平台账号的信息
*/
type ThirdOAuthDomain struct{}

func NewThirdOAuthDomain() *ThirdOAuthDomain {
	domain := new(ThirdOAuthDomain)
	return domain
}

func (domain *ThirdOAuthDomain) DomainName() string {
	return "OauthDomain"

}

// 通过accessCode获取qq user info
func (domain *ThirdOAuthDomain) GetQQOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XQQOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.DomainInnerErrorOnThirdPartRequest(oAuthResult.Error, fmt.Sprintf("QQ OAuth AccessCode:%s", accessCode))
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取wechat user info
func (domain *ThirdOAuthDomain) GetWechatOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XWechatOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.DomainInnerErrorOnThirdPartRequest(oAuthResult.Error, fmt.Sprintf("Wechat OAuth AccessCode:%s", accessCode))
	}

	return oAuthResult.UserInfo, nil
}

// 通过accessCode获取weibo user info
func (domain *ThirdOAuthDomain) GetWeiboOauthUserInfo(accessCode string) (*XOAuth.OAuthAccountInfo, error) {
	var oAuthResult XOAuth.OAuthResult
	oAuthResult = XOAuth.XWeiboOAuthManager().Authorize(accessCode)

	if oAuthResult.Error != nil || !oAuthResult.Result {
		return nil, common.DomainInnerErrorOnThirdPartRequest(oAuthResult.Error, fmt.Sprintf("Weibo OAuth AccessCode:%s", accessCode))
	}

	return oAuthResult.UserInfo, nil
}
