package middleware

import (
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
)

// 响应信息中间件，统一处理每个请求的响应格式
func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 响应在所有请求业务逻辑之后
		ctx.Next()

		// 设置头部信息
		headers := ctx.GetStringMapString(common.ResponseHeaderKey)

		// 封装响应信息
		body ,_ := ctx.Get(common.ResponseDataKey)

	}
}