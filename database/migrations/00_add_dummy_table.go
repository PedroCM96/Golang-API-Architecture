package migrations

import (
	"Duna/database/models"
	"gorm.io/gorm"
)

type AddDummyTable struct{}

func (m AddDummyTable) Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Dummy{})
}
