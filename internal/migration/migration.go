package migration

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return autoMigrateV1(db)
}
