package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AiProjectRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description"`
}

func AiProjectSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:255", "not_exists:ai_projects,name"},
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
