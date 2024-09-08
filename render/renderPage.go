package render

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/billvamva/gomp/internal/components"
	"github.com/gin-gonic/gin"
)

type PageData struct {
	Title   string
	Content template.HTML
}

func RenderPage(c *gin.Context, title string, component components.Component) {
	var buf bytes.Buffer
	if err := component.Render(&buf); err != nil {
		c.AbortWithError(500, err)
		return
	}

	data := PageData{
		Title:   title,
		Content: template.HTML(buf.String()),
	}

	c.HTML(http.StatusOK, "layout.tmpl", data)
}
