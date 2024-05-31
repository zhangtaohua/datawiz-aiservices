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

	return tx.Commit().Error
}

func (aiProject *AiProject) SaveTx(request *requests.AiProjectRequest, isUpdate bool) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&aiProject).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 实际更新的应该是翻译表中
	nameKey := aiProject.Name
	descKey := aiProject.Description

	// todo 这里应该学习如何用事物处理。
	if !isUpdate || (isUpdate && !helpers.Empty(request.Name)) {
		nameTranslationModel := translation.GetByTidLang(nameKey, request.Language)
		if nameTranslationModel.ID == 0 {
			// 没有对应语言的翻译 就创建
			nameTranslationModel = translation.Translation{
				TranslationId:  nameKey,
				Language:       request.Language,
				TranslatedText: request.Name,
			}
			if err := tx.Create(&nameTranslationModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			nameTranslationModel.TranslatedText = request.Name
			if err := tx.Save(&nameTranslationModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if !isUpdate || (isUpdate && !helpers.Empty(request.Description)) {
		descTranslationModel := translation.GetByTidLang(descKey, request.Language)
		if descTranslationModel.ID == 0 {
			descTranslationModel = translation.Translation{
				TranslationId:  descKey,
				Language:       request.Language,
				TranslatedText: request.Description,
			}
			if err := tx.Create(&descTranslationModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			descTranslationModel.TranslatedText = request.Description
			if err := tx.Save(&descTranslationModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
