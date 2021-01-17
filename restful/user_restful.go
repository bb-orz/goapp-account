package restful

import (
	"errors"
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
	oauthGroup.GET("/qq", api.qqOAuthLoginHandler)
	oauthGroup.GET("/weixin", api.wechatOAuthLoginHandler)
	oauthGroup.GET("/weibo", api.weiboOAuthLoginHandler)

	// 用户鉴权访问路由组接口
	userGroup := engine.Group("/user", middleware.JwtAuthMiddleware())
	userGroup.GET("/info", api.getUserInfoHandler)
	userGroup.POST("/set_avatar", api.setAvatarHandler)
	userGroup.POST("/set", api.setUserInfoHandler)
	userGroup.POST("/modified_password", api.modifiedPassword)
	userGroup.POST("/forget_password", api.resetForgetPassword)
	userGroup.POST("/verify_mail", api.verifyEmail)
	userGroup.POST("/verify_phone", api.verifyPhone)
	oauthGroup.GET("/qq_binding", api.qqOAuthBindingHandler)
	oauthGroup.GET("/weixin_binding", api.wechatOAuthBindingHandler)
	oauthGroup.GET("/weibo_binding", api.weiboOAuthBindingHandler)
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
func (api *UserApi) qqOAuthLoginHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.QQLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.QQOAuthLogin(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

/*微信oauth 登录*/
func (api *UserApi) wechatOAuthLoginHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.WechatLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.WechatOAuthLogin(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

/*微博oauth登录*/
func (api *UserApi) weiboOAuthLoginHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.WeiboLoginDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	token, err := userService.WeiboOAuthLogin(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

}

func (api *UserApi) setAvatarHandler(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.SetAvatarUriDTO

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

	userService := services.GetUserService()
	if isPass, err := userService.SetAvatarUri(dto); !isPass || err != nil {
		_ = ctx.Error(common.ErrorOnInnerServer(" Error"))
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

/*设置用户信息*/
func (api *UserApi) setUserInfoHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.SetUserInfoDTO
	var userId uint
	var err error
	err = ctx.ShouldBindJSON(&dto)
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

	// Call Services method ...
	userService := services.GetUserService()
	if isPass, err := userService.SetUserInfos(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

/*获取用户信息*/
func (api *UserApi) getUserInfoHandler(ctx *gin.Context) {
	var dto dtos.GetUserInfoDTO
	var userId uint
	var err error

	// 校验登录用户id是否有获取信息权限
	if userId, err = common.GetUserId(ctx); err != nil {
		_ = ctx.Error(common.ErrorOnAuthenticate("No Permission"))
		return
	} else {
		dto.Id = userId
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

// 修改密码
func (api *UserApi) modifiedPassword(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.ModifiedPasswordDTO

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

	userService := services.GetUserService()
	if isPass, err := userService.ModifiedPassword(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, nil)
}

// 忘记密码:先发送邮箱验证码，用户接收后和新密码一起提交
func (api *UserApi) resetForgetPassword(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.ResetForgetPasswordDTO

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

	userService := services.GetUserService()
	if isPass, err := userService.ResetForgetPassword(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, nil)
}

// 验证账号邮箱
func (api *UserApi) verifyEmail(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.ValidateEmailDTO

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

	userService := services.GetUserService()
	if isPass, err := userService.ValidateEmail(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

// 验证账号手机号码
func (api *UserApi) verifyPhone(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.ValidatePhoneDTO

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

	userService := services.GetUserService()
	if isPass, err := userService.ValidatePhone(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

/*qq 账号绑定*/
func (api *UserApi) qqOAuthBindingHandler(ctx *gin.Context) {
	var ok bool
	var err error
	var userId uint
	var dto dtos.QQBindingDTO

	// Receive Request ...
	err = ctx.ShouldBindJSON(&dto)
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

	// Call Services method ...
	userService := services.GetUserService()
	if ok, err = userService.QQOAuthBinding(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	if ok {
		ctx.Set(common.ResponseDataKey, nil)
	} else {
		_ = ctx.Error(errors.New("QQ OAuth Binding Fail "))
	}
}

/*微信 账号绑定*/
func (api *UserApi) wechatOAuthBindingHandler(ctx *gin.Context) {
	var ok bool
	var err error
	var userId uint
	var dto dtos.WechatBindingDTO

	// Receive Request ...
	err = ctx.ShouldBindJSON(&dto)
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

	// Call Services method ...
	userService := services.GetUserService()
	if ok, err = userService.WechatOAuthBinding(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	if ok {
		ctx.Set(common.ResponseDataKey, nil)
	} else {
		_ = ctx.Error(errors.New("Wechat OAuth Binding Fail "))
	}
}

/*微博账号绑定*/
func (api *UserApi) weiboOAuthBindingHandler(ctx *gin.Context) {
	var ok bool
	var err error
	var userId uint
	var dto dtos.WeiboBindingDTO

	// Receive Request ...
	err = ctx.ShouldBindJSON(&dto)
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

	// Call Services method ...
	userService := services.GetUserService()
	if ok, err = userService.WeiboOAuthBinding(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	if ok {
		ctx.Set(common.ResponseDataKey, nil)
	} else {
		_ = ctx.Error(errors.New("Weibo OAuth Binding Fail "))
	}

}
