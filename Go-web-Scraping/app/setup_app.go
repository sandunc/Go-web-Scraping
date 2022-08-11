package app

import (
	"Go-web-Scraping/routers"

	"github.com/gin-gonic/gin"

	"log"
)

// Function to setup the app object
func SetupApp() *gin.Engine {

	log.Println("Initializing service")
	// Create barebone engine
	app := gin.New()

	// Setup routers
	log.Println("Setting up routers")
	routers.SetupRouters(app)

	return app
}
