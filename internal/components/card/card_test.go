package card

import (
	"bytes"
	"testing"

	"github.com/billvamva/gomp/internal/components"
)

// MockComponent is a simple mock implementation of the Component interface for testing
type MockComponent struct {
	Content string
}

func (m *MockComponent) Render(buf *bytes.Buffer) error {
	_, err := buf.WriteString(m.Content)
	return err
}

func TestNewProjectCard(t *testing.T) {
	card := NewProjectCard("Test Title", "Test Description", "test-class")

	if card.Title != "Test Title" {
		t.Errorf("Expected Title to be 'Test Title', got '%s'", card.Title)
	}
	if card.Description != "Test Description" {
		t.Errorf("Expected Description to be 'Test Description', got '%s'", card.Description)
	}
	if card.Class != "test-class" {
		t.Errorf("Expected Class to be 'test-class', got '%s'", card.Class)
	}
	if len(card.Components) != 0 {
		t.Errorf("Expected Components to be empty, got %d components", len(card.Components))
	}
}

func TestAddComponent(t *testing.T) {
	card := NewProjectCard("Test", "Test", "test")
	mockComponent := &MockComponent{Content: "Mock Content"}

	card.AddComponent(mockComponent)

	if len(card.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(card.Components))
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name     string
		card     *Card
		expected string
	}{
		{
			name:     "Basic card",
			card:     NewProjectCard("Test Title", "Test Description", "test-class"),
			expected: `<div class=test-class><h3>Test Title</h3><p>Test Description</p></div>`,
		},
		{
			name: "Card with component",
			card: func() *Card {
				c := NewProjectCard("With Component", "Has a mock component", "component-class")
				c.AddComponent(&MockComponent{Content: "<span>Mock Component</span>"})
				return c
			}(),
			expected: `<div class=component-class><h3>With Component</h3><p>Has a mock component</p><span>Mock Component</span></div>`,
		},
		{
			name: "Card with multiple components",
			card: func() *Card {
				c := NewProjectCard("Multiple", "Multiple components", "multi-class")
				c.AddComponent(&MockComponent{Content: "<span>First</span>"})
				c.AddComponent(&MockComponent{Content: "<span>Second</span>"})
				return c
			}(),
			expected: `<div class=multi-class><h3>Multiple</h3><p>Multiple components</p><span>First</span><span>Second</span></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.card.Render(&buf)
			if err != nil {
				t.Fatalf("Render returned an error: %v", err)
			}
			result := components.NormalizeWhitespace(buf.String())
			expected := components.NormalizeWhitespace(tt.expected)
			if result != expected {
				t.Errorf("Expected rendered card to be '%s', got '%s'", expected, result)
			}
		})
	}
}
