package translation

import (
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (translation Translation) {
	database.DB.Where("id", idstr).First(&translation)
	return
}

func GetBy(field, value string) (translation Translation) {
	database.DB.Where("? = ?", field, value).First(&translation)
	return
}

func GetByTidLang(translationId, language string) (translation Translation) {
	database.DB.Where(`translation_id = ? AND language = ? `, translationId, language).First(&translation)
	return
}

func GetT(translationId, language string) (msg string) {
	var translation Translation
	database.DB.Model(Translation{}).Where(`translation_id = ? AND language = ? `, translationId, language).Find(&translation)
	msg = translation.TranslatedText
	return
}

func GetTs(translationIds []string, language string) map[string]string {
	var translations []Translation
	length := len(translationIds)
	db := database.DB.Select("translation_id", "translated_text")

	for i := 0; i < length; i++ {
		if i == 0 {
			db = db.Where(`translation_id = ? AND language = ?  `, translationIds[i], language)
		} else {
			db = db.Or(`translation_id = ? AND language = ? `, translationIds[i], language)
		}
	}
	db.Find(&translations)
	msg := make(map[string]string)

	for _, v := range translations {
		msg[v.TranslationId] = v.TranslatedText
	}
	return msg
}

func All() (translation []Translation) {
	database.DB.Find(&translation)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Translation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func IsExistUnion(translationId, language string) bool {
	var count int64
	database.DB.Model(Translation{}).Where(`translation_id = ? AND language = ? `, translationId, language).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (translation []Translation, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Translation{}),
		whereFields,
		&translation,
		app.V1URL(database.TableName(&Translation{})),
		perPage,
	)
	return
}
