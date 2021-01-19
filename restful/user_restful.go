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

	// 非鉴权访问接口
	engine.GET("/email_account_exist", api.isEmailAccountExistHandler) // 是否邮箱账号已存在
	engine.GET("/phone_account_exist", api.isPhoneAccountExistHandler) // 是否手机账号已存在
	engine.POST("/register_email", api.registerEmailHandler)           // 邮箱注册
	engine.POST("/register_phone", api.registerPhoneHandler)           // 手机号码注册
	engine.POST("/login_email", api.loginEmailHandler)                 // 邮箱登录
	engine.POST("/login_phone", api.loginPhoneHandler)                 // 手机号码登录
	engine.GET("/login_qq", api.qqOAuthLoginHandler)                   // qq授权登录
	engine.GET("/login_wechat", api.wechatOAuthLoginHandler)           // 微信授权登录
	engine.GET("/login_weibo", api.weiboOAuthLoginHandler)             // 微博授权登录
	engine.GET("/send_mail_verify_code", api.sendEmailVerifyCode)      // 发送邮箱验证码 for 1校验邮箱，2忘记密码重设邮件
	engine.GET("/send_sms_verify_code", api.sendSmsVerifyCode)         // 发送手机短信验证码,for 1注册、2登录、3绑定
	engine.POST("/reset_forget_password", api.resetForgetPassword)     // 忘记密码，接收到验证码后重设密码

	// 用户鉴权访问路由组接口
	authGroup := engine.Group("/user", middleware.JwtAuthMiddleware())
	authGroup.GET("/info", api.getUserInfoHandler)                  // 获取用户信息
	authGroup.POST("/set_info", api.setUserInfoHandler)             // 设置用户信息
	authGroup.POST("/set_avatar", api.setAvatarHandler)             // 设置用户头像
	authGroup.POST("/email_validate", api.emailValidate)            // 登录后校验邮箱地址，接收邮箱验证码后，账号邮箱验证码校验
	authGroup.POST("/email_binding", api.emailBinding)              // 绑定未注册邮箱
	authGroup.POST("/phone_binding", api.phoneBinding)              // 绑定未注册手机
	authGroup.POST("/modified_password", api.modifiedPassword)      // 修改密码
	authGroup.GET("/qq_binding", api.qqOAuthBindingHandler)         // 登录后绑定qq账号
	authGroup.GET("/weixin_binding", api.wechatOAuthBindingHandler) // 登录后绑定微信账号
	authGroup.GET("/weibo_binding", api.weiboOAuthBindingHandler)   // 登录后绑定微博账号
	authGroup.GET("/logout", api.logoutHandler)                     // 退出登录

}

/* 是否邮箱账号已存在 */
func (api *UserApi) isEmailAccountExistHandler(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.IsEmailAccountExistDTO
	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	isExist, err := userService.IsEmailAccountExist(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"isExist": isExist}))

}

/* 是否手机账号已存在 */
func (api *UserApi) isPhoneAccountExistHandler(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.IsPhoneAccountExistDTO
	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	userService := services.GetUserService()
	isExist, err := userService.IsPhoneAccountExist(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"isExist": isExist}))

}

/*邮箱注册*/
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
	userInfoDTO, err := userService.CreateUserWithEmail(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(userInfoDTO))
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
	userInfoDTO, err := userService.CreateUserWithPhone(dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(userInfoDTO))
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

/*qq oauth 登录*/
func (api *UserApi) qqOAuthLoginHandler(ctx *gin.Context) {
	// Receive Request ...
	var dto dtos.QQLoginDTO

	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	thirdOAuthService := services.GetThirdOAuthService()
	token, err := thirdOAuthService.QQOAuthLogin(dto)
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

	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	thirdOAuthService := services.GetThirdOAuthService()
	token, err := thirdOAuthService.WechatOAuthLogin(dto)
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

	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Call Services method ...
	thirdOAuthService := services.GetThirdOAuthService()
	token, err := thirdOAuthService.WeiboOAuthLogin(dto)
	if err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"token": token}))

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
	thirdOAuthService := services.GetThirdOAuthService()
	if ok, err = thirdOAuthService.QQOAuthBinding(dto); err != nil {
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
	thirdOAuthService := services.GetThirdOAuthService()
	if ok, err = thirdOAuthService.WechatOAuthBinding(dto); err != nil {
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
	thirdOAuthService := services.GetThirdOAuthService()
	if ok, err = thirdOAuthService.WeiboOAuthBinding(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	if ok {
		ctx.Set(common.ResponseDataKey, nil)
	} else {
		_ = ctx.Error(errors.New("Weibo OAuth Binding Fail "))
	}

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

// 验证账号邮箱
func (api *UserApi) emailValidate(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.EmailValidateDTO

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
	if isPass, err := userService.EmailValidate(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

// 绑定账号邮箱
func (api *UserApi) emailBinding(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.EmailValidateDTO

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
	if isPass, err := userService.EmailBinding(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

// 绑定账号手机号码
func (api *UserApi) phoneBinding(ctx *gin.Context) {
	// Receive Request ...
	var userId uint
	var dto dtos.PhoneValidateDTO

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
	if isPass, err := userService.PhoneBinding(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
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
	var dto dtos.ResetForgetPasswordDTO

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userService := services.GetUserService()
	if isPass, err := userService.ResetForgetPassword(dto); !isPass || err != nil {
		_ = ctx.Error(err)
		return
	}

	// Send Data to Response Middleware ...
	ctx.Set(common.ResponseDataKey, nil)
}

func (api *UserApi) sendEmailVerifyCode(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.SendEmailVerifyCodeDTO
	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	mailService := services.GetMailService()
	if err = mailService.SendEmailVerifyCode(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

func (api *UserApi) sendSmsVerifyCode(ctx *gin.Context) {
	// 接收参数由dto封装验证
	var dto dtos.SendPhoneVerifyCodeDTO
	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	smsService := services.GetSmsService()
	if err = smsService.SendPhoneVerifyCode(dto); err != nil {
		_ = ctx.Error(err) // 所有错误最后传递给错误中间件处理
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}
