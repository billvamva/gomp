package header

import (
	"bytes"
	"html/template"
)

type Header struct {
	Title string
	Links []struct{ Text, Url string }
}

const headerTemplate = `
<header>
	<h1>{{.Title}}</h1>
	<nav>
		{{range .Links}}
		<a href="{{.Url}}">{{.Text}}</a>
		{{end}}
	</nav>
</header>
`

func NewHeader(title string, links []struct{ Text, Url string }) *Header {
	return &Header{Title: title, Links: links}
}

func (c *Header) Render(buf *bytes.Buffer) error {
	tmpl := template.Must(template.New("header").Parse(headerTemplate))
	return tmpl.Execute(buf, c)
}
