package dbmodel

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Modèle pour une consultation

type Visit struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	CatID        uint        `json:"cat_id"`
	Date         time.Time   `json:"date"`
	Reason       string      `json:"reason"`
	Veterinarian string      `json:"veterinarian"`
	Treatments   []Treatment `gorm:"foreignKey:VisitID" json:"treatments,omitempty"`
}

// Bind validates and processes incoming requests for the Visit model.
func (v *Visit) Bind(r *http.Request) error {
	if v.CatID == 0 {
		return errors.New("cat_id is required")
	}
	if v.Date.IsZero() {
		return errors.New("date is required")
	}
	if v.Reason == "" {
		return errors.New("reason is required")
	}
	return nil
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
