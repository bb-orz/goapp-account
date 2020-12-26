package user

import (
	"github.com/bb-orz/goinfras/XCache"
)

type userCache struct {
	CommonCache XCache.ICommonCache
}

func NewUserCache() *userCache {
	cache := new(userCache)
	cache.CommonCache = XCache.XCommon()
	return cache
}


// TODO Some User Cache Biz