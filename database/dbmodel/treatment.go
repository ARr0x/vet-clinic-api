package dbmodel

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Treatment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	VisitID     uint      `gorm:"not null" json:"visit_id"`
	Description string    `gorm:"not null" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (t *Treatment) Bind(r *http.Request) error {
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindByID(id uint) (*Treatment, error)
	FindAll() ([]Treatment, error)
	Update(treatment *Treatment) error
	Delete(id uint) error
}

type treatmentRepository struct {
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
	err := r.db.First(&treatment, id).Error
	return &treatment, err
}

func (r *treatmentRepository) FindAll() ([]Treatment, error) {
	var treatments []Treatment
	err := r.db.Find(&treatments).Error
	return treatments, err
}

func (r *treatmentRepository) Update(treatment *Treatment) error {
	return r.db.Save(treatment).Error
}

func (r *treatmentRepository) Delete(id uint) error {
	return r.db.Delete(&Treatment{}, id).Error
}
