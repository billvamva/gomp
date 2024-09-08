package main

import (
	"github.com/billvamva/gomp/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin router
	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*")

	// Serve static files
	router.Static("/static", "web/static")

	// Mount Handlers
	handlers.MountHomeHandler(router)
	handlers.MountCounterHandler(router)

	// Start the server
	router.Run(":8080")
}
