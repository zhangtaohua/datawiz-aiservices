package migrations

import (
	"database/sql"
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Download struct {
		models.BaseModel

		Name      string `gorm:"type:varchar(191);not null;index"`
		AssetPath string `gorm:"type:varchar(191);default:null"`
		Md5       string `gorm:"type:varchar(191);index;default:null"`
		Status    string `gorm:"type:varchar(191);default:-1;comment:状态;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Download{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Download{})
	}

	migrate.Add("2024_06_06_092025_add_downloads_table", up, down)
}
