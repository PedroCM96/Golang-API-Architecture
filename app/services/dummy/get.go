package dummy

import (
	"Duna/app/models"
	"Duna/app/repositories"
)

type GetDummyParams struct {
	Id uint `json:"id" binding:"required"`
}

type GetDummy struct {
	DummyRepository repositories.DummyRepository
}

func (g GetDummy) Exec(p GetDummyParams) (models.Dummy, error) {
	d, err := g.DummyRepository.GetById(p.Id)

	return *d, err
}
