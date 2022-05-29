package system

import (
	"Duna/database/migrations"
)

type Migrate struct {
	Migrator *migrations.Migrator
}

func (c Migrate) New(migrator *migrations.Migrator) *Migrate {
	return &Migrate{migrator}
}

func (c Migrate) Exec(args []string) {
	c.Migrator.Migrate()
}
