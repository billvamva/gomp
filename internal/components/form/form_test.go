package form

import (
	"bytes"
	"strings"
	"testing"
)

// Mock component for testing
type MockComponent struct {
	Content string
}

func (m *MockComponent) Render(buf *bytes.Buffer) error {
	_, err := buf.WriteString(m.Content)
	return err
}

// Form Tests
func TestNewForm(t *testing.T) {
	f := NewForm()
	if f == nil {
		t.Error("NewForm() returned nil")
	}
	if len(f.Components) != 0 {
		t.Errorf("Expected 0 components, got %d", len(f.Components))
	}
	if len(f.Classes) != 0 {
		t.Errorf("Expected 0 classes, got %d", len(f.Classes))
	}
	if len(f.Attributes) != 0 {
		t.Errorf("Expected 0 attributes, got %d", len(f.Attributes))
	}
}

func TestFormAddComponent(t *testing.T) {
	f := NewForm()
	mock := &MockComponent{Content: "Test"}
	f.AddComponent(mock)
	if len(f.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(f.Components))
	}
}

func TestFormWithClass(t *testing.T) {
	f := NewForm().WithClass("form-class")
	if len(f.Classes) != 1 || f.Classes[0] != "form-class" {
		t.Errorf("Expected class 'form-class', got %v", f.Classes)
	}
}

func TestFormWithAttribute(t *testing.T) {
	f := NewForm().WithAttribute("method", "post")
	if len(f.Attributes) != 1 || f.Attributes["method"] != "post" {
		t.Errorf("Expected attribute 'method: post', got %v", f.Attributes)
	}
}

func TestFormRender(t *testing.T) {
	f := NewForm().
		WithClass("test-form").
		WithAttribute("action", "/submit").
		WithAttribute("method", "post")
	f.AddComponent(&MockComponent{Content: "<input type=\"text\" name=\"test\">"})

	var buf bytes.Buffer
	err := f.Render(&buf)
	if err != nil {
		t.Fatalf("Render returned an error: %v", err)
	}

	expected := `<form class="test-form" action="/submit" method="post"><input type="text" name="test"></form>`
	result := strings.TrimSpace(buf.String())
	if result != expected {
		t.Errorf("Expected rendered form to be '%s', got '%s'", expected, result)
	}
}
