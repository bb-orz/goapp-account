package dtos

import (
	"github.com/bb-orz/goinfras/XValidate"
	"time"
)

// OAuthDTO is a mapping object for oauth table in mysql
type OAuthDTO struct {
	Id          uint      `validate:"numeric" json:"id"`            //
	UserId      uint      `validate:"numeric" json:"user_id"`       // user表外键
	Platform    uint      `validate:"numeric" json:"platform"`      // 平台账号类型
	AccessToken string    `validate:"alphanum" json:"access_token"` // 获取三方平台用户信息的accessToken
	OpenId      string    `validate:"alphanum" json:"open_id"`      // 开发者可通过OpenID来获取用户基本信息
	UnionId     string    `validate:"alphanum" json:"union_id"`     // 如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。
	NickName    string    `validate:"alphanum" json:"nick_name"`    // 三方平台账号的用户昵称
	Gender      uint      `validate:"numeric" json:"gender"`        // 用户性别：1男2女
	Avatar      string    `validate:"alphanum" json:"avatar"`       // 用户头像
	CreatedAt   time.Time `validate:"numeric" json:"created_at"`    // 创建或绑定时间
	UpdatedAt   time.Time `validate:"numeric" json:"updated_at"`    // 最近登录时间
	DeletedAt   time.Time `validate:"numeric" json:"deleted_at"`    //
}

func (dto *OAuthDTO) Validate() error {
	return XValidate.V(dto)
}