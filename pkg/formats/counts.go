package formats

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatCount(count int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", count)
}

func PluralizeRecordCount(total int) string {
	message.Set(language.English, "%f%% records",
		plural.Selectf(total, "%d",
			"=1", "%[1]d record.",
			"other", "%[1]d records.",
		),
	)
	p := message.NewPrinter(language.English)
	return p.Sprintf("%f%% records", total)
}
