package ai_project_result

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idstr string, isSkipHooks bool) (aiProjectResult AiProjectResult) {
	if isSkipHooks {
		database.SkipHookDB.Preload(clause.Associations).Where("id", idstr).First(&aiProjectResult)
	} else {
		database.DB.Preload(clause.Associations).Where("id", idstr).First(&aiProjectResult)
	}
	return
}

func GetBy(field, value string) (aiProjectResult AiProjectResult) {
	database.DB.Preload(clause.Associations).Where(field, value).First(&aiProjectResult)
	return
}

func GetByUUID(uuid string) (aiProjectResults []AiProjectResult) {
	database.DB.Preload(clause.Associations).Where(`ai_project_uuid = ?`, uuid).Find(&aiProjectResults)
	return
}

func All() (aiProjectResults []AiProjectResult) {
	database.DB.Preload(clause.Associations).Find(&aiProjectResults)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AiProjectResult{}).Where(field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, preloadFields []paginator.BasePreloadField,
	whereFields []paginator.BaseWhereField, perPage int) (aiProjectResults []AiProjectResult, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AiProjectResult{}),
		preloadFields,
		whereFields,
		&aiProjectResults,
		app.V1URL(database.TableName(&AiProjectResult{})),
		perPage,
	)
	return
}
