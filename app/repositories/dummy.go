package repositories

import "Duna/app/models"

type DummyRepository interface {
	GetById(id uint) (*models.Dummy, error)
	Create(dummy *models.Dummy) error
}
