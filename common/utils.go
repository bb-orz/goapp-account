package common

import (
	"errors"
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/gin-gonic/gin"
)

func GetUserClaim(ctx *gin.Context) *XJwt.UserClaim {
	claim, exists := ctx.Get(ContextTokenUserClaimKey)
	if exists {
		userClaim := claim.(XJwt.UserClaim)
		return &userClaim
	} else {
		return nil
	}
}

func GetUserId(ctx *gin.Context) (uint, error) {
	claim, exists := ctx.Get(ContextTokenUserClaimKey)
	if exists {
		userClaim := claim.(XJwt.UserClaim)
		return userClaim.Id, nil
	} else {
		return 0, errors.New("User Token Not Exist! ")
	}
}

func GetTokenString(ctx *gin.Context) string {
	return ctx.GetString(ContextTokenStringKey)
}
