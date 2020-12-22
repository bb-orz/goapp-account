package common

import "fmt"

// 可暴露给客户端的错误类型
type CError struct {
	Code uint								// 自定义错误码
	Err error								// 基本错误信息
	Message    string						// 格式化错误信息，建议由fmt.Sprintf封装
}

// 实现error 接口
func (err CError) Error() string {
	return err.Message
}



// 客户端错误信息格式
const (
	ClientErrorOnValidateCode = 10001
	ClientErrorOnCheckInfoCode = 10002
	ClientErrorOnNetRequestCode = 10003

)


// 客户端请求错误：验证请求参数DTO错误
func ClientErrorOnValidateParameters(err error) CError  {
	return CError{
		Code:ClientErrorOnValidateCode,
		Err:      err,
		Message:   "[Client Error]: Validate Parameters Error",
	}
}

// 客户端请求错误：与服务端检查信息错误
func ClientErrorOnCheckInformation(err error,info string) CError  {
	return CError{
		Code:ClientErrorOnCheckInfoCode,
		Err:      err,
		Message:  fmt.Sprintf("[Client Error]: Check Information Error | [Info]:%s" ,info),
	}
}

// 客户端请求错误：网络请求错误
func ClientErrorOnNetRequest(err error,info string) CError  {
	return CError{
		Code:ClientErrorOnNetRequestCode,
		Err:      err,
		Message:  fmt.Sprintf("[Client Error]: Network Request Error | [Request]:%s" ,info),
	}
}