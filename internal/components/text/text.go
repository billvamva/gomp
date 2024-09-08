package text

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type Text struct {
	Content    string
	Tag        string
	Classes    []string
	Attributes map[string]string
}

const textTemplate = `{{.OpenTag}}{{.Content}}{{.CloseTag}}`

func NewText(content string) *Text {
	return &Text{
		Content:    content,
		Tag:        "p",
		Classes:    []string{},
		Attributes: make(map[string]string),
	}
}

func (t *Text) WithContent(content string) *Text {
	t.Content = content
	return t
}

func (t *Text) WithTag(tag string) *Text {
	t.Tag = tag
	return t
}

func (t *Text) WithClass(class string) *Text {
	t.Classes = append(t.Classes, class)
	return t
}

func (t *Text) WithAttribute(key, value string) *Text {
	t.Attributes[key] = value
	return t
}

func (t *Text) Render(buf *bytes.Buffer) error {
	tmpl, err := template.New("text").Parse(textTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	openTag := fmt.Sprintf("<%s", t.Tag)
	if len(t.Classes) > 0 {
		openTag += fmt.Sprintf(` class="%s"`, strings.Join(t.Classes, " "))
	}
	for k, v := range t.Attributes {
		openTag += fmt.Sprintf(` %s="%s"`, template.HTMLEscapeString(k), template.HTMLEscapeString(v))
	}
	openTag += ">"

	closeTag := fmt.Sprintf("</%s>", t.Tag)

	data := struct {
		OpenTag  template.HTML
		Content  string
		CloseTag template.HTML
	}{
		OpenTag:  template.HTML(openTag),
		Content:  t.Content,
		CloseTag: template.HTML(closeTag),
	}

	return tmpl.Execute(buf, data)
}
