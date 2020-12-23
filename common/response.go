package common

/* 统一响应信息格式 */

type CResponse struct {
	Code uint // 自定义响应码
	Message string // 自定义码解释
	Data interface{} // 放置任何类型的返回数据

}