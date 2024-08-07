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

		Name        string `gorm:"type:varchar(191);not null;index;"`
		Description string `gorm:"type:varchar(191);default:null;"`
		Code        string `gorm:"type:varchar(191);not null;uniqueIndex;"`
		Type        string `gorm:"type:varchar(191);default:null;comment:模型类型（如分类、回归、聚类等）;"`
		Category    string `gorm:"type:varchar(191);default:null;comment:分类（如AI解释，基础处理， 算法工具等）;"`
		Icon        string `gorm:"type:varchar(191);default:null;"`
		Cover       string `gorm:"type:varchar(191);default:null;"`

		Framework    string `gorm:"type:varchar(191);default:null;comment:深度学习框架（如TensorFlow、PyTorch）;"`
		Algorithm    string `gorm:"type:varchar(191);default:null;"`
		NetStructure string `gorm:"default:null;"`
		BackBone     string `gorm:"default:null;"`

		Accuracy  float32 `gorm:"type:float;comment:准确率;"`
		Precision float32 `gorm:"type:float;comment:精确率;"`
		Recall    float32 `gorm:"type:float;comment:召回率;"`
		F1Score   float32 `gorm:"type:float;comment:F1得分;"`
		AUC       float32 `gorm:"type:float;comment:AUC值;"`

		InputFeatures   string            `gorm:"comment:输入特征;"`
		OutputLabels    string            `gorm:"comment:输出标签;"`
		InputParameters datatypes.JSONMap `gorm:"comment:输入参数;"`
		ExecMethod      datatypes.JSONMap `gorm:"comment:执行方法;"`
		OutputFormatter datatypes.JSONMap `gorm:"comment:输出格式;"`

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
