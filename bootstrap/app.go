package bootstrap

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/config"
)

// SetupLogger 初始化 Logger
func SetupApp() {

	app.SetAppLanguage(
		config.GetString("translator.language"),
	)
}
