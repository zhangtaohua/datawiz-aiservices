// Package download 模型
package download

import (
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/database"
)

type Download struct {
	models.BaseModel

	// Put fields in here
	Name      string `json:"name"`
	AssetPath string `json:"asset_path"`
	Md5       string `json:"md5"`
	Status    string `json:"status"`

	models.CommonTimestampsField
}

func (download *Download) Create() {
	database.DB.Create(&download)
}

func (download *Download) Save() (rowsAffected int64) {
	result := database.DB.Save(&download)
	return result.RowsAffected
}

func (download *Download) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&download)
	return result.RowsAffected
}
