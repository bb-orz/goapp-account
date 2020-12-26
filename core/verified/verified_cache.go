package verified

import (
	"fmt"
	"github.com/bb-orz/goinfras/XCache"
	"strconv"
)

type verifiedCache struct {
	commonCache XCache.ICommonCache
}

func NewMailCache() *verifiedCache {
	cache := new(verifiedCache)
	cache.commonCache = XCache.XCommon()
	return cache
}

// 保存邮箱验证码缓存
func (cache *verifiedCache) SetForgetPasswordVerifiedCode(uid uint, code string) error {
	key := UserCacheForgetPasswordVerifiedCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp( key, code,UserCacheForgetPasswordVerifiedCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取邮箱验证码缓存
func (cache *verifiedCache) GetForgetPasswordVerifiedCode(uid uint) (string, bool) {
	key := UserCacheForgetPasswordVerifiedCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get(key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s",reply)
	return code, false
}

// 保存邮箱验证码缓存
func (cache *verifiedCache) SetUserVerifiedEmailCode(uid uint, code string) error {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp( key,code, UserCacheVerifiedEmailCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取邮箱验证码缓存
func (cache *verifiedCache) GetUserVerifiedEmailCode(uid uint) (string, bool) {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get( key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s",reply)
	return code, false}

// 保存手机验证码缓存
func (cache *verifiedCache) SetUserVerifiedPhoneCode(uid uint, code string) error {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp(key,code, UserCacheVerifiedPhoneCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取手机验证码缓存
func (cache *verifiedCache) GetUserVerifiedPhoneCode(uid uint) (string, bool) {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get(key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s", reply)
	return code, true
}
