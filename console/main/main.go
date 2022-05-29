package main

import (
	"Duna/console"
	"Duna/database/config"
	"Duna/database/connection"
	"Duna/database/migrations"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connecting to database
	databaseConfig := config.GetConfig()
	mySqlConnection := connection.MySqlConnection{
		Config: databaseConfig,
	}

	db := connection.Connect(mySqlConnection)
	m := migrations.NewMigrations()
	migrator := migrations.NewMigrator(db, m)

	kernel := console.NewKernel(&console.KernelDeps{Migrator: migrator})
	router := console.NewRouter(kernel)
	name := os.Args[1]
	args := os.Args[2:]

	router.Exec(name, args)

}
