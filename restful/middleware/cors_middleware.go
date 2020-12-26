package middleware

import (
	"github.com/gin-gonic/gin"
)

// 跨域请求处理策略，cors/jsonp/
func CorsMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		// 需验证origin
		ctx.Next()

	}
}