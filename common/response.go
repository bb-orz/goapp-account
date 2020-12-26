package common

const (
	ResponseHeaderKey = "_resp_header"  // 传递给ResponseMiddleware的http header设置键名
	ResponseDataKey = "_resp_data"		// // 传递给ResponseMiddleware的响应信息设置键名

)

const (
	CodeOK  = (iota + 1) * 1000


)

const (
	MessageOK = "OK"

)


/* 正常响应信息格式 */
type Response struct {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Data interface{} 	`json:"data"`			// 放置任何类型的返回数据
}

/* 客户端错误信息响应格式*/
type CEResponse struct {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Error CError 		`json:"data"`			// 客户端错误信息
}

/*服务端错误信息响应格式*/
type SEResponse struct  {
	Code int			`json:"code"`			// 自定义响应码
	Message string 		`json:"msg"`			// 自定义码解释
	Error error 		`json:"data"`			// 可暴露给用户的服务端错误信息
}


// 正常响应
func ResponseOK(data interface{}) Response {
	return Response{
		Code: CodeOK,
		Message: MessageOK,
		Data:data,
	}
}

// 无访问权限错误响应
func ResponseNoAuth() CEResponse {
	return CEResponse{

	}
}

// 请求参数失败响应
func ResponseValidateFail() CEResponse {
	return CEResponse{}
}

// 数据验证失败响应
func ResponseVerifiedFail() CEResponse  {
	return CEResponse{}
}

func ResponseServerInnerError() SEResponse  {
	return SEResponse{}
}