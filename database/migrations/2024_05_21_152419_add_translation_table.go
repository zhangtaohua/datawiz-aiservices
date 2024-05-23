package migrations

import (
	"database/sql"
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Translation struct {
		models.BaseModel

		TranslationId  string `gorm:"type:varchar(255);not null;uniqueIndex:T_R;"`
		Language       string `gorm:"type:varchar(32);not null;uniqueIndex:T_R;"`
		TranslatedText string `gorm:"type:varchar(255);not null;"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Translation{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Translation{})
	}

	migrate.Add("2024_05_21_152419_add_translation_table", up, down)
}
