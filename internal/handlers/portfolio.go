package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/billvamva/gomp/database"
	"github.com/billvamva/gomp/internal/components/card"
	"github.com/billvamva/gomp/internal/components/header"
	"github.com/billvamva/gomp/internal/components/text"
	"github.com/gin-gonic/gin"
)

func HandleMain(c *gin.Context) {
	var headerHTML, aboutHTML, projectsHTML string
	var err error

	// Create header
	headerHTML, err = renderHeader()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "Error rendering header"})
		return
	}

	// Create about section
	aboutHTML, err = renderAboutSection()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "Error rendering about section"})
		return
	}

	// Create projects section
	projectsHTML, err = renderProjectsSection()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "Error rendering projects section"})
		return
	}

	content := template.HTML(`
        <section id="about">
            <h2>About Me</h2>
            ` + aboutHTML + `
        </section>
        <section id="projects">
            <h2>My Projects</h2>
            <div class="cards-container">
                ` + projectsHTML + `
            </div>
        </section>
    `)

	c.HTML(http.StatusOK, "portfolio.tmpl", gin.H{
		"Header":  template.HTML(headerHTML),
		"Content": content,
	})
}

func renderHeader() (string, error) {
	h := header.NewHeader("Vasileios Vamvakas", []struct{ Text, Url string }{
		{"About Me", "#about"},
		{"Projects", "#projects"},
	})
	var headerBuf bytes.Buffer
	err := h.Render(&headerBuf)
	if err != nil {
		return "", err
	}
	return headerBuf.String(), nil
}

func renderAboutSection() (string, error) {
	about := `Hi, I'm Vasilis! I'm originally from Greece and now based in the UK and I have been a
	professional software engineer for about 2 years now. I have a background in Electronic Engineering, which 
	gave me first exposure in ML/AI Applications and got be hooked in programming. Now I'm pursuing a career in
	backend and platform engineering. I am planning to use this website as kind of a portfolio and in the future maybe a dev log of some sorts.
	Connect with me on my socials :)`
	aboutText := text.NewText(about).WithTag("p")
	var aboutBuf bytes.Buffer
	err := aboutText.Render(&aboutBuf)
	if err != nil {
		return "", err
	}
	return aboutBuf.String(), nil
}

func renderProjectsSection() (string, error) {
	projects, err := database.GetProjects()
	if err != nil {
		return "", err
	}

	var projectsHTML string
	for _, project := range projects {
		projectCard := card.NewProjectCard(project.Name, project.Description, "card")
		for key, value := range project.Tags {
			tagText := text.NewText(key + ": " + value).WithTag("span")
			projectCard.AddComponent(tagText)
		}
		urlText := text.NewText("Visit Project").WithTag("a").WithAttribute("href", project.Url).WithAttribute("target", "_blank")
		projectCard.AddComponent(urlText)

		var cardBuf bytes.Buffer
		err = projectCard.Render(&cardBuf)
		if err != nil {
			return "", err
		}
		projectsHTML += cardBuf.String()
	}

	return projectsHTML, nil
}
