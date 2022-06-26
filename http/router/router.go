package http

import (
	"Duna/database/repositories"
	http "Duna/http/dummy"
	"github.com/gin-gonic/gin"
)

func NewRouter(repositories *repositories.Repositories) *gin.Engine {
	router := gin.Default()
	http.RegisterDummyRoutes(router, repositories)
	// ... Register your groups
	return router
}
