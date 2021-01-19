package verify

// 用户邮箱验证
const (
	DomainName = "VerifyDomain"

	EmailVCTypeCheck       = 1                               // 检测邮箱验证码类型
	EmailVCTypeCheckPrefix = "user.email.check.verify.code." // 缓存邮箱验证码key前缀
	EmailVCTypeCheckExpire = 60 * 60 * 3                     // 缓存邮箱验证码超时时间

	EmailVCTypeForgetPassword       = 2                                        // 忘记密码邮箱验证码类型
	EmailVCTypeForgetPasswordPrefix = "user.email.forgetpassword.verify.code." // 忘记密码重置验证码前缀
	EmailVCTypeForgetPasswordExpire = 60 * 5                                   // 忘记密码重置超时时间

	PhoneVCTypeRegister       = 1                                  // 手机注册短信验证码类型
	PhoneVCTypeRegisterPrefix = "user.phone.register.verify.code." // 缓存手机注册验证码key前缀
	PhoneVCTypeRegisterExpire = 60 * 5                             // 缓存手机注册验证码超时时间

	PhoneVCTypeLogin       = 2                               // 手机登录短信验证码类型
	PhoneVCTypeLoginPrefix = "user.phone.login.verify.code." // 缓存手机登录验证码key前缀
	PhoneVCTypeLoginExpire = 60 * 5                          // 缓存手机登录验证码超时时间

	PhoneVCTypeBinding       = 3                                 // 手机绑定验证码类型
	PhoneVCTypeBindingPrefix = "user.phone.binding.verify.code." // 缓存手机绑定验证码key前缀
	PhoneVCTypeBindingExpire = 60 * 5                            // 缓存手机绑定验证码超时时间

)
