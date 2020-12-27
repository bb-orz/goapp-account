package common

const (
	ResponseHeaderKey = "_resp_header" // 传递给ResponseMiddleware的http header设置键名
	ResponseDataKey   = "_resp_data"   // // 传递给ResponseMiddleware的响应信息设置键名

)

const (
	ResponseOKCode = (iota + 1) * 10000
)

const (
	ResponseOKMessage = "OK"
)

/* 正常响应信息格式 */
type Response struct {
	Code    int         `json:"code"` // 自定义响应码
	Message string      `json:"msg"`  // 自定义码解释
	Data    interface{} `json:"data"` // 放置任何类型的返回数据
}

// 正常响应
func ResponseOK(data interface{}) Response {
	return Response{
		Code:    ResponseOKCode,
		Message: ResponseOKMessage,
		Data:    data,
	}
}
