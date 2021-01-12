package user

import (
	"github.com/bb-orz/goinfras/XCache"
)

type UsersCache struct {
	CommonCache XCache.ICommonCache
}

func NewUserCache() *UsersCache {
	cache := new(UsersCache)
	cache.CommonCache = XCache.XCommon()
	return cache
}

// TODO Some User Cache Biz
