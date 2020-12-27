package common

import (
	"fmt"
	"runtime/debug"
)

// 服务端内部错误信息格式,客户端不可见，用于日志输出
const (
	// 服务端错误顶层包装信息格式
	ErrorOnServerInnerFormat = "[Server Inner Error]: Domain:%s"

	// SQL数据库执行错误信息格式
	ErrorOnSqlQueryFormat      = "[Domain Inner Error]: SQL Query  Error | [CALL]:dao.%s"     // sql查询错误
	ErrorOnSqlInsertFormat     = "[Domain Inner Error]: SQL Insert Error | [CALL]:dao.%s"     // sql插入错误
	ErrorOnSqlUpdateFormat     = "[Domain Inner Error]: SQL Update Error | [CALL]:dao.%s"     // sql更新错误
	ErrorOnSqlDeleteFormat     = "[Domain Inner Error]: SQL Delete Error | [CALL]:dao.%s"     // sql删除错误
	ErrorOnSqlShamDeleteFormat = "[Domain Inner Error]: SQL ShamDelete Error | [CALL]:dao.%s" // sql更新deleted_at字段错误，假删除

	// NoSQL数据库执行错误信息格式
	ErrorOnNoSqlQueryFormat      = "[Domain Inner Error]: NoSQL Query  Error | [CALL]:store.%s"   // nosql查询错误
	ErrorOnNoSqlInsertFormat     = "[Domain Inner Error]: NoSQL Insert Error | [CALL]:store.%s"   // nosql插入错误
	ErrorOnNoSqlUpdateFormat     = "[Domain Inner Error]: NoSQL Update Error | [CALL]:store.%s"   // nosql更新错误
	ErrorOnNoSqlDeleteFormat     = "[Domain Inner Error]: NoSQL Delete Error | [CALL]:store.%s"   // nosql删除错误
	ErrorOnNoSqlShamDeleteFormat = "[Domain Inner Error]: NoSQL ShamDelete Error | [CALL]:dao.%s" // nosql更新deleted_at字段错误，假删除

	// 缓存执行错误信息格式
	ErrorOnCacheSetFormat    = "[Domain Inner Error]: Cache Set Error | [command]:%s"    // 缓存设置错误
	ErrorOnCacheGetFormat    = "[Domain Inner Error]: Cache Get Error | [command]:%s"    // 缓存获取错误
	ErrorOnCacheDeleteFormat = "[Domain Inner Error]: Cache Delete Error | [command]:%s" // 缓存删除错误

	// 服务端内部网络请求报错
	ErrorOnNetRequestFormat = "[Domain Inner Error]: Network Request Error | [Request]:%v"   // 网络请求相关错误
	ErrorOnThirdPartFormat  = "[Domain Inner Error]: Network ThirdPart Error | [Request]:%v" // 第三方接口错误相关错误

	// 算法逻辑类错误信息格式
	ErrorOnAlgorithmFormat = "[Domain Inner Error]: Algorithm Error | [Info]:%s" // 算法执行错误

	// 服务端内部编解码错误
	ErrorOnEncodeFormat = "[Domain Inner Error]: Encode Error | source data:%v" // 编码算法错误
	ErrorOnDecodeFormat = "[Domain Inner Error]: Decode Error | code string:%s" // 解码算法错误

)

// 服务端内部错误类型
type InnerError struct {
	Inner      error                  // 存储我们正在包装的错误。 如果需要调查发生的事情，我们总是希望能够查看到最低级别的错误。
	Message    string                 // 格式化错误信息，建议由fmt.Sprintf封装
	StackTrace string                 // 记录了创建错误时的堆栈跟踪。
	Misc       map[string]interface{} // 创建一个杂项信息存储字段。可以存储并发ID，堆栈跟踪的hash或可能有助于诊断错误的其他上下文信息。
}

// 实现error 接口
func (err InnerError) Error() string {
	return err.Message
}

// 打印包装器的错误信息，供日志记录器写入
func (err InnerError) Printf() string {
	var innerFormat string
	switch err.Inner.(type) {
	case InnerError:
		innerFormat = "\t\t" + err.Inner.(InnerError).Printf()
	default:
		innerFormat = "[Final Error]:" + err.Inner.Error() + "\n =======================================[Inner Error End Off]===========================================\n"
	}
	return fmt.Sprintf("\n\t %s  | \n\t [StackTrace]:=========>  %s \n | [Inner]:%s ", err.Message, err.StackTrace, innerFormat)
}

// 工具函数：服务器内部错误信息在系统各模块传递时的“错误包装器”
func ServerErrorWrapper(err error, messageFormat string, msgArgs ...interface{}) InnerError {
	return InnerError{
		Inner:      err,
		Message:    fmt.Sprintf(messageFormat, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

// 服务端错误顶层包装信息包装方法
func ErrorOnServerInner(err error, domain string) InnerError {
	return ServerErrorWrapper(err, ErrorOnServerInnerFormat, domain)
}

// sql查询错误包装方法
func DomainInnerErrorOnSqlQuery(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnSqlQueryFormat, method)
}

// sql插入错误包装方法
func DomainInnerErrorOnSqlInsert(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnSqlInsertFormat, method)
}

// sql更新错误包装方法
func DomainInnerErrorOnSqlUpdate(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnSqlUpdateFormat, method)
}

// sql删除错误包装方法
func DomainInnerErrorOnSqlDelete(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnSqlDeleteFormat, method)
}

// sql假删除错误包装方法
func DomainInnerErrorOnSqlShamDelete(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnSqlShamDeleteFormat, method)
}

// nosql查询错误包装方法
func DomainInnerErrorOnNoSqlQuery(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnNoSqlQueryFormat, method)
}

// nosql插入错误包装方法
func DomainInnerErrorOnNoSqlInsert(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnNoSqlInsertFormat, method)
}

// nosql更新错误包装方法
func DomainInnerErrorOnNoSqlUpdate(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnNoSqlUpdateFormat, method)
}

// nosql删除错误包装方)
func DomainInnerErrorOnNoSqlDelete(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnNoSqlDeleteFormat, method)
}

// nosql假删除错误包装方法
func DomainInnerErrorOnNoSqlShamDelete(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnNoSqlShamDeleteFormat, method)
}

// 缓存设置错误包装方法
func DomainInnerErrorOnCacheSet(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnCacheSetFormat, method)
}

// 缓存获取错误包装方法
func DomainInnerErrorOnCacheGet(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnCacheGetFormat, method)
}

// 缓存获取错误包装方法
func DomainInnerErrorOnCacheDelete(err error, method string) InnerError {
	return ServerErrorWrapper(err, ErrorOnCacheDeleteFormat, method)
}

// 网络请求错误包装方法
func DomainInnerErrorOnNetRequest(err error, req interface{}) InnerError {
	return ServerErrorWrapper(err, ErrorOnNetRequestFormat, req)
}

// 第三方服务请求错误包装方法
func DomainInnerErrorOnThirdPartRequest(err error, req interface{}) InnerError {
	return ServerErrorWrapper(err, ErrorOnThirdPartFormat, req)
}

// 内部业务算法错误包装方法
func DomainInnerErrorOnAlgorithm(err error, info string) InnerError {
	return ServerErrorWrapper(err, ErrorOnAlgorithmFormat, info)
}

// 数据编码算法错误
func DomainInnerErrorOnEncodeData(err error, data interface{}) InnerError {
	return ServerErrorWrapper(err, ErrorOnEncodeFormat, data)
}

// 数据解码算法错误
func DomainInnerErrorOnDecodeData(err error, code string) InnerError {
	return ServerErrorWrapper(err, ErrorOnDecodeFormat, code)
}
