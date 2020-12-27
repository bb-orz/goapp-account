package common

import (
	"fmt"
)

// 可暴露给客户端的错误信息类型
type CError struct {
	Code uint				`json:"code"`				// 自定义错误码
	Message    string		`json:"message"`				// 格式化错误信息，建议由fmt.Sprintf封装
	Err  error				`json:"error"`				// 基本错误信息
}

// 实现error 接口
func (err CError) Error() string {
	return err.Message
}



// 客户端错误信息码
const (
	ErrorOnValidateCode = 20001 	// 请求参数校验错误码，用于验证是否有效
	ErrorOnVerifyCode = 20002 	// 数据验证错误码，用于验证数据是否核实正确，如验证码
	ErrorOnNetRequestCode = 20003 // 网络请求错误码，调用其他服务时出现的网络错误，如三方接口访问
	ErrorOnAuthenticateCode = 20004 	// 身份验证鉴权错误
)

// 客户端错误信息格式
const (
	ErrorOnValidateMessage = "[Error]: Validate Parameters Error" 	// 请求参数校验错误码，用于验证是否有效
	ErrorOnVerifyMessage = "[Error]: Verify Data Error | [Info]:%s" 	// 数据验证错误码，用于验证数据是否核实正确，如验证码
	ErrorOnNetRequestMessage = "[Error]: Network Request Error | [Request]:%s" // 网络请求错误码，调用其他服务时出现的网络错误，如三方接口访问
	ErrorOnAuthenticateMessage = "" 	// 身份验证鉴权错误
)


// 客户端请求错误：验证请求参数DTO错误
func ErrorOnValidate(err error) CError  {
	return CError{
		Code:		ErrorOnValidateCode,
		Message:   ErrorOnValidateMessage,
		Err:		err,
	}
}

// 客户端请求错误：与服务端检查信息错误
func ErrorOnVerify(err error,info string) CError  {
	return CError{
		Code:	  	ErrorOnVerifyCode,
		Message:  	fmt.Sprintf(ErrorOnVerifyMessage ,info),
		Err:      	err,
	}
}

// 客户端请求错误：网络请求错误
func ErrorOnNetRequest(err error,info string) CError  {
	return CError{
		Code:		ErrorOnNetRequestCode,
		Message:  	fmt.Sprintf( ErrorOnNetRequestMessage,info),
		Err:     	err,
	}
}

// 客户端请求错误：身份认证鉴权错误
func ErrorOnAuthenticate(err error,info string) CError  {
	return CError{
		Code:		ErrorOnAuthenticateCode,
		Message:  	fmt.Sprintf(ErrorOnAuthenticateMessage ,info),
		Err:     	err,
	}
}