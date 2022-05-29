package migrations

import "gorm.io/gorm"

type Migration interface {
	Migrate(db *gorm.DB)
}
