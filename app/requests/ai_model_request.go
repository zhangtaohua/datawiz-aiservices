package requests

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/datatypes"
)

type AiModelRequest struct {
	Name        string                `valid:"name" json:"name"`
	Description string                `valid:"description" json:"description"`
	Type        string                `valid:"type" json:"type"`
	Category    string                `valid:"category" json:"category"`
	Icon        *multipart.FileHeader `valid:"icon" form:"icon"`

	Framework    string `valid:"framework" json:"framework"`
	Algorithm    string `valid:"algorithm" json:"algorithm"`
	NetStructure string `valid:"net_structrue" json:"net_structrue"`
	BackBone     string `valid:"back_bone" json:"back_bone"`

	Accuracy  float32 `valid:"accuracy" json:"accuracy"`
	Precision float32 `valid:"precision" json:"precision"`
	Recall    float32 `valid:"recall" json:"recall"`
	F1Score   float32 `valid:"f1_score" json:"f1_score"`
	AUC       float32 `valid:"auc" json:"auc"`

	InputFeatures   string            `valid:"input_features" json:"input_features"`
	OutputLabels    string            `valid:"out_labels" json:"out_labels"`
	InputParameters datatypes.JSONMap `valid:"input_parameters" json:"input_parameters"`
	ExecMethod      datatypes.JSONMap `valid:"exec_method" json:"exec_method"`

	Size       float32   `valid:"size" json:"size"`
	Version    string    `valid:"version" json:"version"`
	Status     string    `valid:"status" json:"status"`
	DeployedAt time.Time `valid:"deployed_at" json:"deployed_at"`
	RetiredAt  time.Time `valid:"retired_at" json:"retired_at"`

	Language string `valid:"language" json:"language"`
}

func AiModelSave(data interface{}, c *gin.Context) map[string][]string {

	// todo
	// not_exists 这一条规则也是有问题，现在都要到翻译表中去检查。
	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:255", "not_exists:ai_models,name"},
		"description": []string{"min_cn:2", "max_cn:255"},
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:icon": []string{"ext:png,jpg,jpeg", "size:20971520"},
		"language":  []string{"required", "min:2", "max:32", "in:en,zh-CN,zh-TW"},
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
		"file:icon": []string{
			FileExtMsg(" png, jpg, jpeg"),
			FileSizeMaxMsg(" 20MB"),
		},
		"language": []string{
			RequiredMsg("c.language"),
			MinMsg("c.language", "2"),
			MaxMsg("c.language", "32"),
			InMsg("c.language", []string{"en", "zh-CN", "zh-TW"}),
		},
	}
	return validate(data, rules, messages)
}
