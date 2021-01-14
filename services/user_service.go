package services

import (
	"goapp/dtos"
)

/* 定义用户模块的服务层方法，并定义数据传输对象DTO*/
var userService IUserService

// 用于对外暴露账户应用服务，唯一的暴露点，供接口层调用
func GetUserService() IUserService {
	return userService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetUserService(service IUserService) {
	userService = service
}

// 定义用户服务接口
type IUserService interface {
	CreateUserWithEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UsersDTO, error) // 创建邮箱账号
	CreateUserWithPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UsersDTO, error) // 创建手机号码账号

	EmailAuth(dto dtos.AuthWithEmailPasswordDTO) (string, error) // 邮箱账号鉴权
	PhoneAuth(dto dtos.AuthWithPhonePasswordDTO) (string, error) // 手机号码鉴权

	RemoveToken(dto dtos.RemoveTokenDTO) error // 移除登录token缓存

	QQOAuthLogin(dto dtos.QQLoginDTO) (string, error)         // qq三方账号鉴权
	WechatOAuthLogin(dto dtos.WechatLoginDTO) (string, error) // 微信三方账号鉴权
	WeiboOAuthLogin(dto dtos.WeiboLoginDTO) (string, error)   // 微博三方账号鉴权

	GetUserInfo(dto dtos.GetUserInfoDTO) (*dtos.UserInfoDTO, error)    // 获取用户信息
	SetUserInfos(dto dtos.SetUserInfoDTO) error                        // 修改用户信息
	ValidateEmail(dto dtos.ValidateEmailDTO) (bool, error)             // 绑定邮箱，验证邮箱链接
	ValidatePhone(dto dtos.ValidatePhoneDTO) (bool, error)             // 绑定手机，验证短信验证码
	SetStatus(dto dtos.SetStatusDTO) (int, error)                      // 设置用户状态
	ModifiedPassword(dto dtos.ModifiedPasswordDTO) (bool, error)       // 更改用户密码
	ResetForgetPassword(dto dtos.ResetForgetPasswordDTO) (bool, error) // 忘记密码重设
	UploadAvatar() error                                               // 上传头像
}
