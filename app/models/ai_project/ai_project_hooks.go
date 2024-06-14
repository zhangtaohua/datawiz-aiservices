package ai_project

import (
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/pkg/app"

	"gorm.io/gorm"
)

// func (aiProject *AiProject) BeforeSave(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) BeforeCreate(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) AfterCreate(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) AfterUpdate(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) AfterSave(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) BeforeDelete(tx *gorm.DB) (err error) {}
// func (aiProject *AiProject) AfterDelete(tx *gorm.DB) (err error) {}

func (aiProject *AiProject) AfterFind(tx *gorm.DB) (err error) {
	aiProject.CreatedAt = aiProject.CreatedAt.UTC()
	aiProject.UpdatedAt = aiProject.UpdatedAt.UTC()

	namekey := aiProject.Name
	desckey := aiProject.Description
	// tranV := translation.GetTs([]string{namekey, desckey}, app.Language)
	tranV := translation.TryGetTsV2([]string{namekey, desckey}, app.Language)

	aiProject.Name = tranV[namekey]
	aiProject.Description = tranV[desckey]

	return nil
}
