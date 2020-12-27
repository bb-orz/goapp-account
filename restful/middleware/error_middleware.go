package middleware

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
	"net/http"
)

// 统一错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 错误处理在业务处理后，响应之前
		ctx.Next()

		// 判断错误类型是服务端内部错误还是客户端逻辑错误
		e := ctx.Errors.Last()

		if e != nil {
			err := e.Err
			switch err.(type) {
			case common.InnerError: // 服务端内部业务错误
				innerError := err.(common.InnerError)
				// 记录日志，对外隐藏内部错误
				XLogger.XCommon().Error(innerError.Printf())
				// 转换成公共错误输出响应，统一输出服务端执行失败响应信息
				ctx.JSON(http.StatusInternalServerError, common.ErrorOnInnerServer(innerError.Message))
			case common.PublishError:
				publishError := err.(common.PublishError)
				// 返回统一定制的客户端可见错误信息
				switch publishError.Code {
				case common.ErrorOnValidateCode:
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnVerifyCode:
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnNetRequestCode:
					ctx.JSON(http.StatusNotAcceptable, publishError)
				case common.ErrorOnAuthenticateCode:
					ctx.JSON(http.StatusNotAcceptable, publishError)

				default: // 默认客户端错误
					ctx.JSON(http.StatusBadRequest, publishError)
				}

			case gin.Error: // gin 引擎错误，返回统一的错误信息
				errorType := err.(gin.Error).Type
				switch errorType {
				case gin.ErrorTypeBind: // ErrorTypeBind is used when Context.Bind() fails. = 1 << 63
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "ErrorTypeBind"})
				case gin.ErrorTypeRender: // ErrorTypeRender is used when Context.Render() fails. = 1 << 62
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "ErrorTypeRender"})
				case gin.ErrorTypePrivate: // ErrorTypePrivate indicates a private error. = 1 << 0
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "ErrorTypePrivate"})
				case gin.ErrorTypePublic: // ErrorTypePublic indicates a public error. = 1 << 1
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "ErrorTypePublic"})
				case gin.ErrorTypeAny: // ErrorTypeAny indicates any other error. = 1<<64 - 1
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "ErrorTypeAny"})
				default:
					ctx.JSON(http.StatusInternalServerError, gin.H{"GinErrorType": "Gin Engine Error"})
				}
			}
		}
	}
}
