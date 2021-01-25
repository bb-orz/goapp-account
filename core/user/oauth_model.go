package user

import (
	"goapp/dtos"
	"gorm.io/gorm"
)

const OAuthTableName = "oauth"

// OAuthModel is a mapping object for oauth table in mysql
type OAuthModel struct {
	gorm.Model
	UserId      uint   `gorm:"user_id" json:"user_id"`           // user表外键
	Platform    uint   `gorm:"platform" json:"platform"`         // 平台账号类型
	AccessToken string `gorm:"access_token" json:"access_token"` // 获取三方平台用户信息的accessToken
	OpenId      string `gorm:"open_id" json:"open_id"`           // 开发者可通过OpenID来获取用户基本信息
	UnionId     string `gorm:"union_id" json:"union_id"`         // 如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。
	NickName    string `gorm:"nick_name" json:"nick_name"`       // 三方平台账号的用户昵称
	Gender      uint   `gorm:"gender" json:"gender"`             // 用户性别：1男2女
	Avatar      string `gorm:"avatar" json:"avatar"`             // 用户头像
}

func NewOAuthModel() *OAuthModel {
	return new(OAuthModel)
}

func (*OAuthModel) TableName() string {
	return OAuthTableName
}

// To DTO
func (m *OAuthModel) ToDTO() *dtos.OAuthDTO {
	return &dtos.OAuthDTO{
		Id:          m.ID,
		UserId:      m.UserId,
		Platform:    m.Platform,
		AccessToken: m.AccessToken,
		OpenId:      m.OpenId,
		UnionId:     m.UnionId,
		NickName:    m.NickName,
		Gender:      m.Gender,
		Avatar:      m.Avatar,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt.Time,
	}
}

// From DTO
func (m *OAuthModel) FromDTO(dto *dtos.OAuthDTO) {
	m.ID = dto.Id
	m.UserId = dto.UserId
	m.Platform = dto.Platform
	m.AccessToken = dto.AccessToken
	m.OpenId = dto.OpenId
	m.UnionId = dto.UnionId
	m.NickName = dto.NickName
	m.Gender = dto.Gender
	m.Avatar = dto.Avatar
	m.CreatedAt = dto.CreatedAt
	m.UpdatedAt = dto.UpdatedAt
}
