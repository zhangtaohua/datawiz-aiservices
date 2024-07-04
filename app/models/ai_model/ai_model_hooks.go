package ai_model

import (
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/helpers"

	"gorm.io/gorm"
)

// func (aiModel *AiModel) BeforeSave(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) BeforeCreate(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) AfterCreate(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) AfterUpdate(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) AfterSave(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) BeforeDelete(tx *gorm.DB) (err error) {}
// func (aiModel *AiModel) AfterDelete(tx *gorm.DB) (err error) {}
func (aiModel *AiModel) AfterFind(tx *gorm.DB) (err error) {
	aiModel.CreatedAt = aiModel.CreatedAt.UTC()
	aiModel.UpdatedAt = aiModel.UpdatedAt.UTC()

	if !helpers.Empty(aiModel.DeployedAt) {
		aiModel.DeployedAt = aiModel.CreatedAt.UTC()
	}

	if !helpers.Empty(aiModel.RetiredAt) {
		aiModel.RetiredAt = aiModel.RetiredAt.UTC()
	}

	// getTransAiModel() code
	namekey := aiModel.Name
	desckey := aiModel.Description
	inputkey := aiModel.InputFeatures
	outkey := aiModel.OutputLabels

	tranV := translation.TryGetTsV1([]string{namekey, desckey, inputkey, outkey}, app.Language)

	aiModel.Name = tranV[namekey]
	aiModel.Description = tranV[desckey]
	aiModel.InputFeatures = tranV[inputkey]
	aiModel.OutputLabels = tranV[outkey]
	return nil
}
