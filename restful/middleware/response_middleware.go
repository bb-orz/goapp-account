package middleware

import (
	"github.com/gin-gonic/gin"
)

// 响应信息中间件，统一处理每个请求的响应格式
func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 响应在所有请求业务逻辑之后
		ctx.Next()


		// // 设置头部信息
		// headers := ctx.GetStringMapString(common.ResponseHeaderKey)
		// for k,v := range headers {
		// 	ctx.Header(k,v)
		// }
		//
		// // 封装响应信息
		// data ,_ := ctx.Get(common.ResponseDataKey)
		// ctx.JSON(http.StatusOK,data)
	}
}