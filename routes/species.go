package routes

import (
	"net/http"
	"prizepicks/jurassicpark/handlers"

	"github.com/gin-gonic/gin"
)

func addSpecieRoutes(rg *gin.RouterGroup) {
	specie := rg.Group("/specie")

	specie.POST("/", func(c *gin.Context) {
		specie, err := handlers.AddSpecie(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, specie)
	})
	specie.GET("/:id", func(c *gin.Context) {
		specie, err := handlers.GetSpecieById(c)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, specie)
	})
	specie.PATCH("/:id", func(c *gin.Context) {
		specie, err := handlers.UpdateSpecie(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, specie)
	})
}

func addSpeciesRoutes(rg *gin.RouterGroup) {
	species := rg.Group("/species")

	species.GET("/", func(c *gin.Context) {
		species, err := handlers.GetSpecies()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, species)
	})
}
