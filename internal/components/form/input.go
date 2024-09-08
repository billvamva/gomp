package form

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type Input struct {
	Name       string
	Classes    []string
	Attributes map[string]string
	Value      string
}

const inputTemplate = `{{.Tag}}`

func NewInput(name string) *Input {
	return &Input{
		Name:       name,
		Classes:    []string{},
		Attributes: make(map[string]string),
	}
}

func (i *Input) WithClass(class string) *Input {
	i.Classes = append(i.Classes, class)
	return i
}

func (i *Input) UpdateValue(value string) *Input {
	i.Value = value
	return i
}

func (i *Input) WithAttribute(key, value string) *Input {
	i.Attributes[key] = value
	return i
}

func (i *Input) Render(buf *bytes.Buffer) error {
	tmpl, err := template.New("text").Parse(inputTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	tag := "<input"
	tag += fmt.Sprintf(" name=%s", i.Name)
	if len(i.Classes) > 0 {
		tag += fmt.Sprintf(` class="%s"`, strings.Join(i.Classes, " "))
	}
	for k, v := range i.Attributes {
		tag += fmt.Sprintf(` %s="%s"`, template.HTMLEscapeString(k), template.HTMLEscapeString(v))
	}
	if i.Value != "" {
		i.Attributes["value"] = i.Value
	}
	tag += ">"

	data := struct {
		Tag template.HTML
	}{
		Tag: template.HTML(tag),
	}

	return tmpl.Execute(buf, data)
}
