// Package bootstrap 启动程序功能
package bootstrap

import (
	"datawiz-aiservices/database/migrations"
	"datawiz-aiservices/pkg/migrate"
	"os"
)

// SetupCache 缓存
func RunMigration() {
	// 初始化数据表
	length := len(os.Args)
	if length >= 3 {
		command := os.Args[1]
		if command == "migrate" {
			subCmd := os.Args[2]
			switch subCmd {
			// 执行迁移
			case "up":
				runUp()
				break
			// 回滚上一步执行的迁移
			case "down":
				runDown()
				break
			// 回滚所有迁移
			case "reset":
				runReset()
				break
			// 回滚所有迁移，然后再执行所有迁移
			case "refresh":
				runRefresh()
				break
			// 删除所有表，然后执行所有迁移
			case "fresh":
				runFresh()
				break
			}
		}
	}

}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp() {
	migrator().Up()
}

func runDown() {
	migrator().Rollback()
}

func runReset() {
	migrator().Reset()
}

func runRefresh() {
	migrator().Refresh()
}

func runFresh() {
	migrator().Fresh()
}
