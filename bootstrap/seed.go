// Package bootstrap 启动程序功能
package bootstrap

import (
	"datawiz-aiservices/database/seeders"
	"datawiz-aiservices/pkg/console"
	"datawiz-aiservices/pkg/seed"
	"os"
)

// SetupCache 缓存
func RunSeed() {
	// 初始化数据表
	length := len(os.Args)
	if length >= 2 {
		command := os.Args[1]
		if command == "seed" {
			var args []string
			if length >= 3 {
				args = os.Args[2:]
			} else {
				args = []string{}
			}
			runSeeders(args)
		}
	}
}

func runSeeders(args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
