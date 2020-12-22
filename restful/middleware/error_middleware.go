package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		 if len(ctx.Errors) > 0 {



		 }

		 ctx.Next()
	}
}