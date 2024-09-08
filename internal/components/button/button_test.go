package button

import (
	"bytes"
	"strings"
	"testing"
)

func TestWithContent(t *testing.T) {
	b := NewButton("").WithContent("New Content")
	if b.Text != "New Content" {
		t.Errorf("Expected Text to be 'New Content', got '%s'", b.Text)
	}
}

func TestWithClass(t *testing.T) {
	b := NewButton("").WithClass("btn-primary").WithClass("large")
	if len(b.Classes) != 2 {
		t.Errorf("Expected 2 classes, got %d", len(b.Classes))
	}
	if b.Classes[0] != "btn-primary" || b.Classes[1] != "large" {
		t.Errorf("Expected classes to be ['btn-primary', 'large'], got %v", b.Classes)
	}
}

func TestWithAttribute(t *testing.T) {
	b := NewButton("").
		WithAttribute("hx-post", "/action").
		WithAttribute("hx-target", "#result")

	if len(b.Attributes) != 2 {
		t.Errorf("Expected 2 attributes, got %d", len(b.Attributes))
	}
	if b.Attributes["hx-post"] != "/action" || b.Attributes["hx-target"] != "#result" {
		t.Errorf("Expected attributes to be {'hx-post': '/action', 'hx-target': '#result'}, got %v", b.Attributes)
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name     string
		button   *Button
		expected string
	}{
		{
			name:     "Basic button",
			button:   NewButton("Click me").WithClass("btn"),
			expected: `<button class="btn">Click me</button>`,
		},
		{
			name: "Button with multiple classes and attributes",
			button: NewButton("Submit").
				WithClass("btn").
				WithClass("btn-primary").
				WithAttribute("hx-post", "/submit").
				WithAttribute("hx-target", "#result"),
			expected: `<button class="btn btn-primary" hx-post="/submit" hx-target="#result">Submit</button>`,
		},
		{
			name: "Button with special characters",
			button: NewButton("Click & Submit").
				WithAttribute("data-value", "a & b"),
			expected: `<button data-value="a &amp; b">Click &amp; Submit</button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.button.Render(&buf)
			if err != nil {
				t.Fatalf("Render returned an error: %v", err)
			}
			result := strings.TrimSpace(buf.String())
			if result != tt.expected {
				t.Errorf("Expected rendered button to be '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
