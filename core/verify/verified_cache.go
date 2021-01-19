package verify

import (
	"fmt"
	"github.com/bb-orz/goinfras/XCache"
	"strconv"
)

type verifyCache struct {
	commonCache XCache.ICommonCache
}

func NewMailCache() *verifyCache {
	cache := new(verifyCache)
	cache.commonCache = XCache.XCommon()
	return cache
}

// 保存邮箱验证码缓存
func (cache *verifyCache) SetForgetPasswordVerifyCode(uid uint, code string) error {
	key := UserCacheForgetPasswordVerifyCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp(key, code, UserCacheForgetPasswordVerifyCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取邮箱验证码缓存
func (cache *verifyCache) GetForgetPasswordVerifyCode(uid uint) (string, bool) {
	key := UserCacheForgetPasswordVerifyCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get(key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s", reply)
	return code, false
}

// 保存邮箱验证码缓存
func (cache *verifyCache) SetUserVerifyEmailCode(uid uint, code string) error {
	key := UserCacheVerifyEmailCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp(key, code, UserCacheVerifyEmailCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取邮箱验证码缓存
func (cache *verifyCache) GetUserVerifyEmailCode(uid uint) (string, bool) {
	key := UserCacheVerifyEmailCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get(key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s", reply)
	return code, true
}

// 保存手机验证码缓存
func (cache *verifyCache) SetUserVerifyPhoneCode(uid uint, code string) error {
	key := UserCacheVerifyPhoneCodePrefix + strconv.Itoa(int(uid))
	err := cache.commonCache.SetWithExp(key, code, UserCacheVerifyPhoneCodeExpire)
	if err != nil {
		return err
	}
	return nil
}

// 获取手机验证码缓存
func (cache *verifyCache) GetUserVerifyPhoneCode(uid uint) (string, bool) {
	key := UserCacheVerifyPhoneCodePrefix + strconv.Itoa(int(uid))
	reply, isExist := cache.commonCache.Get(key)
	if !isExist {
		return "", false
	}
	code := fmt.Sprintf("%s", reply)
	return code, true
}
