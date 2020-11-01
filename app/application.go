package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication to start the server
func StartApplication() {
	mapUrls()
	router.Run(":8080")

	//Add cors to middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		// AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "X-Server-Date",
		// "X-Token"},
		ExposeHeaders:    []string{"X-Server-Date", "X-Token"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))
}
