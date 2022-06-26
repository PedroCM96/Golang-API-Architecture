package migrations

type Migrations struct {
	orderedMigrations map[uint]Migration
}

func NewMigrations() *Migrations {
	m := make(map[uint]Migration)
	m[0] = AddDummyTable{}
	// ... Register migrations

	return &Migrations{orderedMigrations: m}
}
