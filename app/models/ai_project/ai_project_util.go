package ai_project

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string, isSkipHooks bool) (aiProject AiProject) {
	if isSkipHooks {
		database.SkipHookDB.Where("id", idstr).First(&aiProject)
	} else {
		database.DB.Where("id", idstr).First(&aiProject)
	}
	return
}

func GetBy(field, value string) (aiProject AiProject) {
	database.DB.Where(field, value).First(&aiProject)
	return
}

func All() (aiProjects []AiProject) {
	database.DB.Find(&aiProjects)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AiProject{}).Where(field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (aiProjects []AiProject, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AiProject{}),
		whereFields,
		&aiProjects,
		app.V1URL(database.TableName(&AiProject{})),
		perPage,
	)
	return
}
