package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandlePost(c *gin.Context) {
	title := "post" + c.Param("id")
	date := time.Now().Format(time.DateOnly)
	content := "this is some temp content"
	imageURL := "/static/firstpost.jpg"
	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"Title":    title,
		"Date":     date,
		"ImageURL": imageURL,
		"Content":  content,
	})
}
