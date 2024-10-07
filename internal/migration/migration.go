package migration

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, sugar *zap.SugaredLogger) {
	if err := autoMigrateV1(db); err != nil {
		sugar.Fatalf("Error migration v1: %v", err)
	}
}
