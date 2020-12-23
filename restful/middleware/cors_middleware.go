package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		// 需验证origin



		ctx.Next()
	}
}