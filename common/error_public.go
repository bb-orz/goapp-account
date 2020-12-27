package common

import (
	"github.com/bb-orz/goinfras/XValidate"
	"gopkg.in/go-playground/validator.v9"
)

// 可暴露给客户端的错误信息类型，用于错误中间件的错误响应输出

type PublishError struct {
	Code    uint        `json:"code"`    // 自定义错误码
	Message string      `json:"message"` // 格式化错误信息，建议由fmt.Sprintf封装
	Err     interface{} `json:"error"`   // 基本错误信息
}

// 实现error 接口
func (err PublishError) Error() string {
	return err.Message
}

// 公开错误信息码
const (
	ErrorOnValidateCode         = 20001 // 请求参数校验错误码，用于验证是否有效
	ErrorOnVerifyCode           = 20002 // 数据验证错误码，用于验证数据是否核实正确，如验证码
	ErrorOnNetRequestCode       = 20003 // 网络请求错误码，调用其他服务时出现的网络错误，如三方接口访问
	ErrorOnAuthenticateCode     = 20004 // 身份验证鉴权错误
	ErrorOnCommonBadRequestCode = 20005 // 所有其他错误请求

	ErrorOnServerInnerCode = 30001 // 服务端内部错误码
)

// 公开错误信息格式
const (
	ErrorOnValidateMessage         = "Validate Parameters Fail" // 请求参数校验错误码，用于验证是否有效
	ErrorOnVerifyMessage           = "Verify Data Fail"         // 数据验证错误码，用于验证数据是否核实正确，如验证码
	ErrorOnNetworkRequestMessage   = "Network Request Fail"     // 网络请求错误码，调用其他服务时出现的网络错误，如三方接口访问
	ErrorOnAuthenticateMessage     = "Authentication Fail"      // 身份验证鉴权错误
	ErrorOnServerInnerMessage      = "Server Error"
	ErrorOnCommonBadRequestMessage = "BadRequest"
)

// 验证请求参数错误响应信息
func ErrorOnValidate(err error) PublishError {
	return PublishError{
		Code:    ErrorOnValidateCode,
		Message: ErrorOnValidateMessage,
		Err:     err.(validator.ValidationErrors).Translate(XValidate.XTranslater()),
	}
}

// 与后端数据检查数据不通过响应信息
func ErrorOnVerify(errInfo string) PublishError {
	return PublishError{
		Code:    ErrorOnVerifyCode,
		Message: ErrorOnVerifyMessage,
		Err:     errInfo,
	}
}

// 网络请求错误响应信息
func ErrorOnNetworkRequest(errInfo string) PublishError {
	return PublishError{
		Code:    ErrorOnNetRequestCode,
		Message: ErrorOnNetworkRequestMessage,
		Err:     errInfo,
	}
}

// 身份认证鉴权错误响应信息
func ErrorOnAuthenticate(errInfo string) PublishError {
	return PublishError{
		Code:    ErrorOnAuthenticateCode,
		Message: ErrorOnAuthenticateMessage,
		Err:     errInfo,
	}
}

func ErrorOnBadRequest(errInfo string) PublishError {
	return PublishError{
		Code:    ErrorOnCommonBadRequestCode,
		Message: ErrorOnCommonBadRequestMessage,
		Err:     errInfo,
	}
}

// 服务端内部错误响应信息
func ErrorOnInnerServer(errInfo string) PublishError {
	return PublishError{
		Code:    ErrorOnServerInnerCode,
		Message: ErrorOnServerInnerMessage,
		Err:     errInfo,
	}
}

// 服务端内部错误响应信息
