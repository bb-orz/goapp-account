package dtos

import (
	"github.com/bb-orz/goinfras/XValidate"
	"time"
)

// UserDTO is a mapping object for user table in mysql
type UserDTO struct {
	Id            uint      `validate:"numeric" json:"id"`                  // 用户id
	No            string    `validate:"alphanum" json:"no"`                 // 用户生成编号
	Name          string    `validate:"alphanum" json:"name"`               // 用户名
	Age           uint      `validate:"numeric" json:"age"`                 // 用户年龄
	Gender        uint      `validate:"numeric" json:"gender"`              // 用户性别
	Avatar        string    `validate:"alphanum" json:"avatar"`             // 用户头像
	Email         string    `validate:"alphanum" json:"email"`              // 账户邮箱
	EmailVerified int       `validate:"numeric" json:"email_verified"`      // 邮箱是否已验证
	Phone         string    `validate:"alphanum" json:"phone"`              // 账户手机号码
	PhoneVerified int       `validate:"numeric" json:"phone_verified"`      // 手机号码是否已验证
	Password      string    `validate:"alphanum" json:"password"`           // 用户已加密密码字符串
	Salt          string    `validate:"alphanum" json:"salt"`               // 加密盐
	Status        int       `validate:"numeric" json:"status"`              // 账户状态：1：启用，0：停用
	CreatedAt     time.Time `validate:"numeric" json:"created_at"`          //
	UpdatedAt     time.Time `validate:"numeric" json:"updated_at"`          //
	DeletedAt     time.Time `validate:"required,numeric" json:"deleted_at"` //
}

func (dto *UserDTO) Validate() error {
	return XValidate.V(dto)
}

func (dto *UserDTO) TransToUserInfoDTO() *UserInfoDTO {
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
