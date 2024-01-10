package routes

import (
	"net/http"
	"prizepicks/jurassicpark/handlers"

	"github.com/gin-gonic/gin"
)

func addDinosaurRoutes(rg *gin.RouterGroup) {
	dinosaur := rg.Group("/dinosaur")

	dinosaur.POST("/", func(c *gin.Context) {
		dinosaur, err := handlers.AddDinosaur(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dinosaur)
	})
	dinosaur.GET("/:id", func(c *gin.Context) {
		dinosaur, err := handlers.GetDinosaurById(c)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dinosaur)
	})
	dinosaur.PATCH("/:id", func(c *gin.Context) {
		dinosaur, err := handlers.UpdateDinosaur(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dinosaur)
	})
}

func addDinosaursRoutes(rg *gin.RouterGroup) {
	dinosaurs := rg.Group("/dinosaurs")

	dinosaurs.GET("/", func(c *gin.Context) {
		dinosaurs := handlers.GetDinosaurs()
		c.JSON(http.StatusOK, dinosaurs)
	})
}
