package ai_project_result

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (aiProjectResult AiProjectResult) {
	database.DB.Where("id", idstr).First(&aiProjectResult)
	return
}

func GetBy(field, value string) (aiProjectResult AiProjectResult) {
	database.DB.Where("? = ?", field, value).First(&aiProjectResult)
	return
}

func GetByUUID(uuid string) (aiProjectResults []AiProjectResult) {
	database.DB.Where(`ai_project_uuid = ?`, uuid).Find(&aiProjectResults)
	return
}

func All() (aiProjectResults []AiProjectResult) {
	database.DB.Find(&aiProjectResults)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AiProjectResult{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (aiProjectResults []AiProjectResult, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AiProjectResult{}),
		whereFields,
		&aiProjectResults,
		app.V1URL(database.TableName(&AiProjectResult{})),
		perPage,
	)
	return
}
