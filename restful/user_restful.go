package restful

import (
	"errors"
	"fmt"
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
		XGin.RegisterApi(new(UserApi))
	})
}

type UserApi struct{}

// 设置该模块的API Router
func (api *UserApi) SetRoutes() {
	engine := XGin.XEngine()

	engine.GET("/logout", middleware.JwtAuthMiddleware(), api.logoutHandler)

	// 登录登出接口
	loginGroup := engine.Group("/login")
	loginGroup.POST("/email", api.loginEmailHandler)
	loginGroup.POST("/phone", api.loginPhoneHandler)

	// 邮箱或手机号注册账号接口
	registerGroup := engine.Group("/register")
	registerGroup.POST("/email", api.registerEmailHandler)
	registerGroup.POST("/phone", api.registerPhoneHandler)

	// 第三方平台登录或注册接口
	oauthGroup := engine.Group("/oauth")
	oauthGroup.GET("/qq", api.oauthQQHandler)
	oauthGroup.GET("/weixin", api.oauthWechatHandler)
	oauthGroup.GET("/weibo", api.oauthWeiboHandler)

	// 用户鉴权访问路由组接口
	userGroup := engine.Group("/user", middleware.JwtAuthMiddleware())
	userGroup.GET("/:id", api.getUserInfoHandler)
	userGroup.POST("/set", api.setUserInfoHandler)
}

/*邮箱账号登录*/
func (api *UserApi) loginEmailHandler(ctx *gin.Context) {

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

/*手机号码账号登录*/
func (api *UserApi) loginPhoneHandler(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.AuthWithPhonePasswordDTO
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	// 调用service方法处理核心业务逻辑
	userService := services.GetUserService()
	token, err := userService.PhoneAuth(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}
	// 最后数据传递给响应中间件处理
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))
}

/*用户登出*/
func (api *UserApi) logoutHandler(ctx *gin.Context) {
	// Receive Request ...
	token := ctx.GetString(common.ContextTokenStringKey)
	if token == "" {
		_ = ctx.Error(errors.New("Token is not set in context ")) // 所有错误最后传递给错误中间件处理
		return
	}

	var dto = dtos.RemoveTokenDTO{Token: token}

	// Call Services method ...
	userService := services.GetUserService()
	err := userService.RemoveToken(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, nil)
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
	var dto dtos.UserInfoDTO
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
	err = ctx.ShouldBindUri(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// 校验登录用户id是否有获取信息权限
	userClaim := common.GetUserClaim(ctx)
	fmt.Println(userClaim)

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
