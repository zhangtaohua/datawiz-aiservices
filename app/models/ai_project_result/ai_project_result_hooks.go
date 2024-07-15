package ai_project_result

import (
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/helpers"

	"gorm.io/gorm"
)

// func (aiProjectResult *AiProjectResult) BeforeSave(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) BeforeCreate(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) AfterCreate(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) AfterUpdate(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) AfterSave(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) BeforeDelete(tx *gorm.DB) (err error) {}
// func (aiProjectResult *AiProjectResult) AfterDelete(tx *gorm.DB) (err error) {}
func (aiProjectResult *AiProjectResult) AfterFind(tx *gorm.DB) (err error) {
	aiProjectResult.CreatedAt = aiProjectResult.CreatedAt.UTC()
	aiProjectResult.UpdatedAt = aiProjectResult.UpdatedAt.UTC()

	// getTransAiProjectResult()
	namekey := aiProjectResult.Name
	desckey := aiProjectResult.Description

	tranV := translation.TryGetTsV1([]string{namekey, desckey}, app.Language)

	aiProjectResult.Name = tranV[namekey]
	aiProjectResult.Description = tranV[desckey]

	views := aiProjectResult.Views
	if helpers.Empty(views) {
		views = 1
	} else {
		views = views + 1
	}
	aiProjectResult.PatchNoTx("views", views)

	return nil
}
