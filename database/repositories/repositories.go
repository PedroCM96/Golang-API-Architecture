package repositories

import (
	"Duna/app/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	Dummy repositories.DummyRepository
}

func GetRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Dummy: DatabaseDummyRepository{db},
		// ... Add repositories
	}
}
