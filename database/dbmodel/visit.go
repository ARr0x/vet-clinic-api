package dbmodel

import (
	"gorm.io/gorm"
)

// Modèle pour une consultation
type Visit struct {
	ID           uint   `gorm:"primaryKey"`
	Date         string `gorm:"not null"`
	Reason       string `gorm:"not null"`
	Veterinarian string `gorm:"not null"`
	CatID        uint   `gorm:"not null"` // Relation avec Cat
	Cat          Cat    `gorm:"foreignKey:CatID"`
}

// Interface du repository pour Visit
type VisitRepository interface {
	Create(visit *Visit) error
	FindByID(id uint) (*Visit, error)
	FindAll() ([]Visit, error)
	FindByCatID(catID uint) ([]Visit, error)
	Update(visit *Visit) error
	Delete(id uint) error
}

// Implémentation du repository
type visitRepository struct {
	db *gorm.DB
}
