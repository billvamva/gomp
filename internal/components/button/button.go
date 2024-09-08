package button

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type Button struct {
	Text       string
	Classes    []string
	Attributes map[string]string
}

const buttonTemplate = `{{.OpenTag}}{{.Text}}{{.CloseTag}}`

func NewButton(text string) *Button {
	return &Button{
		Text:       text,
		Classes:    []string{},
		Attributes: make(map[string]string),
	}
}

func (b *Button) WithContent(content string) *Button {
	b.Text = content
	return b
}

func (b *Button) WithClass(class string) *Button {
	b.Classes = append(b.Classes, class)
	return b
}

func (b *Button) WithAttribute(key, value string) *Button {
	b.Attributes[key] = value
	return b
}

func (b *Button) Render(buf *bytes.Buffer) error {
	tmpl := template.Must(template.New("button").Parse(buttonTemplate))
	openTag := "<button"
	if len(b.Classes) > 0 {
		openTag += fmt.Sprintf(` class="%s"`, strings.Join(b.Classes, " "))
	}
	for k, v := range b.Attributes {
		openTag += fmt.Sprintf(` %s="%s"`, template.HTMLEscapeString(k), template.HTMLEscapeString(v))
	}
	openTag += ">"

	closeTag := "</button>"

	data := struct {
		OpenTag  template.HTML
		Text     string
		CloseTag template.HTML
	}{
		OpenTag:  template.HTML(openTag),
		Text:     b.Text,
		CloseTag: template.HTML(closeTag),
	}
	return tmpl.Execute(buf, data)
}
