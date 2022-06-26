package http

import (
	"Duna/app/services/dummy"
	"Duna/database/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
				DummyRepository: repositories.Dummy,
			}

			err := uc.Exec(params)

			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		})

		group.GET("/:id", func(c *gin.Context) {
			id, err := strconv.ParseUint(c.Param("id"), 10, 64)

			if err != nil {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("Bad id: %d", id))
				return
			}

			params := dummy.GetDummyParams{Id: uint(id)}
			uc := dummy.GetDummy{DummyRepository: repositories.Dummy}

			d, e := uc.Exec(params)

			if e != nil {
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			c.JSON(200, gin.H{"dummy": d})
		})
	}
}
