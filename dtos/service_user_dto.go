package dtos

/*
以下为服务方法的待校验传输参数对象
*/

// 验证邮箱数据传输对象
type IsEmailAccountExistDTO struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

// 验证手机号码数据传输对象
type IsPhoneAccountExistDTO struct {
	Phone string `json:"phone" form:"phone" validate:"required,numeric,len=11"`
}

// 创建用户的数据传输对象
type CreateUserWithEmailDTO struct {
	Name       string `json:"name" validate:"required,alphanum,max=16,min=4"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,alphanum,max=20,min=8"`
	RePassword string `json:"repassword" validate:"required,alphanum,eqfield=Password,max=20,min=8"`
}

// 创建用户的数据传输对象
type CreateUserWithPhoneDTO struct {
	Name       string `json:"name" validate:"required,alphanum,max=16,min=4"`
	Phone      string `json:"phone" validate:"required,numeric,len=11"`
	VerifyCode string `json:"verify_code" validate:"required,alphanum,len=6"`
	Password   string `json:"password" validate:"required,alphanum,max=20,min=8"`
	RePassword string `json:"repassword" validate:"required,alphanum,max=20,min=8,eqfield=Password"`
}

// 邮箱密码鉴权数据传输对象
type AuthWithEmailPasswordDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,max=20,min=8"`
}

// 手机号密码鉴权数据传输对象
type AuthWithPhonePasswordDTO struct {
	Phone      string `json:"phone" validate:"required,numeric,len=11"`
	VerifyCode string `json:"verify_code" validate:"required,alphanum,len=6"`
}

// 移除登录鉴权token的缓存
type RemoveTokenDTO struct {
	Token string `json:"token" validate:"required"`
}

type GetUserInfoDTO struct {
	Id uint `json:"id" uri:"id" validate:"required,numeric"`
}

// 修改用户信息的数据传输对象
type SetUserInfoDTO struct {
	Id     uint   `json:"id" uri:"id" validate:"required,numeric"`
	Name   string `json:"name" validate:"alphanum,max=16,min=4"`
	Age    uint   `json:"age" validate:"numeric,max=100,min=0"`
	Avatar string `json:"avatar" validate:"uri"`
	Gender uint   `json:"gender" validate:"numeric,max=1,min=0"`
}

// 验证邮箱数据传输对象
type EmailValidateDTO struct {
	Id         uint   `json:"id" validate:"required,numeric"`
	Email      string `json:"email" validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,alphanum,len=6"`
}

// 验证手机号码数据传输对象
type PhoneValidateDTO struct {
	Id         uint   `json:"id" validate:"required,numeric"`
	Phone      string `json:"phone" validate:"required,numeric,len=11"`
	VerifyCode string `json:"verify_code" validate:"required,numeric,len=6"`
}

// 更改密码数据传输对象
type ModifiedPasswordDTO struct {
	Id    uint   `json:"id" validate:"required,numeric"`
	Old   string `json:"old" validate:"required,alphanum,max=20,min=8"`
	New   string `json:"new" validate:"required,alphanum,max=20,min=8"`
	ReNew string `json:"renew" validate:"required,alphanum,max=20,min=8,eqfield=New"`
}

// 忘记密码重设数据传输对象
type ResetForgetPasswordDTO struct {
	Email      string `json:"email" validate:"required,email"`
	VerifyCode string `json:"code" validate:"required,alphanum,len=6"` // 允许重设密码的key值，服务端生成后被发往邮箱，用户点击过来后接收
	New        string `json:"new" validate:"required,alphanum,max=20,min=8"`
	ReNew      string `json:"renew" validate:"required,alphanum,max=20,min=8,eqfield=New"`
}

type SetAvatarUriDTO struct {
	Id     uint   `json:"id" validate:"required,numeric"`
	Avatar string `json:"avatar" validate:"uri"`
}
