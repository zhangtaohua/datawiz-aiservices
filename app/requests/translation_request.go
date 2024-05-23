package requests

import (
	"datawiz-aiservices/app/models/translation"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TranslationRequest struct {
	TranslationId  string `valid:"translation_id" json:"translation_id"`
	Language       string `valid:"language" json:"language"`
	TranslatedText string `valid:"translated_text" json:"translated_text"`
}

func TranslationSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"translation_id":  []string{"required", "min:2", "max:255"},
		"language":        []string{"required", "min:2", "max:32", "in:en,zh-CN,zh-TW"},
		"translated_text": []string{"required", "min:2", "max:255"},
	}

	messages := govalidator.MapData{
		"translation_id": []string{
			RequiredMsg("c.key"),
			MinMsg("c.key", "2"),
			MaxMsg("c.key", "255"),
		},
		"language": []string{
			RequiredMsg("c.language"),
			MinMsg("c.language", "2"),
			MaxMsg("c.language", "32"),
			InMsg("c.language", []string{"en", "zh-CN", "zh-TW"}),
		},
		"translated_text": []string{
			RequiredMsg("c.transText"),
			MinMsg("c.transText", "2"),
			MaxMsg("c.transText", "32"),
		},
	}

	fmt.Printf("fuckmessages  %v", messages)

	err := validate(data, rules, messages)
	_data := data.(*TranslationRequest)

	count := translation.IsExistUnion(_data.TranslationId, _data.Language)

	if count {
		msg := NotExistUnionMsg([]string{"c.key", "c.language"})
		err["translation_id"] = append(err["translation_id"], msg)
		err["language"] = append(err["language"], msg)
	}

	return err
}
