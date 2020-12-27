package restful

import (
	"github.com/bb-orz/goinfras/XGin"
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
	"goinfras-sample-account/dtos"
	"goinfras-sample-account/restful/middleware"
	"goinfras-sample-account/services"
)

/*
API层，调用相关Service，封装响应返回，并记录日志
*/

func init() {
	// 初始化时注册该模块API
	XGin.RegisterApi(new(UserApi))
}

type UserApi struct{}

// 设置该模块的API Router
func (api *UserApi) SetRoutes() {
	engine := XGin.XEngine()

	engine.POST("/login", api.loginHandler)
	engine.POST("/logout", api.logoutHandler)

	registerGroup := engine.Group("/register")
	registerGroup.POST("/email", api.registerEmailHandler)
	registerGroup.POST("/phone", api.registerPhoneHandler)

	oauthGroup := engine.Group("/oauth")
	oauthGroup.GET("/qq", api.oauthQQHandler)
	oauthGroup.GET("/weixin", api.oauthWechatHandler)
	oauthGroup.GET("/weibo", api.oauthWeiboHandler)

	// 用户鉴权访问路由组
	userGroup := engine.Group("/user", middleware.JwtAuthMiddleware())
	userGroup.GET("/:id", api.getUserInfoHandler)
	userGroup.POST("/set", api.setUserInfoHandler)
}

/*用户登录*/
func (api *UserApi) loginHandler(ctx *gin.Context) {

	// 接收参数由dto封装验证
	var dto dtos.AuthWithEmailPasswordDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// 调用service方法处理核心业务逻辑
	userService := services.GetUserService()
	token, err := userService.EmailAuth(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// 最后数据传递给响应中间件处理
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))
}

/*用户登出*/
func (api *UserApi) logoutHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*邮箱注册注册*/
func (api *UserApi) registerEmailHandler(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.CreateUserWithEmailDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	userDTO, err := userService.CreateUserWithEmail(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, *userDTO)
}

/*手机号码注册注册*/
func (api *UserApi) registerPhoneHandler(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.CreateUserWithPhoneDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	userDTO, err := userService.CreateUserWithPhone(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, *userDTO)

}

/*qq oauth 登录*/
func (api *UserApi) oauthQQHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.QQLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.QQOAuth(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

/*微信oauth 登录*/
func (api *UserApi) oauthWechatHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.WechatLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.WechatOAuth(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

/*微博oauth登录*/
func (api *UserApi) oauthWeiboHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.WeiboLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.WeiboOAuth(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

/*设置用户信息*/

func (api *UserApi) setUserInfoHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.SetUserInfoDTO
	var err error
	err = ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	err = userService.SetUserInfos(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

/*获取用户信息*/
func (api *UserApi) getUserInfoHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.GetUserInfoDTO
	var err error
	err = ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	userDTO, err := userService.GetUserInfo(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, *userDTO)
}
