package user

import (
	"goapp/dtos"
	"gorm.io/gorm"
)

const UsersTableName = "users"

// UsersModel is a mapping object for users table in mysql
type UsersModel struct {
	gorm.Model
	No            string `gorm:"no" json:"no"`                         // 用户生成编号
	Name          string `gorm:"name" json:"name"`                     // 用户名
	Age           uint   `gorm:"age" json:"age"`                       // 用户年龄
	Gender        uint   `gorm:"gender" json:"gender"`                 // 用户性别
	Avatar        string `gorm:"avatar" json:"avatar"`                 // 用户头像
	Email         string `gorm:"email" json:"email"`                   // 账户邮箱
	EmailVerified int    `gorm:"email_verified" json:"email_verified"` // 邮箱是否已验证
	Phone         string `gorm:"phone" json:"phone"`                   // 账户手机号码
	PhoneVerified int    `gorm:"phone_verified" json:"phone_verified"` // 手机号码是否已验证
	Password      string `gorm:"password" json:"password"`             // 用户已加密密码字符串
	Salt          string `gorm:"salt" json:"salt"`                     // 加密盐
	Status        int    `gorm:"status" json:"status"`                 // 账户状态：1：启用，0：停用
}

func NewUsersModel() *UsersModel {
	return new(UsersModel)
}

func (*UsersModel) TableName() string {
	return UsersTableName
}

// To DTO
func (m *UsersModel) ToDTO() *dtos.UsersDTO {
	return &dtos.UsersDTO{
		Id:            m.ID,
		No:            m.No,
		Name:          m.Name,
		Age:           m.Age,
		Gender:        m.Gender,
		Avatar:        m.Avatar,
		Email:         m.Email,
		EmailVerified: m.EmailVerified,
		Phone:         m.Phone,
		PhoneVerified: m.PhoneVerified,
		Password:      m.Password,
		Salt:          m.Salt,
		Status:        m.Status,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     m.DeletedAt.Time,
	}
}

// From DTO
func (m *UsersModel) FromDTO(dto *dtos.UsersDTO) {
	m.ID = dto.Id
	m.No = dto.No
	m.Name = dto.Name
	m.Age = dto.Age
	m.Gender = dto.Gender
	m.Avatar = dto.Avatar
	m.Email = dto.Email
	m.EmailVerified = dto.EmailVerified
	m.Phone = dto.Phone
	m.PhoneVerified = dto.PhoneVerified
	m.Password = dto.Password
	m.Salt = dto.Salt
	m.Status = dto.Status
}

// From DTO
func (m *UsersModel) FromInfoDTO(dto *dtos.UserInfoDTO) {
	m.ID = dto.Id
	m.No = dto.No
	m.Name = dto.Name
	m.Age = dto.Age
	m.Gender = dto.Gender
	m.Avatar = dto.Avatar
	m.Email = dto.Email
	m.EmailVerified = dto.EmailVerified
	m.Phone = dto.Phone
	m.PhoneVerified = dto.PhoneVerified
	m.Status = dto.Status
}
