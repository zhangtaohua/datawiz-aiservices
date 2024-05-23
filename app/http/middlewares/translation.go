package middlewares

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/translator"
	"strings"

	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func Translation() gin.HandlerFunc {
	return func(c *gin.Context) {

		language := c.GetHeader("X-Language") // 获取请求头中的语言设置，默认为英文
		language = strings.ToLower(language)
		if strings.HasPrefix(language, "en") {
			language = "en"
		} else {
			switch language {
			case "zh", "zh_cn", "zh-cn", "zhcn", "zh_ch", "zh-ch", "zhch", "zh_chn", "zh-chn", "zhchn":
				language = "zh-CN"
				break

			case "zh_tw", "zh-tw", "zhtw":
				language = "zh-TW"
				break

			default:
				language = "en"
				break
			}
		}

		app.SetAppLanguage(language)
		translator.TransHandler.SetTranslatorLanguage(language)

		c.Next()
	}
}
