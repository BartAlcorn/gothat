package htmx

import "github.com/a-h/templ"

type Props struct {
	AfterRequest string
	Confirm      string
	Delete       string
	Get          string
	Include      string
	Options      string
	Patch        string
	Post         string
	PushUrl      string
	Put          string
	Swap         string
	Target       string
	Trigger      string
}

func (p *Props) HasMethod() bool {
	return p.Delete != "" || p.Get != "" || p.Options != "" || p.Patch != "" || p.Post != "" || p.Put != ""
}

func (p *Props) Attrbs() templ.Attributes {
	attrbs := make(templ.Attributes)
	if p.Delete != "" {
		attrbs["hx-delete"] = p.Delete
	}
	if p.Get != "" {
		attrbs["hx-get"] = p.Get
	}
	if p.Post != "" {
		attrbs["hx-post"] = p.Post
	}
	if p.Put != "" {
		attrbs["hx-put"] = p.Put
	}
	if p.Patch != "" {
		attrbs["hx-patch"] = p.Patch
	}
	if p.Target != "" {
		attrbs["hx-target"] = p.Target
	}
	if p.Trigger != "" {
		attrbs["hx-trigger"] = p.Trigger
	}
	if p.Swap != "" {
		attrbs["hx-swap"] = p.Swap
	}
	if p.Confirm != "" {
		attrbs["hx-confirm"] = p.Confirm
	}
	if p.Include != "" {
		attrbs["hx-include"] = p.Include
	}
	if p.PushUrl != "" {
		attrbs["hx-push-url"] = p.PushUrl
	}
	if p.AfterRequest != "" {
		attrbs["hx-on::after-request"] = p.AfterRequest
	}
	return attrbs
}
