package middleware

import (
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/gin-gonic/gin"
	"goapp/common"
)

// 用户鉴权中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1.从http头获取token string
		tkStr := ctx.GetHeader("Authorization")
		// fmt.Println("token string:",tkStr)
		if tkStr == "" {
			_ = ctx.Error(common.ErrorOnAuthenticate("Token Parameter on http header is required")) // 所有错误最后传递给错误中间件处理
			ctx.Abort()
			return
		}

		// 2.解码校验token是否合法
		customerClaim, err := XJwt.XTokenUtils().Decode(tkStr)
		if err != nil {
			_ = ctx.Error(common.ErrorOnAuthenticate(err.Error()))
			ctx.Abort()
			return
		}

		// 鉴权通过后设置用户信息
		ctx.Set(common.ContextTokenStringKey, tkStr)
		ctx.Set(common.ContextTokenUserClaimKey, customerClaim.UserClaim)

		ctx.Next()
	}
}
