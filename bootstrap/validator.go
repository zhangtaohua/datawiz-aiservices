// Package bootstrap 启动程序功能
package bootstrap

import (
	"datawiz-aiservices/app/requests/validators"
)

func SetupValidators() {
	validators.Initialize()
}
