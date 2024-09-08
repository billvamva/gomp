package text

import (
	"bytes"
	"testing"

	"github.com/billvamva/gomp/internal/components"
)

func TestNewText(t *testing.T) {
	text := NewText("Hello, World!")
	if text.Content != "Hello, World!" {
		t.Errorf("Expected content to be 'Hello, World!', got '%s'", text.Content)
	}
	if text.Tag != "p" {
		t.Errorf("Expected default tag to be 'p', got '%s'", text.Tag)
	}
	if len(text.Classes) != 0 {
		t.Errorf("Expected no classes by default, got %d", len(text.Classes))
	}
	if len(text.Attributes) != 0 {
		t.Errorf("Expected no attributes by default, got %d", len(text.Attributes))
	}
}

func TestWithTag(t *testing.T) {
	text := NewText("Test").WithTag("h1")
	if text.Tag != "h1" {
		t.Errorf("Expected tag to be 'h1', got '%s'", text.Tag)
	}
}

func TestWithClass(t *testing.T) {
	text := NewText("Test").WithClass("bold").WithClass("italic")
	if len(text.Classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(text.Classes))
	}
	if text.Classes[0] != "bold" || text.Classes[1] != "italic" {
		t.Errorf("Expected classes to be ['bold', 'italic'], got %v", text.Classes)
	}
}

func TestWithAttribute(t *testing.T) {
	text := NewText("Test").WithAttribute("data-id", "123").WithAttribute("aria-label", "Test Label")
	if len(text.Attributes) != 2 {
		t.Errorf("Expected 2 attributes, got %d", len(text.Attributes))
	}
	if text.Attributes["data-id"] != "123" || text.Attributes["aria-label"] != "Test Label" {
		t.Errorf("Expected attributes to be {'data-id': '123', 'aria-label': 'Test Label'}, got %v", text.Attributes)
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name     string
		text     *Text
		expected string
	}{
		{
			name:     "Basic paragraph",
			text:     NewText("Hello, World!"),
			expected: "<p>Hello, World!</p>",
		},
		{
			name:     "Heading with class",
			text:     NewText("Title").WithTag("h1").WithClass("main-title"),
			expected: `<h1 class="main-title">Title</h1>`,
		},
		{
			name:     "Span with multiple classes and attribute",
			text:     NewText("Important").WithTag("span").WithClass("bold").WithClass("red").WithAttribute("data-priority", "high"),
			expected: `<span class="bold red" data-priority="high">Important</span>`,
		},
		{
			name:     "Empty content",
			text:     NewText("").WithTag("br"),
			expected: "<br></br>",
		},
		{
			name:     "Escaping HTML content",
			text:     NewText("<script>alert('XSS');</script>"),
			expected: "<p>&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.text.Render(&buf)
			if err != nil {
				t.Fatalf("Render returned an error: %v", err)
			}
			result := buf.String()
			if components.NormalizeWhitespace(result) != components.NormalizeWhitespace(tt.expected) {
				t.Errorf("Expected rendered text to be '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestRenderError(t *testing.T) {
	// This test is to ensure that the Render method doesn't panic with nil values
	text := &Text{}
	var buf bytes.Buffer
	err := text.Render(&buf)
	if err != nil {
		t.Errorf("Expected no error for nil values, got: %v", err)
	}
}
