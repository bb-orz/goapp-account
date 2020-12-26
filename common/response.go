package common

const (
	ResponseHeaderKey = "_resp_header"  // 传递给ResponseMiddleware的http header设置键名
	ResponseDataKey = "_resp_data"		// // 传递给ResponseMiddleware的响应信息设置键名

)

const (
	ResponseCodeOK  = (iota + 1) * 1000
	ResponseCodeAuthFail
	ResponseCodeValidateFail
	ResponseCodeVerifiedFail
	ResponseCodeServerInnerError
)

const (
	ResponseMessageOK = "OK"
	ResponseMessageAuthFail = "Auth Fail"
	ResponseMessageValidateFail = "Validate Fail"
	ResponseMessageVerifiedFail = "Verified Fail"
	ResponseMessageServerInnerError = "Server Inner Error"
)


/* 正常响应信息格式 */
type Response struct {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Data interface{} 	`json:"error"`			// 放置任何类型的返回数据
}

/* 客户端错误信息响应格式*/
type CEResponse struct {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Error CError 		`json:"error"`			// 客户端错误信息
}

/*服务端错误信息响应格式*/
type SEResponse struct  {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Error string 		`json:"error"`			// 可暴露给用户的服务端错误信息
}


// 正常响应
func ResponseOK(data interface{}) Response {
	return Response{
		Code: ResponseCodeOK,
		Message: ResponseMessageOK,
		Data:data,
	}
}

// 无访问权限错误响应
func ResponseAuthFail(err CError) CEResponse {
	return CEResponse{
		Code:ResponseCodeAuthFail,
		Message :ResponseMessageAuthFail,
		Error : err,
	}
}

// 请求参数失败响应
func ResponseValidateFail(err CError) CEResponse {
	return CEResponse{
		Code:ResponseCodeValidateFail,
		Message :ResponseMessageValidateFail,
		Error : err,
	}
}

// 数据验证失败响应
func ResponseVerifiedFail(err CError) CEResponse  {
	return CEResponse{
		Code:ResponseCodeVerifiedFail,
		Message :ResponseMessageVerifiedFail,
		Error : err,
	}
}

func ResponseServerInnerError() SEResponse  {
	return SEResponse{
		Code:ResponseCodeServerInnerError,
		Message :ResponseMessageServerInnerError,
		Error : "Server Inner Error,Please contact developer or administrator",
	}
}