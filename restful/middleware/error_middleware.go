package middleware

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
	"gopkg.in/go-playground/validator.v9"
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
			case common.SError:
				XLogger.XCommon().Error(err.(common.SError).Printf()) 						 // 服务端内部业务错误，需记录日志，并返回统一的服务端错误信息，隐藏内部错误
				ctx.JSON(http.StatusInternalServerError,common.ResponseServerInnerError())	 // 输出统一的服务器错误响应信息
			case common.CError:
				// 返回统一定制的客户端错误信息
				switch err.(common.CError).Err.(type){
				case validator.ValidationErrors:  // 请求参数验证错误
					// TODO 继续优化封装
					ctx.JSON(http.StatusNotAcceptable,map[string]interface{}{
						"code":common.ResponseCodeValidateFail,
						"message":err.(common.CError).Message,
						"error":err.(common.CError).Err.(validator.ValidationErrors).Translate(XValidate.XTranslater()),
					})
				default: // 默认客户端错误
					ctx.JSON(http.StatusBadRequest, err.(common.CError))
				}

			case gin.Error:			// gin 引擎错误，返回统一的错误信息
				errorType := err.(gin.Error).Type
				switch errorType {
				case gin.ErrorTypeBind : 	// ErrorTypeBind is used when Context.Bind() fails. = 1 << 63
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"ErrorTypeBind"})
				case gin.ErrorTypeRender:	// ErrorTypeRender is used when Context.Render() fails. = 1 << 62
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"ErrorTypeRender"})
				case gin.ErrorTypePrivate:	// ErrorTypePrivate indicates a private error. = 1 << 0
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"ErrorTypePrivate"})
				case gin.ErrorTypePublic:	// ErrorTypePublic indicates a public error. = 1 << 1
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"ErrorTypePublic"})
				case gin.ErrorTypeAny:		// ErrorTypeAny indicates any other error. = 1<<64 - 1
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"ErrorTypeAny"})
				default:
					ctx.JSON(http.StatusInternalServerError,gin.H{"GinErrorType":"Gin Engine Error"})
				}
			}
		}
	}
}