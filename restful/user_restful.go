package restful

import (
	"github.com/bb-orz/goinfras/XGin"
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/gin-gonic/gin"
	"goinfras-sample-account/common"
	"goinfras-sample-account/dtos"
	"goinfras-sample-account/restful/middleware"
	"goinfras-sample-account/services"
	"net/http"
)

/*
API层，调用相关Service，封装响应返回，并记录日志
*/

func init() {
	// 初始化时注册该模块API
	XGin.RegisterApi(new(UserApi))
}

type UserApi struct {}

// 设置该模块的API Router
func (api *UserApi) SetRoutes() {
	engine := XGin.XEngine()

	// 如TokenUtils服务已初始化，添加中间件
	var authMiddleware gin.HandlerFunc
	if tku := XJwt.XTokenUtils(); tku == nil {
		authMiddleware = middleware.JwtAuthMiddleware()
	}

	engine.POST("/login", api.loginHandler)
	engine.POST("/logout", api.logoutHandler)

	registerGroup := engine.Group("/register")
	registerGroup.POST("/email", api.registerEmailHandler)
	registerGroup.POST("/phone", api.registerPhoneHandler)

	oauthGroup := engine.Group("/oauth")
	oauthGroup.GET("/qq", api.oauthQQHandler)
	oauthGroup.GET("/weixin", api.oauthWeixinHandler)
	oauthGroup.GET("/weibo", api.oauthWeiboHandler)

	userGroup := engine.Group("/user", authMiddleware)
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
	ctx.Set(common.ResponseDataKey,common.ResponseOK(gin.H{"token":token}))
}

/*用户登出*/
func (api *UserApi) logoutHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*邮箱注册注册*/
func (api *UserApi) registerEmailHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*手机号码注册注册*/
func (api *UserApi) registerPhoneHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*qq oauth 登录*/
func (api *UserApi) oauthQQHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*微信oauth 登录*/
func (api *UserApi) oauthWeixinHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*微博oauth登录*/
func (api *UserApi) oauthWeiboHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*设置用户信息*/

func (api *UserApi) setUserInfoHandler(ctx *gin.Context) {
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...

}

/*获取用户信息*/
func (api *UserApi) getUserInfoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK,gin.H{"id":id})
	// TODO Receive Request ...

	// TODO Call Services method ...

	// TODO Send Data to Response Middleware ...


}
