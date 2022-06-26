package repositories

import (
	"Duna/app/models"
	"gorm.io/gorm"
)

type DatabaseDummyRepository struct {
	db *gorm.DB
}

func (r DatabaseDummyRepository) GetById(id uint) (*models.Dummy, error) {
	dummy := &models.Dummy{ID: id}
	var result *models.Dummy

	r.db.Model(dummy).First(&result)

	return result, nil
}

func (r DatabaseDummyRepository) Create(dummy *models.Dummy) error {
	err := r.db.Create(dummy)
	return err.Error
}
