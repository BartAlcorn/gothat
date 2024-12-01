package showcase

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/bartalcorn/gothat/pkg/components/breadcrumb"
	"github.com/bartalcorn/gothat/pkg/components/button"
	"github.com/bartalcorn/gothat/pkg/components/checkbox"
	"github.com/bartalcorn/gothat/pkg/components/combobox"
	"github.com/bartalcorn/gothat/pkg/components/dialog"
	"github.com/bartalcorn/gothat/pkg/components/disclosure"
	"github.com/bartalcorn/gothat/pkg/components/dropdown"
	"github.com/bartalcorn/gothat/pkg/components/menu"
	"github.com/bartalcorn/gothat/pkg/components/modal"
	"github.com/bartalcorn/gothat/pkg/components/popover"
	"github.com/bartalcorn/gothat/pkg/components/radio"
	"github.com/bartalcorn/gothat/pkg/components/toggle"
	"github.com/bartalcorn/gothat/pkg/controller"
)

type handlers struct{}

var handle *handlers

type Handler interface {
	base(w http.ResponseWriter, r *http.Request)
}

func (h *handlers) base(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "module", "Showcase")
	ctx = controller.AddContent(ctx, base())
	controller.Render(w, r.WithContext(ctx))
}

func (h *handlers) element(w http.ResponseWriter, r *http.Request) {
	el := r.PathValue("element")
	ctx := controller.AddContent(r.Context(), elements[el])
	controller.Render(w, r.WithContext(ctx))
}

var elements = map[string]templ.Component{
	"breadcrumb": breadcrumb.Showcase(),
	"button":     button.Showcase(),
	"checkbox":   checkbox.Showcase(),
	"combobox":   combobox.Showcase(),
	"dialog":     dialog.Showcase(),
	"disclosure": disclosure.Showcase(),
	"dropdown":   dropdown.Showcase(),
	"menu":       menu.Showcase(),
	"modal":      modal.Showcase(),
	"popover":    popover.Showcase(),
	"radio":      radio.Showcase(),
	"toggle":     toggle.Showcase(),
}
