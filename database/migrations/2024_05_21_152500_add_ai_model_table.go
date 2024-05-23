package migrations

import (
	"database/sql"
	"datawiz-aiservices/app/models"
	"datawiz-aiservices/pkg/migrate"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func init() {

	type AiModel struct {
		models.BaseModel
		models.BaseUUIDModel

		Name        string `gorm:"type:varchar(255);not null;index;"`
		Description string `gorm:"type:varchar(255);default:null;"`
		Type        string `gorm:"type:varchar(255);default:null;"`

		Framework    string `gorm:"type:varchar(255);default:null;"`
		Algorithm    string `gorm:"type:varchar(255);default:null;"`
		NetStructure string `gorm:"type:varchar(255);default:null;"`
		BackBone     string `gorm:"type:varchar(255);default:null;"`

		Accuracy  float32 `gorm:"type:float;comment:准确率;"`
		Precision float32 `gorm:"type:float;comment:精确率;"`
		Recall    float32 `gorm:"type:float;comment:召回率;"`
		F1Score   float32 `gorm:"type:float;comment:F1得分;"`
		AUC       float32 `gorm:"type:float;comment:AUC值;"`

		InputFeatures datatypes.JSONMap `gorm:"comment:输入参数;"`
		OutputLabels  datatypes.JSONMap `gorm:"comment:输出标签;"`
		ExecMethod    datatypes.JSONMap `gorm:"comment:执行方法;"`

		Size       float32   `gorm:"comment:模型大小;"`
		Version    string    `gorm:"type:varchar(32);default:null;comment:版本号;"`
		Status     string    `gorm:"type:varchar(32);default:null;comment:状态;"`
		DeployedAt time.Time `gorm:"column:deployed_at;comment:部署时间;"`
		RetiredAt  time.Time `gorm:"column:retired_at;comment:废弃时间;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AiModel{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AiModel{})
	}

	migrate.Add("2024_05_21_152500_add_ai_model_table", up, down)
}
