package formats

import (
	"strings"
)

func CommaList(items []string) string {
	var sb strings.Builder
	for i, item := range items {
		sb.WriteString(item)
		if i < len(items)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

func TitleCaseCommaList(items []string) string {
	var sb strings.Builder
	for i, item := range items {
		sb.WriteString(TitleCase(item))
		if i < len(items)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
