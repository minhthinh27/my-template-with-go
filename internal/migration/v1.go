package migration

import (
	"gorm.io/gorm"
	"my-template-with-go/internal/entity"
)

func autoMigrateV1(db *gorm.DB) error {
	return migrateDB(db)
}

func migrateDB(db *gorm.DB) error {
	db = db.Session(&gorm.Session{SkipHooks: false})
	return db.Migrator().AutoMigrate(
		&entity.Article{},
		&entity.User{},
	)

}
