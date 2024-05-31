// Package ai_model 模型
package ai_model

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/helpers"
	"time"

	"gorm.io/datatypes"
)

type AiModel struct {
	models.BaseModel
	models.BaseUUIDModel

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Icon        string `json:"icon"`

	Framework    string `json:"framework"`
	Algorithm    string `json:"algorithm"`
	NetStructure string `json:"net_structrue"`
	BackBone     string `json:"back_bone"`

	Accuracy  float32 `json:"accuracy"`
	Precision float32 `json:"precision"`
	Recall    float32 `json:"recall"`
	F1Score   float32 `json:"f1_score"`
	AUC       float32 `json:"auc"`

	InputFeatures   string            `json:"input_features"`
	OutputLabels    string            `json:"output_labels"`
	InputParameters datatypes.JSONMap `json:"input_parameters"`
	ExecMethod      datatypes.JSONMap `json:"exec_method"`

	Size       float32   `json:"size"`
	Version    string    `json:"version"`
	Status     string    `json:"status"`
	DeployedAt time.Time `json:"deployed_at"`
	RetiredAt  time.Time `json:"retired_at"`

	models.CommonTimestampsField
}

func (aiModel *AiModel) Create() {
	database.DB.Create(&aiModel)
}

func (aiModel *AiModel) Save() (rowsAffected int64) {
	result := database.DB.Save(&aiModel)
	return result.RowsAffected
}

func (aiModel *AiModel) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&aiModel)
	return result.RowsAffected
}

func (aiModel *AiModel) CreateTx(request *requests.AiModelRequest) error {
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
	inputUUID := helpers.UUID()
	outputUUID := helpers.UUID()

	aiModel.Name = nameUUID
	aiModel.Description = descUUID
	aiModel.InputFeatures = inputUUID
	aiModel.OutputLabels = outputUUID
	if err := tx.Create(&aiModel).Error; err != nil {
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

	inTranslationModel := translation.Translation{
		TranslationId:  inputUUID,
		Language:       request.Language,
		TranslatedText: request.InputFeatures,
	}
	if err := tx.Create(&inTranslationModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	outTranslationModel := translation.Translation{
		TranslationId:  outputUUID,
		Language:       request.Language,
		TranslatedText: request.OutputLabels,
	}
	if err := tx.Create(&outTranslationModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (aiModel *AiModel) SaveTx(request *requests.AiModelRequest) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&aiModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 实际更新的应该是翻译表中
	nameKey := aiModel.Name
	descKey := aiModel.Description
	inputKey := aiModel.InputFeatures
	outputKey := aiModel.OutputLabels

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

	descTranslationModel = translation.GetByTidLang(inputKey, request.Language)
	if descTranslationModel.ID == 0 {
		descTranslationModel = translation.Translation{
			TranslationId:  inputKey,
			Language:       request.Language,
			TranslatedText: request.InputFeatures,
		}
		if err := tx.Create(&descTranslationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		descTranslationModel.TranslatedText = request.InputFeatures
		if err := tx.Save(&descTranslationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	descTranslationModel = translation.GetByTidLang(outputKey, request.Language)
	if descTranslationModel.ID == 0 {
		descTranslationModel = translation.Translation{
			TranslationId:  outputKey,
			Language:       request.Language,
			TranslatedText: request.OutputLabels,
		}
		if err := tx.Create(&descTranslationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		descTranslationModel.TranslatedText = request.OutputLabels
		if err := tx.Save(&descTranslationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
