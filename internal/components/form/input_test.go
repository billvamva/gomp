package form

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewInput(t *testing.T) {
	i := NewInput("test-input")
	if i.Name != "test-input" {
		t.Errorf("Expected name 'test-input', got '%s'", i.Name)
	}
	if len(i.Classes) != 0 {
		t.Errorf("Expected 0 classes, got %d", len(i.Classes))
	}
	if len(i.Attributes) != 0 {
		t.Errorf("Expected 0 attributes, got %d", len(i.Attributes))
	}
}

func TestInputWithClass(t *testing.T) {
	i := NewInput("test").WithClass("input-class")
	if len(i.Classes) != 1 || i.Classes[0] != "input-class" {
		t.Errorf("Expected class 'input-class', got %v", i.Classes)
	}
}

func TestInputWithAttribute(t *testing.T) {
	i := NewInput("test").WithAttribute("type", "text")
	if len(i.Attributes) != 1 || i.Attributes["type"] != "text" {
		t.Errorf("Expected attribute 'type: text', got %v", i.Attributes)
	}
}

func TestInputRender(t *testing.T) {
	i := NewInput("test-input").
		WithClass("form-control").
		WithAttribute("type", "text").
		WithAttribute("placeholder", "Enter text")

	var buf bytes.Buffer
	err := i.Render(&buf)
	if err != nil {
		t.Fatalf("Render returned an error: %v", err)
	}

	expected := `<input class="form-control" type="text" placeholder="Enter text">`
	result := strings.TrimSpace(buf.String())
	if result != expected {
		t.Errorf("Expected rendered input to be '%s', got '%s'", expected, result)
	}
}

func TestInputRenderEscaping(t *testing.T) {
	i := NewInput("test").
		WithAttribute("data-value", "a & b")

	var buf bytes.Buffer
	err := i.Render(&buf)
	if err != nil {
		t.Fatalf("Render returned an error: %v", err)
	}

	expected := `<input data-value="a &amp; b">`
	result := strings.TrimSpace(buf.String())
	if result != expected {
		t.Errorf("Expected rendered input to be '%s', got '%s'", expected, result)
	}
}
