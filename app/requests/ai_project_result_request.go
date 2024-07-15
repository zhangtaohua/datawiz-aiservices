package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/datatypes"
)

type AiProjectResultRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description,omitempty"`

	Input  datatypes.JSONMap `json:"input"`
	Output datatypes.JSONMap `json:"output"`

	Progress uint32 `json:"progress"`
	Status   string `json:"status"`
	Views    uint32 `json:"views"`

	UserID        string `json:"user_id"`
	AiModelUUID   string `json:"ai_model_uuid"`
	AiProjectUUID string `json:"ai_project_uuid"`

	Language string `valid:"language" json:"language"`
}

func AiProjectResultSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:191", "not_exists:ai_project_results,name"},
		"description": []string{"min_cn:3", "max_cn:191"},
		"language":    []string{"required", "min:2", "max:32", "in:en,zh-CN,zh-TW"},
	}
	messages := govalidator.MapData{
		"name": []string{
			RequiredMsg("c.name"),
			MinCnMsg("c.name", "2"),
			MaxCnMsg("c.name", "191"),
			NotExistMsg("c.name"),
		},
		"description": []string{
			MinCnMsg("c.description", "2"),
			MaxCnMsg("c.description", "191"),
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
