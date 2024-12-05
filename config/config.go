package config

import (
	"log"
	"os"

	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Config centralise la configuration de l'application
type Config struct {
	DB                  *gorm.DB
	CatRepository       dbmodel.CatRepository
	VisitRepository     dbmodel.VisitRepository
	TreatmentRepository dbmodel.TreatmentRepository
}

// New initialise la configuration principale
func New() *Config {
	db, err := gorm.Open(sqlite.Open("vet_clinic.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if os.Getenv("MIGRATE") == "true" {
		database.Migrate(db)
	}

	return &Config{
		DB:                  db,
		CatRepository:       dbmodel.NewCatRepository(db),
		VisitRepository:     dbmodel.NewVisitRepository(db),
		TreatmentRepository: dbmodel.NewTreatmentRepository(db),
	}
}
