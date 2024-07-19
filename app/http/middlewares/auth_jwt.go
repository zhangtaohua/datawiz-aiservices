// Package middlewares Gin 中间件
package middlewares

import (
	"datawiz-aiservices/pkg/helpers"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.GetHeader("X-User-Id")
		if helpers.Empty(userID) {
			response.Unauthorized(c, translator.TransHandler.T("r.authErr"))
			return
		}

		c.Next()
	}
}
