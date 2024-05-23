// Package ai_model 模型
package ai_model

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/database"
	"time"

	"gorm.io/datatypes"
)

type AiModel struct {
	models.BaseModel
	models.BaseUUIDModel

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`

	Framework    string `json:"framework"`
	Algorithm    string `json:"algorithm"`
	NetStructure string `json:"net_structrue"`
	BackBone     string `json:"back_bone"`

	Accuracy  float32 `json:"accuracy"`
	Precision float32 `json:"precision"`
	Recall    float32 `json:"recall"`
	F1Score   float32 `json:"f1_score"`
	AUC       float32 `json:"auc"`

	InputFeatures datatypes.JSONMap `json:"input_features"`
	OutputLabels  datatypes.JSONMap `json:"out_labels"`
	ExecMethod    datatypes.JSONMap `json:"exec_method"`

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
