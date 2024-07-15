// Package ai_project_result 模型
package ai_project_result

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/app/models/ai_model"
	"datawiz-aiservices/app/models/ai_project"
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/helpers"

	"gorm.io/datatypes"
)

type AiProjectResult struct {
	models.BaseModel
	models.BaseUUIDModel

	// 基本参数
	Name        string `json:"name"`
	Description string `json:"description"`

	// 输入参数
	Input datatypes.JSONMap `json:"input"`

	// 输出结果
	Output datatypes.JSONMap `json:"output"`

	Progress uint32 `json:"progress"`
	Status   string `json:"status"`
	Views    uint32 `json:"views"`

	// 其他信息
	UserID        string `json:"user_id"`
	AiModelUUID   string `json:"ai_model_uuid"`
	AiProjectUUID string `json:"ai_project_uuid"`

	AiModel   ai_model.AiModel     `json:"ai_model" gorm:"foreignKey:AiModelUUID;references:UUID"`
	AiProject ai_project.AiProject `json:"ai_project" gorm:"foreignKey:AiProjectUUID;references:UUID"`

	models.CommonTimestampsField
}

func (aiProjectResult *AiProjectResult) Create() {
	database.DB.Create(&aiProjectResult)
}

func (aiProjectResult *AiProjectResult) Save() (rowsAffected int64) {
	result := database.DB.Save(&aiProjectResult)
	return result.RowsAffected
}

func (aiProjectResult *AiProjectResult) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&aiProjectResult)
	return result.RowsAffected
}

func (aiProjectResult *AiProjectResult) PatchNoTx(field string, value interface{}) (rowsAffected int64) {
	result := database.DB.Model(&aiProjectResult).Omit("AiModel").Omit("AiProject").Omit("Associations").Select(field).Update(field, value)
	return result.RowsAffected
}

func (aiProjectResult *AiProjectResult) CreateTx(request *requests.AiProjectResultRequest) error {
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

	aiProjectResult.Name = nameUUID
	aiProjectResult.Description = descUUID
	if err := tx.Create(&aiProjectResult).Error; err != nil {
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

func (aiProjectResult *AiProjectResult) SaveTx(request *requests.AiProjectResultRequest, isUpdate bool) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 实际更新的应该是翻译表中
	nameKey := aiProjectResult.Name
	descKey := aiProjectResult.Description

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
