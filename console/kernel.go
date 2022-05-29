package console

import (
	"Duna/console/commands"
	"Duna/console/commands/system"
	"Duna/database/migrations"
)

type KernelDeps struct {
	Migrator *migrations.Migrator
}

type Kernel struct {
	deps *KernelDeps
}

func NewKernel(deps *KernelDeps) *Kernel {
	return &Kernel{deps}
}

func (k Kernel) commands() map[string]func() Command {
	return map[string]func() Command{
		"system:migrate": func() Command {
			migrate := new(system.Migrate)
			migrate.Migrator = k.deps.Migrator
			return migrate
		},

		"dummy": func() Command {
			return new(commands.DummyCommand)
		},
		// ... Register commands
	}
}
