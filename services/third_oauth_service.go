package services

import "goapp/dtos"

/* 定义用户模块的服务层方法，并定义数据传输对象DTO*/
var thirdOAuthService IThirdOAuthService

// 用于对外暴露账户应用服务，唯一的暴露点，供接口层调用
func GetThirdOAuthService() IThirdOAuthService {
	return thirdOAuthService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetThirdOAuthService(service IThirdOAuthService) {
	thirdOAuthService = service
}

// 定义用户服务接口
type IThirdOAuthService interface {
	QQOAuthLogin(dto dtos.QQLoginDTO) (string, error)         // qq三方账号鉴权
	WechatOAuthLogin(dto dtos.WechatLoginDTO) (string, error) // 微信三方账号鉴权
	WeiboOAuthLogin(dto dtos.WeiboLoginDTO) (string, error)   // 微博三方账号鉴权

	QQOAuthBinding(dto dtos.QQBindingDTO) (bool, error)         // qq三方账号绑定
	WechatOAuthBinding(dto dtos.WechatBindingDTO) (bool, error) // 微信三方账号绑定
	WeiboOAuthBinding(dto dtos.WeiboBindingDTO) (bool, error)   // 微博三方账号绑定
}
