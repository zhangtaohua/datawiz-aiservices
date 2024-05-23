// Package ai_project_result 模型
package ai_project_result

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/app/models/ai_model"
	"datawiz-aiservices/app/models/ai_project"
	"datawiz-aiservices/pkg/database"

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

	Status string `json:"status"`

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
