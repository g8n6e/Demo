package routes

import (
	"net/http"

	"prizepicks/jurassicpark/handlers"

	"github.com/gin-gonic/gin"
)

func addCageRoutes(rg *gin.RouterGroup) {
	cage := rg.Group("/cage")

	cage.POST("/", func(c *gin.Context) {
		cage, err := handlers.AddCage(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cage)
	})
	cage.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "cage by id")
	})
	cage.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "update cage")
	})
}

func addCagesRoutes(rg *gin.RouterGroup) {
	cages := rg.Group("/cages")

	cages.GET("/", func(c *gin.Context) {
		cages := handlers.GetCages()
		c.JSON(http.StatusOK, cages)
	})
}
