package dtos

import "time"

/*
以下为服务方法的待校验传输参数对象
*/

// 用户数据传输对象
type UserInfoDTO struct {
	Id            uint      `validate:"required,numeric" json:"id"`             // 用户id
	No            string    `validate:"required,alphanumunicode" json:"no"`     // 用户生成编号
	Name          string    `validate:"required,alphanumunicode" json:"name"`   // 用户名
	Age           int       `validate:"required,numeric" json:"age"`            // 用户年龄
	Gender        int       `validate:"required,numeric" json:"gender"`         // 用户性别
	Avatar        string    `validate:"required,alphanumunicode" json:"avatar"` // 用户头像
	Email         string    `validate:"required,alphanumunicode" json:"email"`  // 账户邮箱
	EmailVerified int       `validate:"required,numeric" json:"email_verified"` // 邮箱是否已验证
	Phone         string    `validate:"required,alphanumunicode" json:"phone"`  // 账户手机号码
	PhoneVerified int       `validate:"required,numeric" json:"phone_verified"` // 手机号码是否已验证
	Status        int       `validate:"required,numeric" json:"status"`         // 账户状态：1：启用，0：停用
	CreatedAt     time.Time `validate:"required,numeric" json:"created_at"`     //
	UpdatedAt     time.Time `validate:"required,numeric" json:"updated_at"`     //
	DeletedAt     time.Time `validate:"required,numeric" json:"deleted_at"`     //
}

type UserOAuthInfoDTO struct {
	Id            uint      `validate:"required,numeric" json:"id"`             // 用户id
	No            string    `validate:"required,alphanumunicode" json:"no"`     // 用户生成编号
	Name          string    `validate:"required,alphanumunicode" json:"name"`   // 用户名
	Age           int       `validate:"required,numeric" json:"age"`            // 用户年龄
	Gender        int       `validate:"required,numeric" json:"gender"`         // 用户性别
	Avatar        string    `validate:"required,alphanumunicode" json:"avatar"` // 用户头像
	Email         string    `validate:"required,alphanumunicode" json:"email"`  // 账户邮箱
	EmailVerified int       `validate:"required,numeric" json:"email_verified"` // 邮箱是否已验证
	Phone         string    `validate:"required,alphanumunicode" json:"phone"`  // 账户手机号码
	PhoneVerified int       `validate:"required,numeric" json:"phone_verified"` // 手机号码是否已验证
	Status        int       `validate:"required,numeric" json:"status"`         // 账户状态：1：启用，0：停用
	CreatedAt     time.Time `validate:"required,numeric" json:"created_at"`     //
	UpdatedAt     time.Time `validate:"required,numeric" json:"updated_at"`     //
	DeletedAt     time.Time `validate:"required,numeric" json:"deleted_at"`     //
	OAuths        []OauthsDTO
}

type QQLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

type WechatLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

type WeiboLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

// 创建用户的数据传输对象
type CreateUserWithEmailDTO struct {
	Name       string `validate:"required,alphanum" json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Password   string `validate:"required,alphanumunicode" json:"password"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password" json:"repassword"`
}

// 创建用户的数据传输对象
type CreateUserWithPhoneDTO struct {
	Name       string `validate:"required,alphanum" json:"name"`
	Phone      string `validate:"required,numeric,eq=11" json:"phone"`
	Password   string `validate:"required,alphanumunicode" json:"password"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password" json:"repassword"`
}

// 邮箱密码鉴权数据传输对象
type AuthWithEmailPasswordDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,alphanumunicode" json:"password"`
}

// 手机号密码鉴权数据传输对象
type AuthWithPhonePasswordDTO struct {
	Phone    string `validate:"required,numeric,eq=11" json:"phone"`
	Password string `validate:"required,alphanumunicode" json:"password"`
}

// 移除登录鉴权token的缓存
type RemoveTokenDTO struct {
	Token string `validate:"required" json:"token"`
}

type GetUserInfoDTO struct {
	ID uint `validate:"required,numeric" json:"id" uri:"id"`
}

// 修改用户新息的数据传输对象
type SetUserInfoDTO struct {
	ID     uint   `validate:"required,numeric" json:"id"`
	Name   string `validate:"alpha" json:"name"`
	Age    uint   `validate:"numeric" json:"age"`
	Avatar string `validate:"alphanumunicode" json:"avatar"`
	Gender uint   `validate:"numeric" json:"gender"`
	Status uint   `validate:"numeric" json:"status"`
}

// 设置用户状态数据传输对象
type SetStatusDTO struct {
	ID     uint `validate:"required,numeric" json:"id"`
	Status uint `validate:"required,numeric" json:"status"` // TODO 验证枚举0/1/2
}

// 验证邮箱数据传输对象
type ValidateEmailDTO struct {
	ID         uint   `validate:"required,numeric" json:"id"`
	VerifyCode string `validate:"required,alphanum" json:"verify_code"`
}

// 验证手机号码数据传输对象
type ValidatePhoneDTO struct {
	ID         uint   `validate:"required,numeric" json:"id"`
	VerifyCode string `validate:"required,alphanum" json:"verify_code"`
}

// 更改密码数据传输对象
type ChangePasswordDTO struct {
	ID    uint   `validate:"required,numeric" json:"id"`
	Old   string `validate:"required,alphanumunicode" json:"old"`
	New   string `validate:"required,alphanumunicode" json:"new"`
	ReNew string `validate:"required,alphanumunicode" json:"renew"`
}

// 忘记密码数据传输对象
type ForgetPasswordDTO struct {
	ID    uint   `validate:"required,numeric" json:"id"`
	Code  string `validate:"required,alphanum" json:"code"` // 允许重设密码的key值，服务端生成后被发往邮箱，用户点击过来后接收
	New   string `validate:"required,alphanumunicode" json:"new"`
	ReNew string `validate:"required,alphanumunicode" json:"renew"`
}