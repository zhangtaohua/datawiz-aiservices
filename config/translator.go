// Package config 站点配置信息
package config

import "datawiz-aiservices/pkg/config"

func init() {
	config.Add("translator", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认翻译语言
			"language": config.Env("TRANSLATER_LANGUAGE", "en"),

			// i18n 配置目录
			"dir": config.Env("TRANSLATER_DIR", "./assets/i18n/"),
		}
	})
}
