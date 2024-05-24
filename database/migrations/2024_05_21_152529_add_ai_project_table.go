package migrations

import (
	"database/sql"
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type AiProject struct {
		models.BaseModel
		models.BaseUUIDModel

		Name        string `gorm:"type:varchar(191);not null;index;"`
		Description string `gorm:"type:varchar(191);default:null;"`
		UserID      string `gorm:"type:varchar(191);not null;index"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AiProject{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AiProject{})
	}

	migrate.Add("2024_05_21_152529_add_ai_project_table", up, down)
}
