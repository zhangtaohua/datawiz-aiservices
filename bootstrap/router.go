// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"net/http"
	"strings"

	"datawiz-aiservices/app/http/middlewares"
	"datawiz-aiservices/pkg/translator"
	"datawiz-aiservices/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)
	// gin 实例
	router = gin.New()

	// 注册全局中间件
	registerGlobalMiddleWare()

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler()

	return router
}

func registerGlobalMiddleWare() {
	router.Use(
		// gin.Logger(),
		middlewares.Logger(),
		// gin.Recovery(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setup404Handler() {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, translator.TransHandler.T("r.noPageErr"))
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": translator.TransHandler.T("r.noRouterErr"),
			})
		}
	})
}
