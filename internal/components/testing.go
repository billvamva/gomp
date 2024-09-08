package components

import (
	"regexp"
	"strings"
)

func NormalizeWhitespace(s string) string {
	// Remove all newlines and extra spaces
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	// Remove spaces after '<' and before '>'
	s = regexp.MustCompile(`\s*<\s*`).ReplaceAllString(s, "<")
	s = regexp.MustCompile(`\s*>\s*`).ReplaceAllString(s, ">")
	return s
}
