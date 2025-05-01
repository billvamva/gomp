package card

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/billvamva/gomp/internal/components"
)

type Card struct {
	Title       string
	Description string
	Class       string
	Image       string
	Components  []components.Component
}

const cardTemplate = `
<div class={{.Class}}>
	<h3>{{.Title}}</h3>
	<p>{{.Description}}</p>
	{{.ComponentsHTML}}
</div>
`

const cardWithImageTemplate = `
<div class={{.Class}}>
	<h3>{{.Title}}</h3>
	<img src={{.Image}}>
	<p>{{.Description}}</p>
	{{.ComponentsHTML}}
</div>
`

func NewCard(title, description, class, img string) *Card {
	return &Card{
		Title:       title,
		Description: description,
		Class:       class,
		Image:       img,
	}
}

func (c *Card) AddComponent(component components.Component) {
	c.Components = append(c.Components, component)
}

func (c *Card) Render(buf *bytes.Buffer) error {
	componentsBuffer := new(bytes.Buffer)
	for _, component := range c.Components {
		if err := component.Render(componentsBuffer); err != nil {
			return fmt.Errorf("error rendering component: %w", err)
		}
	}
	data := struct {
		*Card
		ComponentsHTML template.HTML
	}{
		Card:           c,
		ComponentsHTML: template.HTML(componentsBuffer.String()),
	}

	var tmpl *template.Template
	if c.Image == "" {
		tmpl = template.Must(template.New("card").Parse(cardTemplate))
	} else {
		tmpl = template.Must(template.New("card").Parse(cardWithImageTemplate))
	}
	return tmpl.Execute(buf, data)
}
