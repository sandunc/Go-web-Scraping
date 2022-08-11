package main

import (
	"Go-web-Scraping/app"
	"Go-web-Scraping/utils"
	"fmt"

	"github.com/gin-gonic/gin"

	"log"
)

func init() {
	// Set gin mode
	mode := utils.GetEnvVar("GIN_MODE")
	gin.SetMode(mode)
}

func main() {
	// Setup the app
	app := app.SetupApp()

	// Read ADDR and port
	addr := utils.GetEnvVar("GIN_ADDR")
	port := utils.GetEnvVar("GIN_PORT")

	if err := app.Run(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

}
