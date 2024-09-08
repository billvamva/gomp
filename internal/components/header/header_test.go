package header

import (
	"bytes"
	"testing"

	"github.com/billvamva/gomp/internal/components"
)

func TestNewHeader(t *testing.T) {
	title := "My Website"
	links := []struct{ Text, Url string }{
		{Text: "Home", Url: "/"},
		{Text: "About", Url: "/about"},
	}

	header := NewHeader(title, links)

	if header.Title != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, header.Title)
	}

	if len(header.Links) != len(links) {
		t.Errorf("Expected %d links, got %d", len(links), len(header.Links))
	}

	for i, link := range header.Links {
		if link != links[i] {
			t.Errorf("Expected link %d to be %v, got %v", i, links[i], link)
		}
	}
}

func TestHeaderRender(t *testing.T) {
	testCases := []struct {
		name     string
		header   *Header
		expected string
	}{
		{
			name: "Basic header",
			header: NewHeader("My Website", []struct{ Text, Url string }{
				{Text: "Home", Url: "/"},
				{Text: "About", Url: "/about"},
			}),
			expected: `
<header>
	<h1>My Website</h1>
	<nav>
		<a href="/">Home</a>
		<a href="/about">About</a>
	</nav>
</header>`,
		},
		{
			name:   "Header with no links",
			header: NewHeader("Empty Site", nil),
			expected: `
<header>
	<h1>Empty Site</h1>
	<nav>
	</nav>
</header>`,
		},
		{
			name: "Header with special characters",
			header: NewHeader("Test & Demo", []struct{ Text, Url string }{
				{Text: "Home & Start", Url: "/?start=true"},
			}),
			expected: `
<header>
	<h1>Test &amp; Demo</h1>
	<nav>
		<a href="/?start=true">Home &amp; Start</a>
	</nav>
</header>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tc.header.Render(&buf)
			if err != nil {
				t.Fatalf("Render returned an error: %v", err)
			}

			result := buf.String()
			if components.NormalizeWhitespace(result) != components.NormalizeWhitespace(tc.expected) {
				t.Errorf("Expected rendered header to be:\n%s\n\nGot:\n%s", tc.expected, result)
			}
		})
	}
}
