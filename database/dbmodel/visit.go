package dbmodel

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	CatID        uint        `gorm:"not null" json:"cat_id"`
	Date         time.Time   `json:"date"`
	Reason       string      `gorm:"not null" json:"reason"`
	Veterinarian string      `gorm:"not null" json:"veterinarian"`
	Treatments   []Treatment `gorm:"foreignKey:VisitID" json:"treatments,omitempty"`
}

func (v *Visit) Bind(r *http.Request) error {
	if v.Reason == "" {
		return errors.New("reason is required")
	}
	if v.Veterinarian == "" {
		return errors.New("veterinarian is required")
	}
	return nil
}

type VisitRepository interface {
	Create(visit *Visit) error
	FindByID(id uint) (*Visit, error)
	FindAll() ([]Visit, error)
	Update(visit *Visit) error
	Delete(id uint) error
}

type visitRepository struct {
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
	err := r.db.Preload("Treatments").First(&visit, id).Error
	return &visit, err
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	err := r.db.Preload("Treatments").Find(&visits).Error
	return visits, err
}

func (r *visitRepository) Update(visit *Visit) error {
	return r.db.Save(visit).Error
}

func (r *visitRepository) Delete(id uint) error {
	return r.db.Delete(&Visit{}, id).Error
}
