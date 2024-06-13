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
	database.DB.Where(field, value).First(&translation)
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

func TryGetT(translationId, language string) (msg string) {
	// todo
	// 建议从一张表中获取，估计还要创建一张表
	languageList := []string{"zh-CN", "zh-TW", "en"}

	var translation Translation
	database.DB.Model(Translation{}).Where(`translation_id = ? AND language = ? `, translationId, language).Find(&translation)
	if translation.ID == 0 {
		for _, tryLanguage := range languageList {
			if tryLanguage != language {
				database.DB.Model(Translation{}).Where(`translation_id = ? AND language = ? `, translationId, tryLanguage).Find(&translation)
				if translation.ID != 0 {
					break
				}
			}
		}
	}

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
	msgs := make(map[string]string)

	for _, v := range translations {
		msgs[v.TranslationId] = v.TranslatedText
	}
	return msgs
}

// TryGetTs 版本V1 查询次数很多，最后结果可能会是不同的语言的组合，且也有可能查不到。
func TryGetTsV1(translationIds []string, language string) map[string]string {
	msgs := make(map[string]string)
	length := len(translationIds)

	for i := 0; i < length; i++ {
		msg := TryGetT(translationIds[i], language)
		msgs[translationIds[i]] = msg
	}
	return msgs
}

// TryGetTs 版本V2  查询数据库次数为尝试语言次数。
// 最后可能会有部分项是找不到对应的翻译语言的，除非数据保证一种语言是一定全存在的。
func TryGetTsV2(translationIds []string, language string) map[string]string {
	var translations []Translation
	length := len(translationIds)

	// todo
	// 建议从一张表中获取，估计还要创建一张表
	languageList := []string{"zh-CN", "zh-TW", "en"}

	// default
	db := database.DB.Select("translation_id", "translated_text")
	for i := 0; i < length; i++ {
		if i == 0 {
			db = db.Where(`translation_id = ? AND language = ?  `, translationIds[i], language)
		} else {
			db = db.Or(`translation_id = ? AND language = ? `, translationIds[i], language)
		}
	}
	db.Find(&translations)

	translationsLength := len(translations)

	// 请求默认值没有结果，就按顺序查找
	if translationsLength != length {
		for _, tryLanguage := range languageList {
			tryDb := database.DB.Select("translation_id", "translated_text")
			for i := 0; i < length; i++ {
				if i == 0 {
					tryDb = tryDb.Where(`translation_id = ? AND language = ?  `, translationIds[i], tryLanguage)
				} else {
					tryDb = tryDb.Or(`translation_id = ? AND language = ? `, translationIds[i], tryLanguage)
				}
			}
			tryDb.Find(&translations)

			tryTranslationsLength := len(translations)
			// 严格来说是要等于
			if tryTranslationsLength == length {
				break
			}
		}
	}

	// finally
	msgs := make(map[string]string)
	for _, v := range translations {
		msgs[v.TranslationId] = v.TranslatedText
	}
	return msgs
}

func All() (translation []Translation) {
	database.DB.Find(&translation)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Translation{}).Where(field, value).Count(&count)
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
