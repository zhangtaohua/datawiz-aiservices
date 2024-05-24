// Package routes 注册路由
package routes

import (
	controllers "datawiz-aiservices/app/http/controllers/api/v1"
	"datawiz-aiservices/app/http/middlewares"
	"datawiz-aiservices/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1/ai")
	} else {
		v1 = r.Group("/v1/ai")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("2000-H"), middlewares.Translation())
	{
		v1.StaticFS("/uploads", http.Dir("./public/uploads"))
		// v1.Static("/uploads", "./public/uploads")

		tc := new(controllers.TranslationsController)
		tcGroup := v1.Group("/translations")
		{
			tcGroup.GET("", tc.Index)
			tcGroup.GET("/:id", tc.Show)
			tcGroup.POST("", tc.Store)
			tcGroup.PUT("/:id", tc.Update)
			tcGroup.DELETE("/:id", tc.Delete)
		}

		amc := new(controllers.AiModelsController)
		amcGroup := v1.Group("/models")
		{
			amcGroup.GET("", amc.Index)
			amcGroup.GET("/:id", amc.Show)
			amcGroup.POST("", amc.Store)
			amcGroup.PUT("/:id", amc.Update)
			amcGroup.DELETE("/:id", amc.Delete)
		}

		apc := new(controllers.AiProjectsController)
		apcGroup := v1.Group("/projects")
		{
			apcGroup.GET("", apc.Index)
			apcGroup.GET("/:id", apc.Show)
			apcGroup.POST("", apc.Store)
			apcGroup.PUT("/:id", apc.Update)
			apcGroup.DELETE("/:id", apc.Delete)
		}

		aprc := new(controllers.AiProjectResultsController)
		aprcGroup := v1.Group("/project/results")
		{
			aprcGroup.GET("", aprc.Index)
			aprcGroup.GET("/:id", aprc.Show)
			aprcGroup.POST("", aprc.Store)
			aprcGroup.PUT("/:id", aprc.Update)
			aprcGroup.DELETE("/:id", aprc.Delete)
		}
		// end V1
	}
}
