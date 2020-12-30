package middleware

import (
	"github.com/gin-gonic/gin"
	"goapp/common"
	"net/http"
)

// 响应信息中间件，统一处理每个请求的响应格式
func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 响应在所有请求业务逻辑之后
		ctx.Next()

		if len(ctx.Errors.Errors()) == 0 {
			// 设置头部信息
			headers := ctx.GetStringMapString(common.ResponseHeaderKey)
			for k, v := range headers {
				ctx.Header(k, v)
			}

			// 有数据正常响应信息
			data, isExist := ctx.Get(common.ResponseDataKey)
			if isExist {
				if data != nil {
					ctx.JSON(http.StatusOK, data)
				} else {
					ctx.JSON(http.StatusOK, common.ResponseOK(nil))
				}
			}
		}
	}
}
