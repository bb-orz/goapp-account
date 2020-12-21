package main

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

var app *goinfras.Application // 应用实例

func main() {
	// 初始化Viper配置加载器，导入配置，启动参数由命令行flag输入
	fmt.Println("Viper Config Loading  ......")
	viperCfg := goinfras.ViperLoader()

	// 注册应用组件启动器
	fmt.Println("Register Starters  ......")
	RegisterStarter()

	// 创建应用程序启动管理器
	app = goinfras.NewApplication(viperCfg)

	// 运行应用,启动已注册的资源组件
	fmt.Println("Application Starting ......")
	app.Up()
}
