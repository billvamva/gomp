package counter

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/billvamva/gomp/internal/components/button"
	"github.com/billvamva/gomp/internal/components/card"
	"github.com/billvamva/gomp/internal/components/form"
	"github.com/billvamva/gomp/internal/components/header"
	"github.com/billvamva/gomp/internal/components/text"
)

type Counter struct {
	Count           int
	CountText       template.HTML
	Header          template.HTML
	IncrementButton template.HTML
	DecrementButton template.HTML
	Cards           []template.HTML
	FormComponent   *form.Form
}

const counterTemplate = `
	<div id="counter" hx-target="this" hx-swap="outerHTML">
	{{.Header}}
	{{.CountText}}<h2>{{.Count}}</h2>
	{{.IncrementButton}}
	{{.DecrementButton}}
	<div class="cards-container">
		{{range .Cards}}
			{{.}}
		{{end}}
	</div>
    <div id="form-container" hx-target="this" hx-swap="innerHTML">
	{{ .Form }} 
	</div>
</div>
`

func NewCounter(initCount int, countText *text.Text, header *header.Header, incrementButton *button.Button, decrementButton *button.Button, nameForm *form.Form) *Counter {
	countTextBuf := new(bytes.Buffer)
	countText.Render(countTextBuf)
	headerBuf := new(bytes.Buffer)
	header.Render(headerBuf)
	incBuf := new(bytes.Buffer)
	incrementButton.Render(incBuf)
	decBuf := new(bytes.Buffer)
	decrementButton.Render(decBuf)

	counter := &Counter{
		Count:           initCount,
		CountText:       template.HTML(countTextBuf.String()),
		Header:          template.HTML(headerBuf.String()),
		IncrementButton: template.HTML(incBuf.String()),
		DecrementButton: template.HTML(decBuf.String()),
		FormComponent:   nameForm,
	}

	counter.updateCards()

	return counter
}

func (c *Counter) updateCards() {
	c.Cards = make([]template.HTML, c.Count)
	for i := 0; i < c.Count; i++ {
		cardComponent := card.NewProjectCard(
			fmt.Sprintf("Card %d", i+1),
			"This is a simple card",
			"simple-card",
		)
		cardBuf := new(bytes.Buffer)
		cardComponent.Render(cardBuf)
		c.Cards[i] = template.HTML(cardBuf.String())
	}
}

func (c *Counter) Render(buf *bytes.Buffer) error {
	tmpl := template.Must(template.New("counter").Parse(counterTemplate))
	formBuf := new(bytes.Buffer)
	c.FormComponent.Render(formBuf)

	data := struct {
		*Counter
		Form template.HTML
	}{
		Counter: c,
		Form:    template.HTML(formBuf.String()),
	}
	return tmpl.Execute(buf, data)
}

func (c *Counter) Increment() {
	c.Count++
	c.updateCards()
}

func (c *Counter) Decrement() {
	if c.Count > 0 {
		c.Count--
		c.updateCards()
	}
}

func (c *Counter) UpdateForm(newForm *form.Form) {
	c.FormComponent = newForm
}

func (c *Counter) RenderFormOnly(buf *bytes.Buffer) error {
	return c.FormComponent.Render(buf)
}
