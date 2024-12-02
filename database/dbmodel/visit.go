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

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Create(treatment *Treatment) error {
	return r.db.Create(treatment).Error
}

func (r *treatmentRepository) FindByID(id uint) (*Treatment, error) {
	var treatment Treatment
	err := r.db.Preload("Visit").First(&treatment, id).Error
	return &treatment, err
}

func (r *treatmentRepository) FindByVisitID(visitID uint) ([]Treatment, error) {
	var treatments []Treatment
	err := r.db.Where("visit_id = ?", visitID).Find(&treatments).Error
	return treatments, err
}

func (r *treatmentRepository) Update(treatment *Treatment) error {
	return r.db.Save(treatment).Error
}

func (r *treatmentRepository) Delete(id uint) error {
	return r.db.Delete(&Treatment{}, id).Error
}
