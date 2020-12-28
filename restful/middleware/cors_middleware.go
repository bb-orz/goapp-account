package middleware

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 跨域请求处理策略，cors/jsonp/
func CorsMiddleware() gin.HandlerFunc {
	// 获取应用全局配置实例
	viperConfig := goinfras.XApp().Sctx.Configs()

	var config cors.Config
	err := viperConfig.UnmarshalKey("GinCors", &config)
	if err != nil {
		XLogger.XCommon().Error("Gin Cors Middleware Error", zap.Error(err))
	}

	return cors.New(config)
}
