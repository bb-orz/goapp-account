package middleware

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-gonic/gin"
	"goapp/common"
	"net/http"
)

// 统一错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 错误处理在业务处理后，响应之前
		ctx.Next()

		// 判断错误类型是服务端内部错误还是可暴露给用户的错误信息
		e := ctx.Errors.Last()
		if e != nil {
			// 处理自定义错误
			err := e.Err
			switch err.(type) {
			case common.InnerError: // 服务端内部业务错误信息封装
				innerError := err.(common.InnerError)
				// 记录日志，对外隐藏内部错误
				XLogger.XCommon().Error(innerError.Printf())
				// 转换成公共错误输出响应，统一输出服务端执行失败响应信息
				ctx.JSON(http.StatusInternalServerError, common.ErrorOnInnerServer(innerError.Message))
			case common.PublishError: // 可暴露给用户的错误信息
				publishError := err.(common.PublishError)
				// 返回统一定制的客户端可见错误信息封装
				switch publishError.Code {
				case common.ErrorOnValidateCode: // 请求参数校验错误信息封装
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnVerifyCode: // 验证数据核实错误信息封装
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnNetRequestCode: // 服务端网络请求错误信息封装
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnAuthenticateCode: // 鉴权未通过错误信息封装
					ctx.JSON(http.StatusNonAuthoritativeInfo, publishError)
				}
			default: // gin 引擎引发的错误信息封装
				ctx.JSON(http.StatusBadRequest, common.ErrorOnBadRequest(e.Error()))
			}
			ctx.Abort()
		}

	}

}
