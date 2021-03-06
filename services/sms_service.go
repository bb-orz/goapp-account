package services

import "goapp/dtos"

/* 定义短信服务模块的服务层方法，并定义数据传输对象DTO*/
var smsService ISmsService

// 用于对外暴露短信服务，唯一的暴露点，供接口层调用
func GetSmsService() ISmsService {
	return smsService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetSmsService(service ISmsService) {
	smsService = service
}

type ISmsService interface {
	SendPhoneVerifyCode(dto dtos.SendPhoneVerifyCodeDTO) error // 发送短信验证码
}
