package ai_model

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string, isSkipHooks bool) (aiModel AiModel) {
	if isSkipHooks {
		database.SkipHookDB.Where("id", idstr).First(&aiModel)
	} else {
		database.DB.Where("id", idstr).First(&aiModel)
	}
	return
}

func GetBy(field, value string) (aiModel AiModel) {
	database.DB.Where(field, value).First(&aiModel)
	return
}

func All() (aiModels []AiModel) {
	database.DB.Find(&aiModels)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AiModel{}).Where(field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, preloadFields []paginator.BasePreloadField,
	whereFields []paginator.BaseWhereField, perPage int) (aiModels []AiModel, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AiModel{}),
		preloadFields,
		whereFields,
		&aiModels,
		app.V1URL(database.TableName(&AiModel{})),
		perPage,
	)
	return
}
