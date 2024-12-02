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

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *Cat) error {
	return r.db.Create(cat).Error
}

func (r *catRepository) FindByID(id uint) (*Cat, error) {
	var cat Cat
	err := r.db.First(&cat, id).Error
	return &cat, err
}

func (r *catRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	err := r.db.Find(&cats).Error
	return cats, err
}

func (r *catRepository) Update(cat *Cat) error {
	return r.db.Save(cat).Error
}

func (r *catRepository) Delete(id uint) error {
	return r.db.Delete(&Cat{}, id).Error
}
