package middleware

import (
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 跨域请求处理策略，cors/jsonp/
func CorsMiddleware(viperConfig *viper.Viper) gin.HandlerFunc {
	// 获取应用全局配置实例
	var config cors.Config
	err := viperConfig.UnmarshalKey("GinCors", &config)
	if err != nil {
		XLogger.XCommon().Error("Gin Cors Middleware Error", zap.Error(err))
	}

	if config.AllowAllOrigins {
		return cors.New(cors.Config{AllowAllOrigins: true})
	}

	return cors.New(config)
}
