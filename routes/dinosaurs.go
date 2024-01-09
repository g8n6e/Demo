package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addDinosaurRoutes(rg *gin.RouterGroup) {
	dino := rg.Group("/dinosaur")

	dino.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "add dinosaur")
	})
	dino.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "dino by id")
	})
	dino.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "update dino")
	})
}

func addDinosaursRoutes(rg *gin.RouterGroup) {
	dinos := rg.Group("/dinosaurs")

	dinos.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "all dinosaurs")
	})
}
