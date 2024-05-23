package requests

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/datatypes"
)

type AiModelRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description"`
	Type        string `valid:"type" json:"type"`

	Framework    string `valid:"framework" json:"framework"`
	Algorithm    string `valid:"algorithm" json:"algorithm"`
	NetStructure string `valid:"net_structrue" json:"net_structrue"`
	BackBone     string `valid:"back_bone" json:"back_bone"`

	Accuracy  float32 `valid:"accuracy" json:"accuracy"`
	Precision float32 `valid:"precision" json:"precision"`
	Recall    float32 `valid:"recall" json:"recall"`
	F1Score   float32 `valid:"f1_score" json:"f1_score"`
	AUC       float32 `valid:"auc" json:"auc"`

	InputFeatures datatypes.JSONMap `valid:"input_features" json:"input_features"`
	OutputLabels  datatypes.JSONMap `valid:"out_labels" json:"out_labels"`
	ExecMethod    datatypes.JSONMap `valid:"exec_method" json:"exec_method"`

	Size       float32   `valid:"size" json:"size"`
	Version    string    `valid:"version" json:"version"`
	Status     string    `valid:"status" json:"status"`
	DeployedAt time.Time `valid:"deployed_at" json:"deployed_at"`
	RetiredAt  time.Time `valid:"retired_at" json:"retired_at"`
}

func AiModelSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:255", "not_exists:ai_models,name"},
		"description": []string{"min_cn:2", "max_cn:255"},
	}

	messages := govalidator.MapData{
		"name": []string{
			RequiredMsg("c.name"),
			MinCnMsg("c.name", "2"),
			MaxCnMsg("c.name", "255"),
			NotExistMsg("c.name"),
		},
		"description": []string{
			MinCnMsg("c.description", "2"),
			MaxCnMsg("c.description", "255"),
		},
	}
	return validate(data, rules, messages)
}
