package middleware

import "github.com/gin-gonic/gin"

func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 响应在所有请求业务逻辑之后
		ctx.Next()

		// 封装响应信息


	}
}