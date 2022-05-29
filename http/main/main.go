package main

import (
	"Duna/app"
	"Duna/database/config"
	"Duna/database/connection"
	"Duna/database/repositories"
	http "Duna/http/router"
	"github.com/joho/godotenv"
	"log"
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

	repos := repositories.GetRepositories(db)
	router := http.NewRouter(repos)

	a := app.NewApp(repos, router)
	e := a.Start()

	if e != nil {
		panic(e)
	}
}
