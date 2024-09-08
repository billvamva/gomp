package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MountHomeHandler(router *gin.Engine) {
	// Set up routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.tmpl", gin.H{
			"Title":   "Welcome",
			"Content": "Welcome to Gomp",
		})
	})
}
