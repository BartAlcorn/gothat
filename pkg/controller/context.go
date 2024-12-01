package controller

import (
	"context"

	"github.com/a-h/templ"
)

func AddContent(ctx context.Context, content templ.Component) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value("content") == nil {
		cslice := []templ.Component{content}
		ctx = context.WithValue(ctx, "content", cslice)
		return ctx
	}

	cslice := ctx.Value("content").([]templ.Component)
	cslice = append(cslice, content)
	return context.WithValue(ctx, "content", cslice)
}
