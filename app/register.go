package main

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XGin"
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XStore/XMongo"
	"github.com/bb-orz/goinfras/XStore/XRedis"
	"github.com/bb-orz/goinfras/XValidate"
	"goinfras-sample-account/restful/middleware"
	_ "goinfras-sample-account/restful" // 自动载入Restful API模块
	_ "goinfras-sample-account/core"
)

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter() {
	goinfras.RegisterStarter(XGlobal.NewStarter())

	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册mongodb启动器
	goinfras.RegisterStarter(XMongo.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())
	// 注册Redis连接池
	goinfras.RegisterStarter(XRedis.NewStarter())

	// 注册Oauth Manager
	goinfras.RegisterStarter(XOAuth.NewStarter())

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 注册gin web 服务启动器
	// TODO add your gin middlewares
	goinfras.RegisterStarter(XGin.NewStarter(middleware.CorsMiddleware(),middleware.ErrorMiddleware()))

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
