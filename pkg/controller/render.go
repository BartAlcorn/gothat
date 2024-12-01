package controller

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/bartalcorn/gothat/pkg/controller/tmpls"
)

func Render(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	if ctx.Value("module") != nil {
		m := ctx.Value("module").(string)
		SetModule(m)
		if IsHTMX(r) {
			ctx = AddContent(ctx, tmpls.Navigation(m, true))
		}
	}
	if ctx.Value("content") != nil {
		fmt.Println("content", len(ctx.Value("content").([]templ.Component)))
	}

	if IsHTMX(r) {
		for _, tc := range ctx.Value("content").([]templ.Component) {
			err := tc.Render(ctx, w)
			if err != nil {
				fmt.Println("ERROR rendering template:", err)
			}
		}
		fmt.Println("*** HTMX ***")
		return
	}

	err := tmpls.Base(&tmpls.Props{
		Content: tmpls.Main(&tmpls.Props{
			Contents: ctx.Value("content").([]templ.Component),
			Module:   module,
		}),
	}).Render(ctx, w)
	if err != nil {
		fmt.Println("ERROR rendering template:", err)
	}
}

func IsHTMX(r *http.Request) bool {
	return r.Header.Get("Hx-Request") == "true"
}
