package formats

import (
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

func UpperCase(str string) string {
	caser := cases.Upper(language.English)
	return caser.String(str)
}

func FormatDate(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format(time.RFC3339)
}
