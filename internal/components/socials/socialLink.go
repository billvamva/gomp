package socials

import (
	"bytes"
	"html/template"
)

type SocialLink struct {
	Name string
	URL  string
}

type SocialMedia struct {
	Links []SocialLink
}

const socialMediaTemplate = `
<div class="social-media">
    {{range .Links}}
    <a href="{{.URL}}" target="_blank" rel="noopener noreferrer" class="social-link">{{.Name}}</a>
    {{end}}
</div>
`

func NewSocialMedia(links []SocialLink) *SocialMedia {
	return &SocialMedia{Links: links}
}

func (s *SocialMedia) Render(buf *bytes.Buffer) error {
	tmpl, err := template.New("socialMedia").Parse(socialMediaTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(buf, s)
}
