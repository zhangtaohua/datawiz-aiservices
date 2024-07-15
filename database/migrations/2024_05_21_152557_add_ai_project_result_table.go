package migrations

import (
	"database/sql"
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/migrate"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func init() {

	type AiModel struct {
		models.BaseModel
		models.BaseUUIDModel
	}

	type AiProject struct {
		models.BaseModel
		models.BaseUUIDModel
	}

	type AiProjectResult struct {
		models.BaseModel
		models.BaseUUIDModel

		Name        string `gorm:"type:varchar(191);not null;index;"`
		Description string `gorm:"type:varchar(191);default:null;"`

		Input datatypes.JSONMap `gorm:"comment:输入参数;"`

		Output datatypes.JSONMap `gorm:"comment:输出结果;"`

		Progress uint32 `gorm:"default:0;comment:进度条;"`
		Status   string `gorm:"type:varchar(32);default:null;comment:状态;"`
		Views    uint32 `gorm:"default:0;comment:浏览次数;"`

		UserID        string `gorm:"type:varchar(191);not null;index;"`
		AiModelUUID   string `gorm:"type:varchar(191);not null;index;"`
		AiProjectUUID string `gorm:"type:varchar(191);not null;index;"`

		AiModel   AiModel   `gorm:"foreignKey:AiModelUUID;references:UUID;"`
		AiProject AiProject `gorm:"foreignKey:AiProjectUUID;references:UUID;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AiProjectResult{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AiProjectResult{})
	}

	migrate.Add("2024_05_21_152557_add_ai_project_result_table", up, down)
}
