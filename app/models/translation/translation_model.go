// Package Translation 模型
package translation

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/database"
)

type Translation struct {
	models.BaseModel

	// Put fields in here
	TranslationId  string `json:"translation_id"`
	Language       string `json:"language"`
	TranslatedText string `json:"translated_text"`
}

func (translation *Translation) Create() {
	database.DB.Create(&translation)
}

func (translation *Translation) Save() (rowsAffected int64) {
	result := database.DB.Save(&translation)
	return result.RowsAffected
}

func (translation *Translation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&translation)
	return result.RowsAffected
}
