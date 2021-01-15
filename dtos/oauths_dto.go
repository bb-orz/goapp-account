package dtos

import (
	"github.com/bb-orz/goinfras/XValidate"
	"time"
)

// OauthsDTO is a mapping object for oauths table in mysql
type OauthsDTO struct {
	Id          uint      `validate:"required,numeric" json:"id"`                   //
	UserId      uint      `validate:"required,numeric" json:"user_id"`              // user表外键
	Platform    uint      `validate:"required,numeric" json:"platform"`             // 平台账号类型
	AccessToken string    `validate:"required,alphanumunicode" json:"access_token"` // 获取三方平台用户信息的accessToken
	OpenId      string    `validate:"required,alphanumunicode" json:"open_id"`      // 开发者可通过OpenID来获取用户基本信息
	UnionId     string    `validate:"required,alphanumunicode" json:"union_id"`     // 如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。
	NickName    string    `validate:"required,alphanumunicode" json:"nick_name"`    // 三方平台账号的用户昵称
	Gender      uint      `validate:"required,numeric" json:"gender"`               // 用户性别：1男2女
	Avatar      string    `validate:"required,alphanumunicode" json:"avatar"`       // 用户头像
	CreatedAt   time.Time `validate:"required,numeric" json:"created_at"`           // 创建或绑定时间
	UpdatedAt   time.Time `validate:"required,numeric" json:"updated_at"`           // 最近登录时间
	DeletedAt   time.Time `validate:"required,numeric" json:"deleted_at"`           //
}

func (dto *OauthsDTO) Validate() error {
	return XValidate.V(dto)
}
