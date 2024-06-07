package download

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (download Download) {
	database.DB.Where("id", idstr).First(&download)
	return
}

func GetBy(field, value string) (download Download) {
	database.DB.Where("? = ?", field, value).First(&download)
	return
}

func All() (downloads []Download) {
	database.DB.Find(&downloads)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Download{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (downloads []Download, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Download{}),
		whereFields,
		&downloads,
		app.V1URL(database.TableName(&Download{})),
		perPage,
	)
	return
}
