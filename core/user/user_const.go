package user

const (
	DomainName = "UserDomain"
	// 用户状态相关
	UserStatusNotVerified  = iota // 未验证 0
	UserStatusNormal              // 已验证 1
	UserStatusDeactivation        // 已停用 2

)

// 用户绑定的第三方账号平台
const (
	QQOauthPlatform = iota
	WechatOauthPlatform
	WeiboOauthPlatform
)

const (
	UserCacheTokenPrefix = "user.token."    // 用户鉴权令牌缓存前缀
	UserCacheTokenExpire = 60 * 60 * 23 * 7 // 用户鉴权令牌缓存超时
)
