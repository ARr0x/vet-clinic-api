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

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *visitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	err := r.db.Preload("Cat").First(&visit, id).Error
	return &visit, err
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	err := r.db.Preload("Cat").Find(&visits).Error
	return visits, err
}

func (r *visitRepository) FindByCatID(catID uint) ([]Visit, error) {
	var visits []Visit
	err := r.db.Where("cat_id = ?", catID).Find(&visits).Error
	return visits, err
}

func (r *visitRepository) Update(visit *Visit) error {
	return r.db.Save(visit).Error
}

func (r *visitRepository) Delete(id uint) error {
	return r.db.Delete(&Visit{}, id).Error
}
