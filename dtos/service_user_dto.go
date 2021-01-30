package dtos

/*
以下为服务方法的待校验传输参数对象
*/

// 验证邮箱数据传输对象
type IsEmailAccountExistDTO struct {
	Email string `validate:"required,email" json:"email" form:"email"`
}

// 验证手机号码数据传输对象
type IsPhoneAccountExistDTO struct {
	Phone string `validate:"required,numeric,len=11" json:"phone" form:"phone"`
}

// 创建用户的数据传输对象
type CreateUserWithEmailDTO struct {
	Name       string `validate:"required,alphanum,max=16,min=4" json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Password   string `validate:"required,alphanum,max=20,min=8" json:"password"`
	RePassword string `validate:"required,alphanum,eqfield=Password,max=20,min=8" json:"repassword"`
}

// 创建用户的数据传输对象
type CreateUserWithPhoneDTO struct {
	Name       string `validate:"required,alphanum,max=16,min=4" json:"name"`
	Phone      string `validate:"required,numeric,len=11" json:"phone"`
	VerifyCode string `validate:"required,alphanum,len=6" json:"verify_code"`
	Password   string `validate:"required,alphanum,max=20,min=8" json:"password"`
	RePassword string `validate:"required,alphanum,max=20,min=8,eqfield=Password" json:"repassword"`
}

// 邮箱密码鉴权数据传输对象
type AuthWithEmailPasswordDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,alphanum,max=20,min=8" json:"password"`
}

// 手机号密码鉴权数据传输对象
type AuthWithPhonePasswordDTO struct {
	Phone      string `validate:"required,numeric,len=11" json:"phone"`
	VerifyCode string `validate:"required,alphanum,len=6" json:"verify_code"`
}

// 移除登录鉴权token的缓存
type RemoveTokenDTO struct {
	Token string `validate:"required" json:"token"`
}

type GetUserInfoDTO struct {
	Id uint `validate:"required,numeric" json:"id" uri:"id"`
}

// 修改用户信息的数据传输对象
type SetUserInfoDTO struct {
	Id     uint   `validate:"required,numeric" json:"id" uri:"id"`
	Name   string `validate:"alphanum,max=16,min=4" json:"name"`
	Age    uint   `validate:"numeric,max=100,min=0" json:"age"`
	Avatar string `validate:"uri" json:"avatar"`
	Gender uint   `validate:"numeric,max=1,min=0" json:"gender"`
}

// 验证邮箱数据传输对象
type EmailValidateDTO struct {
	Id         uint   `validate:"required,numeric" json:"id"`
	Email      string `validate:"required,email" json:"email"`
	VerifyCode string `validate:"required,alphanum,len=6" json:"verify_code"`
}

// 验证手机号码数据传输对象
type PhoneValidateDTO struct {
	Id         uint   `validate:"required,numeric" json:"id"`
	Phone      string `validate:"required,numeric,len=11" json:"phone"`
	VerifyCode string `validate:"required,numeric,len=6" json:"verify_code"`
}

// 更改密码数据传输对象
type ModifiedPasswordDTO struct {
	Id    uint   `validate:"required,numeric" json:"id"`
	Old   string `validate:"required,alphanum,max=20,min=8" json:"old"`
	New   string `validate:"required,alphanum,max=20,min=8" json:"new"`
	ReNew string `validate:"required,alphanum,max=20,min=8,eqfield=New" json:"renew"`
}

// 忘记密码重设数据传输对象
type ResetForgetPasswordDTO struct {
	Email      string `validate:"required,email" json:"email"`
	VerifyCode string `validate:"required,alphanum,len=6" json:"code"` // 允许重设密码的key值，服务端生成后被发往邮箱，用户点击过来后接收
	New        string `validate:"required,alphanum,max=20,min=8" json:"new"`
	ReNew      string `validate:"required,alphanum,max=20,min=8,eqfield=New" json:"renew"`
}

type SetAvatarUriDTO struct {
	Id     uint   `validate:"required,numeric" json:"id"`
	Avatar string `validate:"uri" json:"avatar"`
}
