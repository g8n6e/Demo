package routes

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"prizepicks/jurassicpark/models"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	getRoutes()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func getRoutes() {
	defaultRG := router.Group("")
	addDinosaurRoutes(defaultRG)
	addDinosaursRoutes(defaultRG)
	addCageRoutes(defaultRG)
	addCagesRoutes(defaultRG)
}

func SetupTestRouter() *gin.Engine {
	models.ConnectDatabase("../demo.db")
	getRoutes()
	return router
}
