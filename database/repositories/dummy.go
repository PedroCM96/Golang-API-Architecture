package repositories

import (
	"Duna/database/models"
	"Duna/database/repositories/errors"
	"gorm.io/gorm"
)

type DummyRepository struct {
	Db *gorm.DB
}

func (r DummyRepository) Create(d *models.Dummy) (*models.Dummy, error) {
	result := r.Db.Create(d)
	if result.Error != nil {
		return nil, &errors.CreateDummyError{Err: result.Error}
	}
	return d, nil
}
