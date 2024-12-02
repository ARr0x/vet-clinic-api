package dbmodel

import (
	"gorm.io/gorm"
)

// Modèle pour un traitement
type Treatment struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Dosage  string `gorm:"not null"`
	VisitID uint   `gorm:"not null"` // Relation avec Visit
	Visit   Visit  `gorm:"foreignKey:VisitID"`
}

// Interface du repository pour Treatment
type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindByID(id uint) (*Treatment, error)
	FindByVisitID(visitID uint) ([]Treatment, error)
	Update(treatment *Treatment) error
	Delete(id uint) error
}

// Implémentation du repository
type treatmentRepository struct {
	db *gorm.DB
}
