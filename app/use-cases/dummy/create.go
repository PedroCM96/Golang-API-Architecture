package dummy

import (
	"Duna/database/models"
	"Duna/database/repositories"
)

type CreateDummyParams struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreateDummy struct {
	DummyRepository *repositories.DummyRepository
}

func (c CreateDummy) Exec(p CreateDummyParams) error {
	d, vErr := models.NewDummy(p.Name, p.Email)

	if vErr != nil {
		return vErr
	}

	_, err := c.DummyRepository.Create(d)
	return err
}
