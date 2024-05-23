package seeders

import (
	"datawiz-aiservices/database/factories"
	"datawiz-aiservices/pkg/console"
	"datawiz-aiservices/pkg/logger"
	"datawiz-aiservices/pkg/seed"
	"fmt"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedTranslationsTable", func(db *gorm.DB) {

		translations := factories.MakeTranslations()

		result := db.Table("translations").Create(&translations)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
