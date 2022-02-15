package gnt

import (
	"strings"
)

var (
	escapeMarkdownV2 = []string{
		"_", "*", "[", "]", "(", ")", "~", "`", ">",
		"#", "+", "-", "=", "|", "{", "}", ".", "!",
	}
	removeMarkdown = []string{
		"\r", "\n", "[", "]", "(", ")", "`", ">", "#", "-",
	}
)

func EscapedMarkdownV2(s string) string {
	for _, v := range escapeMarkdownV2 {
		s = strings.ReplaceAll(s, v, `\`+v)
	}
	return s
}

func FilterBody(s string) string {
	if s == "" {
		return "\n\n"
	}
	for _, v := range removeMarkdown {
		s = strings.ReplaceAll(s, v, "")
	}
	if len(s) > 1414 {
		s = s[:1414]
	}
	return "\n\n" + EscapedMarkdownV2(s) + "\n\n"
}
