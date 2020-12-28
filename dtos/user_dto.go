package dtos

import "time"

// 用户数据传输对象
type UserDTO struct {
	Uid         uint
	No          string
	Name        string
	Age         uint
	Avatar      string
	Gender      uint
	Email       string
	EmailVerify bool
	Phone       string
	PhoneVerify bool
	Password    string
	Salt        string
	Status      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// 用户数据传输对象
type UserInfoDTO struct {
	Uid       uint
	No        string
	Name      string
	Age       uint
	Avatar    string
	Gender    uint
	Email     string
	Phone     string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (d *UserDTO) TransToUserInfoDTO() *UserInfoDTO {
	return &UserInfoDTO{
		Uid:       d.Uid,
		No:        d.No,
		Name:      d.Name,
		Age:       d.Age,
		Avatar:    d.Avatar,
		Gender:    d.Gender,
		Email:     d.Email,
		Phone:     d.Phone,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

// 第三方平台授权账号信息传输对象
type OAuthDTO struct {
	Platform    uint
	UserId      uint
	AccessToken string
	OpenId      string
	UnionId     string
	NickName    string
	Gender      uint
	Avatar      string
}

// 包含第三方账号绑定信息的用户信息传输对象
type UserOAuthsDTO struct {
	User       UserDTO
	UserOAuths []OAuthDTO
}

/*
以下为服务方法的待校验传输参数对象
*/

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
