package form

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/billvamva/gomp/internal/components"
)

type Form struct {
	Components []components.Component
	Classes    []string
	Attributes map[string]string
}

const formTemplate = `{{.OpenTag}}{{.Content}}{{.CloseTag}}`

func NewForm() *Form {
	return &Form{
		Components: []components.Component{},
		Classes:    []string{},
		Attributes: make(map[string]string),
	}
}

func (f *Form) AddComponent(component components.Component) *Form {
	f.Components = append(f.Components, component)
	return f
}

func (f *Form) ResetComponents() *Form {
	f.Components = []components.Component{}
	return f
}

func (f *Form) WithClass(class string) *Form {
	f.Classes = append(f.Classes, class)
	return f
}

func (f *Form) WithAttribute(key, value string) *Form {
	f.Attributes[key] = value
	return f
}

func (f *Form) Render(buf *bytes.Buffer) error {
	tmpl, err := template.New("text").Parse(formTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	openTag := "<form"
	if len(f.Classes) > 0 {
		openTag += fmt.Sprintf(` class="%s"`, strings.Join(f.Classes, " "))
	}
	for k, v := range f.Attributes {
		openTag += fmt.Sprintf(` %s="%s"`, template.HTMLEscapeString(k), template.HTMLEscapeString(v))
	}
	openTag += ">"

	closeTag := "</form>"

	componentsBuffer := new(bytes.Buffer)
	for _, component := range f.Components {
		if err := component.Render(componentsBuffer); err != nil {
			return fmt.Errorf("error rendering component: %w", err)
		}
	}

	data := struct {
		OpenTag  template.HTML
		Content  template.HTML
		CloseTag template.HTML
	}{
		OpenTag:  template.HTML(openTag),
		Content:  template.HTML(componentsBuffer.String()),
		CloseTag: template.HTML(closeTag),
	}

	return tmpl.Execute(buf, data)
}
