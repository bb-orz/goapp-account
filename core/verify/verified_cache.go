package verify

import (
	"fmt"
	"github.com/bb-orz/goinfras/XCache"
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
func (cache *verifyCache) SetEmailVerifyCode(vcType uint, email string, code string) error {
	var key string
	var err error

	switch vcType {
	case EmailVCTypeCheck:
		key = EmailVCTypeCheckPrefix + email
		err = cache.commonCache.SetWithExp(key, code, EmailVCTypeCheckExpire)
		if err != nil {
			return err
		}
	case EmailVCTypeForgetPassword:
		key = EmailVCTypeForgetPasswordPrefix + email
		err = cache.commonCache.SetWithExp(key, code, EmailVCTypeForgetPasswordExpire)
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取邮箱验证码缓存
func (cache *verifyCache) GetEmailVerifyCode(vcType uint, email string) (string, bool) {
	var key string
	var code string
	switch vcType {
	case EmailVCTypeCheck:
		key = EmailVCTypeCheckPrefix + email
		reply, isExist := cache.commonCache.Get(key)
		if !isExist {
			return "", false
		}
		code = fmt.Sprintf("%s", reply)
		return code, true
	case EmailVCTypeForgetPassword:
		key = EmailVCTypeForgetPasswordPrefix + email
		reply, isExist := cache.commonCache.Get(key)
		if !isExist {
			return "", false
		}
		code = fmt.Sprintf("%s", reply)
		return code, true
	}

	return "", false
}

// 保存手机验证码缓存
func (cache *verifyCache) SetUserPhoneVerifyCode(vcType uint, phone string, code string) error {
	var key string
	var err error

	switch vcType {
	case PhoneVCTypeRegister:
		key = PhoneVCTypeRegisterPrefix + phone
		err = cache.commonCache.SetWithExp(key, code, PhoneVCTypeRegisterExpire)
		if err != nil {
			return err
		}
	case PhoneVCTypeLogin:
		key = PhoneVCTypeLoginPrefix + phone
		err = cache.commonCache.SetWithExp(key, code, PhoneVCTypeLoginExpire)
		if err != nil {
			return err
		}
	case PhoneVCTypeBinding:
		key = PhoneVCTypeBindingPrefix + phone
		err = cache.commonCache.SetWithExp(key, code, PhoneVCTypeBindingExpire)
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取手机验证码缓存
func (cache *verifyCache) GetUserPhoneVerifyCode(vcType uint, phone string) (string, bool) {
	var key string
	var code string
	switch vcType {
	case PhoneVCTypeRegister:
		key = PhoneVCTypeRegisterPrefix + phone
		reply, isExist := cache.commonCache.Get(key)
		if !isExist {
			return "", false
		}
		code = fmt.Sprintf("%s", reply)
		return code, true
	case PhoneVCTypeLogin:
		key = PhoneVCTypeLoginPrefix + phone
		reply, isExist := cache.commonCache.Get(key)
		if !isExist {
			return "", false
		}
		code = fmt.Sprintf("%s", reply)
		return code, true
	case PhoneVCTypeBinding:
		key = PhoneVCTypeBindingPrefix + phone
		reply, isExist := cache.commonCache.Get(key)
		if !isExist {
			return "", false
		}
		code = fmt.Sprintf("%s", reply)
		return code, true
	}

	return "", false
}
