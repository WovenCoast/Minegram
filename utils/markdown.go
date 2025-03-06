package utils

import "strings"

// EscapeMarkdown escapes characters that have special meaning in Markdown
// to prevent user-provided text from being interpreted as Markdown formatting
func EscapeMarkdown(text string) string {
	// Characters to escape: _ * [ ] ( ) ~ ` > # + - = | { } . !
	replacements := []struct {
		from string
		to   string
	}{
		{"\\", "\\\\"},
		{"_", "\\_"},
		{"*", "\\*"},
		{"[", "\\["},
		{"]", "\\]"},
		{"(", "\\("},
		{")", "\\)"},
		{"~", "\\~"},
		{"`", "\\`"},
		{">", "\\>"},
		{"#", "\\#"},
		{"+", "\\+"},
		{"-", "\\-"},
		{"=", "\\="},
		{"|", "\\|"},
		{"{", "\\{"},
		{"}", "\\}"},
		{".", "\\."},
		{"!", "\\!"},
	}

	escapedText := text
	for _, r := range replacements {
		escapedText = strings.ReplaceAll(escapedText, r.from, r.to)
	}

	return escapedText
}
