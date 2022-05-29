package migrations

import (
	"gorm.io/gorm"
)

type Migrator struct {
	db         *gorm.DB
	migrations map[uint]Migration
}

func NewMigrator(db *gorm.DB, migrations *Migrations) *Migrator {
	return &Migrator{migrations: migrations.orderedMigrations, db: db}
}

func (m Migrator) Migrate() {
	for _, migration := range m.migrations {
		migration.Migrate(m.db)
	}
}
