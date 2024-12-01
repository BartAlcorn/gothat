package home

import (
	"context"
	"net/http"

	"github.com/bartalcorn/gothat/pkg/controller"
)

type handlers struct{}

var handle *handlers

type Handler interface {
	base(w http.ResponseWriter, r *http.Request)
}

func (h *handlers) base(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "module", "Home")
	ctx = controller.AddContent(ctx, base())
	controller.Render(w, r.WithContext(ctx))
}
