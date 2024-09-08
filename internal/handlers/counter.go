package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/billvamva/gomp/internal/components/button"
	counter "github.com/billvamva/gomp/internal/components/counter"
	"github.com/billvamva/gomp/internal/components/form"
	"github.com/billvamva/gomp/internal/components/header"
	"github.com/billvamva/gomp/internal/components/text"
	"github.com/billvamva/gomp/render"
	"github.com/gin-gonic/gin"
)

type FormData struct {
	Name  string
	Email string
}

var (
	counterComponent *counter.Counter
	f                *form.Form
	formData         FormData
)

func init() {
	header := header.NewHeader("", []struct{ Text, Url string }{
		{"Posts", "/posts"},
		{"About me", "/about"},
	})

	textComponent := text.NewText("Count").WithTag("h2").WithAttribute("hx-target", "#content").WithAttribute("hx-swap", "beforeend").WithAttribute("style", "color: red; font-weight: bold;")

	incButton := button.NewButton("Increment").
		WithAttribute("hx-post", "increment").WithClass("btn btn-primary")

	decButton := button.NewButton("Decrement").
		WithAttribute("hx-post", "decrement").
		WithClass("btn btn-secondary")

	f = form.NewForm().WithAttribute("id", "name-email").
		WithAttribute("hx-post", "submit").WithAttribute("hx-include", "[name]")

	// Add form components
	f.AddComponent(form.NewInput("name").
		WithAttribute("type", "text").
		UpdateValue(formData.Name).
		WithAttribute("placeholder", "Enter your name"))

	f.AddComponent(form.NewInput("email").
		WithAttribute("type", "email").
		UpdateValue(formData.Email).
		WithAttribute("placeholder", "Enter your email"))

	f.AddComponent(text.NewText("").WithAttribute("hx-target", "#content").WithAttribute("hx-swap", "beforeend").WithAttribute("style", "color: red; font-weight: bold;"))

	f.AddComponent(button.NewButton("submit").WithAttribute("type", "submit"))

	counterComponent = counter.NewCounter(0, textComponent, header, incButton, decButton, f)
}

type CounterHandler struct{}

func MountCounterHandler(router *gin.Engine) {
	counterGroup := router.Group("/counter")
	{
		counterGroup.GET("/", HandleCounterPage)
		counterGroup.POST("increment", HandleIncrement)
		counterGroup.POST("decrement", HandleDecrement)
		counterGroup.POST("submit", HandleSubmit)
	}
}

func HandleCounterPage(c *gin.Context) {
	render.RenderPage(c, "Counter", counterComponent)
}

func HandleIncrement(c *gin.Context) {
	counterComponent.Increment()
	renderCounter(c)
}

func HandleDecrement(c *gin.Context) {
	counterComponent.Decrement()
	renderCounter(c)
}

func HandleSubmit(c *gin.Context) {
	// Log raw body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusNotFound, "Error Reading form body")
		return
	}

	formBody, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		c.String(http.StatusNotFound, "Error Parsing Data")
		return
	}

	name := formBody.Get("name")
	email := formBody.Get("email")

	log.Printf("Parsed form data - Name: %s, Email: %s", name, email)

	renderForm(c)
}

// updating component
func renderCounter(c *gin.Context) {
	var buf bytes.Buffer
	if err := counterComponent.Render(&buf); err != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}
	updateDom(c, &buf)
}

// updating component
func renderForm(c *gin.Context) {
	// remove existing components
	f.ResetComponents()

	// Add form components
	f.AddComponent(form.NewInput("name").
		WithAttribute("type", "text").
		UpdateValue("").
		WithAttribute("placeholder", "Enter your name"))

	f.AddComponent(form.NewInput("email").
		WithAttribute("type", "email").
		UpdateValue("").
		WithAttribute("placeholder", "Enter your email"))

	f.AddComponent(text.NewText("Success").WithAttribute("hx-target", "#content").WithAttribute("hx-swap", "beforeend").WithAttribute("style", "color: red; font-weight: bold;"))

	f.AddComponent(button.NewButton("submit").WithAttribute("type", "submit"))

	// replace form in counter
	counterComponent.UpdateForm(f)

	// render counter form only
	var buf bytes.Buffer
	if err := counterComponent.RenderFormOnly(&buf); err != nil {
		return
	}

	updateDom(c, &buf)
}

func updateDom(c *gin.Context, buf *bytes.Buffer) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, buf.String())
}
