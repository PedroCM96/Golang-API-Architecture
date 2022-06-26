package migrations

import (
	"Duna/app/models"
	"gorm.io/gorm"
)

type AddDummyTable struct{}

func (m AddDummyTable) Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Dummy{})
}
