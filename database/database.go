package database

import (
	"log"

	"vet-clinic-api/database/dbmodel"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Cat{},
		&dbmodel.Visit{},
		&dbmodel.Treatment{},
	)
	log.Println("Database migrated successfully")
}
