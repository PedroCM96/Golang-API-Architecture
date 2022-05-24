package http

import (
	use_cases "AcademyManagement/app/dummy/use-cases"
	"github.com/gin-gonic/gin"
)

func RegisterDummyRoutes(router *gin.Engine) {
	courses := router.Group("dummy")
	{
		courses.GET("/ping", func(c *gin.Context) {

			pong := use_cases.DummyUseCase()
			c.String(200, pong)
		})
	}
}
