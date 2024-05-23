// Package ai_project 模型
package ai_project

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/helpers"
)

type AiProject struct {
	models.BaseModel
	models.BaseUUIDModel

	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`

	models.CommonTimestampsField
}

func (aiProject *AiProject) Create() {
	database.DB.Create(&aiProject)
}

func (aiProject *AiProject) Save() (rowsAffected int64) {
	result := database.DB.Save(&aiProject)
	return result.RowsAffected
}

func (aiProject *AiProject) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&aiProject)
	return result.RowsAffected
}

func (aiProject *AiProject) CreateTx(request *requests.AiProjectRequest) error {
	// 生成UUID作为 translated_id
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 生成UUID作为 translated_id

	nameUUID := helpers.UUID()
	descUUID := helpers.UUID()

	aiProject.Name = nameUUID
	aiProject.Description = descUUID
	if err := tx.Create(&aiProject).Error; err != nil {
		tx.Rollback()
		return err
	}

	nameTranslationModel := translation.Translation{
		TranslationId:  nameUUID,
		Language:       request.Language,
		TranslatedText: request.Name,
	}
	if err := tx.Create(&nameTranslationModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	descTranslationModel := translation.Translation{
		TranslationId:  descUUID,
		Language:       request.Language,
		TranslatedText: request.Description,
	}
	if err := tx.Create(&descTranslationModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
