package requests

import (
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/pkg/helpers"
	"datawiz-aiservices/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	UserID string `valid:"user_id" form:"user_id"`

	Name        string `valid:"name" form:"name"`
	Description string `valid:"description" form:"description"`
	Category    string `valid:"category" form:"category"`
	Status      string `valid:"status" form:"status"`

	StartCreatedAt time.Time `valid:"start_created_at" form:"start_created_at"`
	EndCreatedAt   time.Time `valid:"end_created_at" form:"end_created_at"`

	StartUpdatedAt time.Time `valid:"start_updated_at" form:"start_updated_at"`
	EndUpdatedAt   time.Time `valid:"end_updated_at" form:"end_updated_at"`

	Language string `valid:"language" form:"language"`

	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,created_at,updated_at",
		},
		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
	}
	return validate(data, rules, messages)
}

func MakeWhereFields(request PaginationRequest, isTs bool) []paginator.BaseWhereField {
	whereFields := []paginator.BaseWhereField{}

	if !helpers.Empty(request.UserID) {
		whereField := paginator.BaseWhereField{
			Query:  `user_id = ?`,
			Values: []interface{}{request.UserID},
		}
		whereFields = append(whereFields, whereField)
	}

	if !helpers.Empty(request.Name) {
		if isTs {
			var uuids []string
			translations := translation.GetByTextLang("%"+request.Name+"%", request.Language)
			for _, v := range translations {
				uuids = append(uuids, v.TranslationId)
			}
			whereField := paginator.BaseWhereField{
				Query:  `name IN ?`,
				Values: []interface{}{uuids},
			}
			whereFields = append(whereFields, whereField)
		} else {
			whereField := paginator.BaseWhereField{
				Query:  `name LIKE ?`,
				Values: []interface{}{"%" + request.Name + "%"},
			}
			whereFields = append(whereFields, whereField)
		}

	}

	if !helpers.Empty(request.Description) {
		if isTs {
			var uuids []string
			translations := translation.GetByTextLang("%"+request.Description+"%", request.Language)
			for _, v := range translations {
				uuids = append(uuids, v.TranslationId)
			}
			whereField := paginator.BaseWhereField{
				Query:  `description IN ?`,
				Values: []interface{}{uuids},
			}
			whereFields = append(whereFields, whereField)
		} else {
			whereField := paginator.BaseWhereField{
				Query:  `description LIKE ?`,
				Values: []interface{}{"%" + request.Description + "%"},
			}
			whereFields = append(whereFields, whereField)
		}
	}

	if !helpers.Empty(request.Category) {
		whereField := paginator.BaseWhereField{
			Query:  `category = ?`,
			Values: []interface{}{request.Category},
		}
		whereFields = append(whereFields, whereField)
	}

	if !helpers.Empty(request.Status) {
		whereField := paginator.BaseWhereField{
			Query:  `status = ?`,
			Values: []interface{}{request.Status},
		}
		whereFields = append(whereFields, whereField)
	}

	if !helpers.Empty(request.StartCreatedAt) && !helpers.Empty(request.EndCreatedAt) {
		whereField := paginator.BaseWhereField{
			Query:  `created_at BETWEEN ? AND ?`,
			Values: []interface{}{request.StartCreatedAt.Local(), request.EndCreatedAt.Local()},
		}
		whereFields = append(whereFields, whereField)
	} else if !helpers.Empty(request.StartUpdatedAt) && !helpers.Empty(request.EndUpdatedAt) {
		whereField := paginator.BaseWhereField{
			Query:  `created_at BETWEEN ? AND ?`,
			Values: []interface{}{request.StartUpdatedAt.Local(), request.EndUpdatedAt.Local()},
		}
		whereFields = append(whereFields, whereField)
	}

	return whereFields
}

// modelNames := []string{"Product.Category", "Product.Supplier")
func MakePreloadFields(modelNames []string) []paginator.BasePreloadField {
	preloadFields := []paginator.BasePreloadField{}
	for i := 0; i < len(modelNames); i++ {
		preloadField := paginator.BasePreloadField{
			Query:  modelNames[i],
			Values: []interface{}{},
		}
		preloadFields = append(preloadFields, preloadField)
	}
	return preloadFields
}
