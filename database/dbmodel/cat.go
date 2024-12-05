package dbmodel

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Cat struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Age       int       `gorm:"not null" json:"age"`
	Breed     string    `gorm:"not null" json:"breed"`
	Weight    float64   `json:"weight"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Visits    []Visit   `gorm:"foreignKey:CatID" json:"visits,omitempty"`
}

func (c *Cat) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Age < 0 {
		return errors.New("age must be non-negative")
	}
	if c.Breed == "" {
		return errors.New("breed is required")
	}
	if c.Weight < 0 {
		return errors.New("weight must be non-negative")
	}
	return nil
}

type CatRepository interface {
	Create(cat *Cat) error
	FindByID(id uint) (*Cat, error)
	FindAll() ([]Cat, error)
	Update(cat *Cat) error
	Delete(id uint) error
}

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
	err := r.db.Preload("Visits").First(&cat, id).Error
	return &cat, err
}

func (r *catRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	err := r.db.Preload("Visits").Find(&cats).Error
	return cats, err
}

func (r *catRepository) Update(cat *Cat) error {
	return r.db.Save(cat).Error
}

func (r *catRepository) Delete(id uint) error {
	return r.db.Delete(&Cat{}, id).Error
}
