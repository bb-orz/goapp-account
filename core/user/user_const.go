package user

const (
	DomainName = "UserDomain"
	// 用户状态相关
	UserStatusNotVerify    = 0 // 未验证 0
	UserStatusNormal       = 1 // 已验证 1
	UserStatusDeactivation = 2 // 已停用 2

)

const (
	UserEmailNotVerify = 0
	UserEmailVerify    = 1
	UserPhoneNotVerify = 0
	UserPhoneVerify    = 1
)

// 用户绑定的第三方账号平台
const (
	QQOauthPlatform = iota
	WechatOauthPlatform
	WeiboOauthPlatform
)
