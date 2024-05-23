// Package bootstrap 启动程序功能
package bootstrap

import (
	"datawiz-aiservices/pkg/config"
	"datawiz-aiservices/pkg/translator"
)

// SetupCache 缓存
func SetupTranslator() {

	translator.NewTranslator(
		config.GetString("translator.language"),
		config.GetString("translator.dir"),
	)
}
