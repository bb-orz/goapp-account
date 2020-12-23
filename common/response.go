package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeOK  = (iota + 1) * 1000


)

const (
	StatusOK = "OK"

)


/* 统一响应信息格式 */
type Response struct {
	Code int			`json:"code"`			// 自定义响应码
	Status string 		`json:"status"`			// 自定义码解释
	Data interface{} 	`json:"data"`			// 放置任何类型的返回数据
}

func GinResponseOK(ctx *gin.Context ,data interface{})  {
	ctx.JSON(http.StatusOK,Response{
		Code: CodeOK,
		Status: StatusOK,
		Data:data,
	})
}