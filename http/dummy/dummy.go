package http

import (
	"Duna/app/use-cases/dummy"
	"Duna/database/models"
	"Duna/database/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterDummyRoutes(router *gin.Engine, repositories *repositories.Repositories) {
	group := router.Group("dummy")
	{
		group.POST("/", func(c *gin.Context) {
			var params dummy.CreateDummyParams
			parseError := c.BindJSON(&params)

			if parseError != nil {
				c.String(http.StatusBadRequest, "Invalid request")
				return
			}

			uc := dummy.CreateDummy{
				DummyRepository: repositories.DummyRepository,
			}

			err := uc.Exec(params)

			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		})

		group.GET("/", func(c *gin.Context) {
			c.JSON(200, models.Dummy{Email: "abc@gmail.com", Name: "xD"})
		})
	}
}
