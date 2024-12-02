package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "vet-clinic-api/database"
    "vet-clinic-api/database/dbmodel"
)

type Config struct {
    // Connexion aux repositories
    CatRepository     dbmodel.CatRepository
    TreatmentRepository  dbmodel.TreatmentRepository
	VisitRepository  dbmodel.VisitRepository
}

func New() (*Config, error) {
    config := Config{}

    // Initialisation de la connexion à la base de données
    databaseSession, err := gorm.Open(sqlite.Open("vet-clinic-api.db"), &gorm.Config{})
    if err != nil {
        return &config, err
    }

    database.Migrate(databaseSession)

    // Initialisation des repositories
    config.CatRepository = dbmodel.NewCatRepository(databaseSession)
    config.TreatmentRepository = dbmodel.NewTreatmentRepository(databaseSession)
	config.VisitRepository = dbmodel.NewVisitRepository(databaseSession)


    return &config, nil
}