package common

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"time"
)

func GetUserClaim(ctx *gin.Context) *XJwt.UserClaim {
	claim, exists := ctx.Get(ContextTokenUserClaimKey)
	if exists {
		userClaim := claim.(XJwt.UserClaim)
		return &userClaim
	} else {
		return nil
	}
}

func GetUserId(ctx *gin.Context) (uint, error) {
	claim, exists := ctx.Get(ContextTokenUserClaimKey)
	if exists {
		userClaim := claim.(XJwt.UserClaim)
		return userClaim.Id, nil
	} else {
		return 0, errors.New("User Token Not Exist! ")
	}
}

func GetTokenString(ctx *gin.Context) string {
	return ctx.GetString(ContextTokenStringKey)
}

// 幂运算
func Powerf(x float64, n int) float64 {
	ans := 1.0
	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

// 生成固定位随机数字
func RandomNumber(l int) (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(int32(Powerf(10.00, l))))
	return code, nil
}

// 生成固定长度随机字符串
func RandomString(l int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	var rendomString []byte
	bytes := []byte(chars)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		rendomString = append(rendomString, bytes[r.Intn(len(bytes))])
	}
	return string(rendomString)
}

// 用户密码加盐生成哈希
func HashPassword(src string) (hashStr, salt string) {
	// 获取随机盐值字符串
	salt = RandomString(4)

	hash := sha1.New()

	_, _ = io.WriteString(hash, src)
	_, _ = io.WriteString(hash, salt)

	hashBytes := hash.Sum(nil)
	// 组合输出40位哈希字符
	hashStr = fmt.Sprintf("%x", hashBytes)
	return
}

// 校验密码
func ValidatePassword(passStr, salt, passHash string) bool {
	// 重新计算密码哈希，与之前的校验
	hash := sha1.New()
	_, _ = io.WriteString(hash, passStr)
	_, _ = io.WriteString(hash, salt)
	hashBytes := hash.Sum(nil)
	// 组合输出40位哈希字符
	hashStr := fmt.Sprintf("%x", hashBytes)

	return hashStr == passHash
}

func IsStringItemExist(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
