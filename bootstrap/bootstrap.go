package bootstrap

import (
	"datawiz-aiservices/pkg/config"
	"datawiz-aiservices/pkg/console"
	"datawiz-aiservices/pkg/logger"
	"embed"
)

func Bootstrap(databaseMigrationFS embed.FS) {
	SetupApp()

	// 初始化 Logger
	SetupLogger()

	// 初始化数据库
	SetupDB()

	// 初始化 Redis
	SetupRedis()

	// 初始化数据表
	RunMigration(databaseMigrationFS)

	// 初始化预置数据
	RunSeed()

	// 初始化缓存
	SetupCache()

	// 初始化i18n 翻译
	SetupTranslator()

	// 初始化 自定义较验规则
	SetupValidators()

	// 初始化路由绑定
	router := SetupRoute()

	// 运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("Main", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
