package user

import (
	"github.com/bb-orz/goinfras/XCache"
)

type UserCache struct {
	CommonCache XCache.ICommonCache
}

func NewUserCache() *UserCache {
	cache := new(UserCache)
	cache.CommonCache = XCache.XCommon()
	return cache
}

// TODO Some User Cache Biz
