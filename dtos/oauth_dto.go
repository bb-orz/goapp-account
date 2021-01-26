package dtos

import (
	"github.com/bb-orz/goinfras/XValidate"
	"time"
)

// OAuthDTO is a mapping object for oauths table in mysql
type OAuthDTO struct {
	Id          uint      `json:"id"`           //
	UserId      uint      `json:"user_id"`      // user表外键
	Platform    uint      `json:"platform"`     // 平台账号类型
	AccessToken string    `json:"access_token"` // 获取三方平台用户信息的accessToken
	OpenId      string    `json:"open_id"`      // 开发者可通过OpenID来获取用户基本信息
	UnionId     string    `json:"union_id"`     // 如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。
	NickName    string    `json:"nick_name"`    // 三方平台账号的用户昵称
	Gender      uint      `json:"gender"`       // 用户性别：1男2女
	Avatar      string    `json:"avatar"`       // 用户头像
	CreatedAt   time.Time `json:"created_at"`   // 创建或绑定时间
	UpdatedAt   time.Time `json:"updated_at"`   // 最近登录时间
	DeletedAt   time.Time `json:"deleted_at"`   //
}

func (dto *OAuthDTO) Validate() error {
	return XValidate.V(dto)
}
