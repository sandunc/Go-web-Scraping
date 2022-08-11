package routers

import (
	"Go-web-Scraping/controllers"

	"github.com/gin-gonic/gin"
)

// Function to setup routers and router groups
func SetupRouters(app *gin.Engine) {

	v1 := app.Group("/v1")
	{
		v1.POST("/result", controllers.GetResult)

	}

}
