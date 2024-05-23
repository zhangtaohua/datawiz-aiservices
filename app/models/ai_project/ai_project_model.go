// Package ai_project 模型
package ai_project

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/database"
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
