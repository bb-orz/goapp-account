package dtos

import (
	"time"
)

// UserDTO is a mapping object for user table in mysql
type UserDTO struct {
	Id            uint      `json:"id"`             // 用户id
	No            string    `json:"no"`             // 用户生成编号
	Name          string    `json:"name"`           // 用户名
	Age           uint      `json:"age"`            // 用户年龄
	Gender        uint      `json:"gender"`         // 用户性别
	Avatar        string    `json:"avatar"`         // 用户头像
	Email         string    `json:"email"`          // 账户邮箱
	EmailVerified uint      `json:"email_verified"` // 邮箱是否已验证
	Phone         string    `json:"phone"`          // 账户手机号码
	PhoneVerified uint      `json:"phone_verified"` // 手机号码是否已验证
	Password      string    `json:"password"`       // 用户已加密密码字符串
	Salt          string    `json:"salt"`           // 加密盐
	Status        uint      `json:"status"`         // 账户状态：1：启用，0：停用
	CreatedAt     time.Time `json:"created_at"`     //
	UpdatedAt     time.Time `json:"updated_at"`     //
	DeletedAt     time.Time `json:"deleted_at"`     //
}

func (dto *UserDTO) ToInfoDTO() *UserInfoDTO {
	return &UserInfoDTO{
		Id:            dto.Id,
		No:            dto.No,
		Name:          dto.Name,
		Age:           dto.Age,
		Gender:        dto.Gender,
		Avatar:        dto.Avatar,
		Email:         dto.Email,
		EmailVerified: dto.EmailVerified,
		Phone:         dto.Phone,
		PhoneVerified: dto.PhoneVerified,
		Status:        dto.Status,
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
		DeletedAt:     dto.DeletedAt,
	}
}

// 用户数据传输对象
type UserInfoDTO struct {
	Id            uint      `json:"id"`             // 用户id
	No            string    `json:"no"`             // 用户生成编号
	Name          string    `json:"name"`           // 用户名
	Age           uint      `json:"age"`            // 用户年龄
	Gender        uint      `json:"gender"`         // 用户性别
	Avatar        string    `json:"avatar"`         // 用户头像
	Email         string    `json:"email"`          // 账户邮箱
	EmailVerified uint      `json:"email_verified"` // 邮箱是否已验证
	Phone         string    `json:"phone"`          // 账户手机号码
	PhoneVerified uint      `json:"phone_verified"` // 手机号码是否已验证
	Status        uint      `json:"status"`         // 账户状态：1：启用，0：停用
	CreatedAt     time.Time `json:"created_at"`     //
	UpdatedAt     time.Time `json:"updated_at"`     //
	DeletedAt     time.Time `json:"deleted_at"`     //
}

type UserOAuthInfoDTO struct {
	Id            uint       `json:"id"`             // 用户id
	No            string     `json:"no"`             // 用户生成编号
	Name          string     `json:"name"`           // 用户名
	Age           uint       `json:"age"`            // 用户年龄
	Gender        uint       `json:"gender"`         // 用户性别
	Avatar        string     `json:"avatar"`         // 用户头像
	Email         string     `json:"email"`          // 账户邮箱
	EmailVerified uint       `json:"email_verified"` // 邮箱是否已验证
	Phone         string     `json:"phone"`          // 账户手机号码
	PhoneVerified uint       `json:"phone_verified"` // 手机号码是否已验证
	Status        uint       `json:"status"`         // 账户状态：1：启用，0：停用
	CreatedAt     time.Time  `json:"created_at"`     //
	UpdatedAt     time.Time  `json:"updated_at"`     //
	DeletedAt     time.Time  `json:"deleted_at"`     //
	OAuth         []OAuthDTO `json:"oauth_list"`
}
