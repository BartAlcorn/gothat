package tmpls

import "github.com/a-h/templ"

type Props struct {
	Content  templ.Component
	Contents []templ.Component
	Module   string
	Attrs    templ.Attributes
}
