package main

import (
	http "AcademyManagement/http/router"
)

func main() {
	router := http.Router()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
