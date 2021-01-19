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
	IsEmailAccountExist(dto dtos.IsEmailAccountExistDTO) (bool, error) // 邮箱账号是否存在
	IsPhoneAccountExist(dto dtos.IsPhoneAccountExistDTO) (bool, error) // 手机账号是否存在

	CreateUserWithEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UserInfoDTO, error) // 创建邮箱账号
	CreateUserWithPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UserInfoDTO, error) // 创建手机号码账号

	EmailAuth(dto dtos.AuthWithEmailPasswordDTO) (string, error) // 邮箱账号鉴权
	PhoneAuth(dto dtos.AuthWithPhonePasswordDTO) (string, error) // 手机号码鉴权

	RemoveToken(dto dtos.RemoveTokenDTO) error // 移除登录token缓存

	GetUserInfo(dto dtos.GetUserInfoDTO) (*dtos.UserInfoDTO, error)    // 获取用户信息
	SetUserInfos(dto dtos.SetUserInfoDTO) (bool, error)                // 修改用户信息
	EmailValidate(dto dtos.EmailValidateDTO) (bool, error)             // 校验邮箱地址，验证邮箱链接
	EmailBinding(dto dtos.EmailValidateDTO) (bool, error)              // 绑定邮箱，验证邮箱链接
	PhoneBinding(dto dtos.PhoneValidateDTO) (bool, error)              // 绑定手机，验证短信验证码
	ModifiedPassword(dto dtos.ModifiedPasswordDTO) (bool, error)       // 更改用户密码
	ResetForgetPassword(dto dtos.ResetForgetPasswordDTO) (bool, error) // 忘记密码重设
	SetAvatarUri(dto dtos.SetAvatarUriDTO) (bool, error)               // 上传头像

}
