package formats

import (
	"strings"
)

func StripFirstStringElement(str string, sep string) string {
	s := strings.Split(str, sep)
	return strings.Join(s[1:], sep)
}

func LastStringElement(str string, sep string) string {
	s := strings.Split(str, sep)
	return s[len(s)-1]
}
