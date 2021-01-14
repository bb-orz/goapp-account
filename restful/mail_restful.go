package restful

import (
	"github.com/bb-orz/goinfras/XGin"
	"github.com/gin-gonic/gin"
	"goapp/common"
	"goapp/dtos"
	"goapp/restful/middleware"
	"goapp/services"
	"sync"
)

/*
API层，调用相关Service，封装响应返回，并记录日志
*/

func init() {
	var once sync.Once
	once.Do(func() {
		// 初始化时注册该模块API
		XGin.RegisterApi(new(MailApi))
	})
}

type MailApi struct{}

// 设置该模块的API Router
func (api *MailApi) SetRoutes() {
	engine := XGin.XEngine()

	// 用户鉴权访问路由组接口
	userGroup := engine.Group("/mail", middleware.JwtAuthMiddleware())
	userGroup.GET("/send_email_verify_code", api.sendEmailForVerifyEmailAddress)
	userGroup.POST("/send_forget_password_verify_code", api.sendEmailForForgetPassword)

}

func (api *MailApi) sendEmailForVerifyEmailAddress(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var userId uint
	var dto dtos.SendEmailForVerifyDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// 校验登录用户id是否有获取信息权限
	if userId, err = common.GetUserId(ctx); err != nil {
		_ = ctx.Error(common.ErrorOnAuthenticate("No Permission"))
		return
	} else {
		dto.Id = userId
	}

	mailService := services.GetMailService()
	if err = mailService.SendEmailForVerify(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

func (api *MailApi) sendEmailForForgetPassword(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var userId uint
	var dto dtos.SendEmailForgetPasswordDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// 校验登录用户id是否有获取信息权限
	if userId, err = common.GetUserId(ctx); err != nil {
		_ = ctx.Error(common.ErrorOnAuthenticate("No Permission"))
		return
	} else {
		dto.Id = userId
	}

	mailService := services.GetMailService()
	if err = mailService.SendEmailForgetPassword(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}
