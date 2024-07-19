// Package auth 授权相关逻辑
package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	jwtPayload := c.GetHeader("X-Jwt-Payload")
	fmt.Printf("jwtPayload %v", jwtPayload)
	userID := c.GetHeader("X-User-Id")
	return userID
}
