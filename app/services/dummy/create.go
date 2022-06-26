package dummy

import (
	"Duna/app/models"
	"Duna/app/repositories"
)

type CreateDummyParams struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreateDummy struct {
	DummyRepository repositories.DummyRepository
}

func (c CreateDummy) Exec(p CreateDummyParams) error {
	d, _ := models.NewDummy(p.Name, p.Email)

	err := c.DummyRepository.Create(d)
	return err
}
