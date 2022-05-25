package main

import (
	"AcademyManagement/database/config"
	"AcademyManagement/database/connection"
	http "AcademyManagement/http/router"
	"fmt"
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
	connection.Connect(mySqlConnection)
	fmt.Println(databaseConfig)

	// Initialize router
	router := http.Router()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
