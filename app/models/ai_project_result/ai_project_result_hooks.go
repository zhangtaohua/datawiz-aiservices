package ai_project_result

import "gorm.io/gorm"

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

	return nil
}
