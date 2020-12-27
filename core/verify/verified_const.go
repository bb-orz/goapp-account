package verify

// 用户邮箱验证
const (
	DomainName                       = "VerifyDomain"
	UserCacheVerifyEmailCodePrefix = "user.verify.email.code." // 缓存邮箱验证码key前缀
	UserCacheVerifyEmailCodeExpire = 60 * 60 * 3                 // 缓存邮箱验证码超时时间

	UserCacheVerifyPhoneCodePrefix = "user.verify.phone.code." // 缓存手机验证码key前缀
	UserCacheVerifyPhoneCodeExpire = 60 * 5                      // 缓存手机验证码超时时间

	UserCacheForgetPasswordVerifyCodePrefix = "user.verify.forgetpassword.code." // 忘记密码重置验证码前缀
	UserCacheForgetPasswordVerifyCodeExpire = 60 * 5                               // 忘记密码重置超时时间

)
