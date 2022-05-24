package http

import (
	http "AcademyManagement/http/dummy"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	http.RegisterDummyRoutes(router)
	return router
}
