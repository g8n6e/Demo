package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCageRoutes(rg *gin.RouterGroup) {
	dino := rg.Group("/cage")

	dino.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "add cage")
	})
	dino.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "cage by id")
	})
	dino.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "update cage")
	})
}

func addCagesRoutes(rg *gin.RouterGroup) {
	dinos := rg.Group("/cages")

	dinos.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "all cages")
	})
}
