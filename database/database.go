package database

import (
	"vet-clinic-api/database/dbmodel"

	"gorm.io/gorm"
)

// Migrate exécute les migrations pour tous les modèles.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&dbmodel.Cat{},
		&dbmodel.Visit{},
		&dbmodel.Treatment{},
	)
}
