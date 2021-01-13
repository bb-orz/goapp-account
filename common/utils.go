package common

import (
	"github.com/gin-gonic/gin"
)

func GetUserClaim(ctx *gin.Context) map[string]interface{} {
	// value, exists := ctx.Get(ContextTokenUserClaimKey)
	// if exists {
	// 	fmt.Printf("ContextTokenUserClaimKey: %+v \n", value)
	// }

	return ctx.GetStringMap(ContextTokenUserClaimKey)
}

func GetTokenString(ctx *gin.Context) string {
	return ctx.GetString(ContextTokenStringKey)
}
