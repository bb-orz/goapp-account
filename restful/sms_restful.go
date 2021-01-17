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
		XGin.RegisterApi(new(SmsApi))
	})
}

type SmsApi struct{}

// 设置该模块的API Router
func (api *SmsApi) SetRoutes() {
	engine := XGin.XEngine()

	// 用户鉴权访问路由组接口
	userGroup := engine.Group("/sms", middleware.JwtAuthMiddleware())
	userGroup.GET("/send_sms_verify_code", api.sendSmsForVerifyPhoneNum)

}

func (api *SmsApi) sendSmsForVerifyPhoneNum(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var userId uint
	var dto dtos.SendPhoneVerifyCodeDTO
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

	smsService := services.GetSmsService()
	if err = smsService.SendPhoneVerifyCode(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}
