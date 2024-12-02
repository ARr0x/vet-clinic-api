package database

import (
	"log"

	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&dbmodel.AnimalSound{})
	DB.AutoMigrate(&dbmodel.AgeEntry{})
	log.Println("Database connected and migrated")
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.AgeEntry{},
		&dbmodel.AnimalSound{},
	)
	log.Println("Database migrated successfully")
}
