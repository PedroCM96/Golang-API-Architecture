package repositories

import "gorm.io/gorm"

type Repositories struct {
	DummyRepository *DummyRepository
}

func GetRepositories(db *gorm.DB) *Repositories {
	repositories := new(Repositories)
	// Add your repositories here
	dummy := &DummyRepository{Db: db}
	repositories.DummyRepository = dummy
	return repositories
}
