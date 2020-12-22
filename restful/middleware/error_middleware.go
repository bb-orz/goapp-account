package middleware

import (
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO 判断错误类型是服务端内部错误还是客户端逻辑错误

		 if len(ctx.Errors) > 0 {
		 	e := ctx.Errors[0].Err
			switch e.(type) {
			case common.SError :
				// TODO 服务端错误记录日志 ...，并返回统一的服务端错误信息，隐藏内部错误



			case common.CError:
				// TODO 返回定制的客户端错误信息



			}


		 }

		 ctx.Next()
	}
}