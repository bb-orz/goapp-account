package restful

import (
	"github.com/bb-orz/goinfras/XGin"
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/gin-gonic/gin"
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
		authMiddleware = XGin.JwtAuthMiddleware()
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
	userGroup.GET("/get", api.getUserInfoHandler)
	userGroup.POST("/set", api.setUserInfoHandler)
}

/*用户登录*/
func (api *UserApi) loginHandler(ctx *gin.Context) {
	// TODO Receive Request ...


	// TODO Call Services method ...



	// TODO Send Response ...
}

/*用户登出*/
func (api *UserApi) logoutHandler(ctx *gin.Context) {

}

/*邮箱注册注册*/
func (api *UserApi) registerEmailHandler(ctx *gin.Context) {

}

/*手机号码注册注册*/
func (api *UserApi) registerPhoneHandler(ctx *gin.Context) {

}

/*qq oauth 登录*/
func (api *UserApi) oauthQQHandler(ctx *gin.Context) {

}

/*微信oauth 登录*/
func (api *UserApi) oauthWeixinHandler(ctx *gin.Context) {

}

/*微博oauth登录*/
func (api *UserApi) oauthWeiboHandler(ctx *gin.Context) {

}

/*设置用户信息*/

func (api *UserApi) setUserInfoHandler(ctx *gin.Context) {

}

/*获取用户信息*/
func (api *UserApi) getUserInfoHandler(ctx *gin.Context) {

}
