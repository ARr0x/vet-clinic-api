package dbmodel

import (
	"gorm.io/gorm"
)

// Modèle pour un chat
type Cat struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Breed     string `gorm:"not null"`
	Weight    float64
	CreatedAt string
}

// Interface du repository pour Cat
type CatRepository interface {
	Create(cat *Cat) error
	FindByID(id uint) (*Cat, error)
	FindAll() ([]Cat, error)
	Update(cat *Cat) error
	Delete(id uint) error
}

// Implémentation du repository
type catRepository struct {
	db *gorm.DB
}
