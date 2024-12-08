// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package button

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/bartalcorn/gothat/pkg/htmx"
	"github.com/bartalcorn/gothat/pkg/tws"
)

type Variant string
type Size string

type Props struct {
	// Variant determines the visual style of the button.
	Variant Variant `default:"primary"`

	// Label is the content of the button.
	// Default: "" (empty string)
	Label string

	// Size sets the size of the button.
	Size Size `default:"md"`

	// FullWidth determines whether the button should take up the full width of its container.
	// Default: false
	FullWidth bool

	// Click is the JavaScript function to be called when the button is clicked. Usuallu by AlpineJS
	// Default: ""
	Click string

	// Prevent is a bool to indicate whether the default action should occur onClick
	// Default: false
	Prevent bool

	// HX allows passing additional HTMX attributes to the button or anchor element.
	// Default: nil
	HX htmx.Props

	// Disabled can be either a bool or a string.
	// If bool (Go), it directly controls the disabled state.
	// If string, it's treated as a JavaScript expression for dynamic disabling.
	Disabled bool

	// IconLeft specifies an icon component to be displayed on the left side of the button text.
	// Use IconLeft with a label for an Icon button
	// Default: nil
	IconLeft templ.Component

	// IconRight specifies an icon component to be displayed on the right side of the button text.
	// Default: nil
	IconRight templ.Component

	// Content - for truly dynamic button content
	// Default: nil
	Content templ.Component

	// BoundClass specifies additional CSS classes to apply when the bounded condition is met.
	// Default: "" (empty string)
	BoundClass string

	// Class specifies additional CSS classes to apply to the button to override any previous class.
	// Default: "" (empty string)
	Class string

	// Attrs allows passing additional HTML attributes to the button or anchor element.
	// Default: nil
	Attrs templ.Attributes
}

const (
	base           string = "inline-block items-center justify-center space-x-4 space-y-3 whitespace-nowrap text-base leading-4 transition-colors no-ring"
	diabled        string = " disabled:text-white disabled:border-gray-400 disabled:bg-gray-400 disabled:hover:border-gray-600 disabled:hover:bg-gray-600"
	diabledSpecial string = " disabled:border-gray-600 disabled:bg-gray-600"

	Primary   Variant = "primary"
	Secondary Variant = "secondary"
	Text      Variant = "text"
	Delete    Variant = "delete"
	Error     Variant = "error"
	Warning   Variant = "warning"
	Success   Variant = "success"
	Info      Variant = "info"
	Icon      Variant = "icon"
	Gray      Variant = "gray"

	Md Size = "md"
	Sm Size = "sm"
	Lg Size = "lg"
)

func (p *Props) variant() string {
	switch p.Variant {
	case Secondary:
		return "border border-blue-500 bg-white text-blue-500 hover:bg-blue-50" + diabled
	case Text:
		return "text-base hover:text-blue-600"
	case Delete:
		return "border border-rose-500 bg-white text-rose-500 hover:bg-rose-500-light disabled:border-gray-600 disabled:bg-transparent"
	case Error:
		if p.Label != "" {
			return "border border-rose-500 bg-white text-rose-500 bg-rose-500-light" + diabledSpecial
		}
		return "text-rose-500 hover:text-blue-500 disabled:text-gray-400"
	case Icon:
		return "text-gray-300 hover:text-blue-500"
	case Warning:
		if p.Label != "" {

			return "border border-yellow-500 bg-white text-yellow-500 bg-yellow-500-light" + diabledSpecial
		}
		return "text-yellow-500 hover:text-blue-500 disabled:text-gray-400"
	case Success:
		if p.Label != "" {
			return "border border-green-500 bg-white text-green-500 bg-green-500-light" + diabledSpecial
		}
		return "text-green-500 hover:text-blue-500 disabled:text-gray-400"
	case Info:
		if p.Label != "" {
			return "border border-blue-500 bg-white text-blue-500 bg-blue-500-light" + diabledSpecial
		}
		return "text-blue-500 hover:text-blue-500 disabled:text-gray-400"
	default:
		return "inline-flex w-fit items-center justify-between gap-2 whitespace-nowrap border-zinc-300 bg-zinc-100 px-4 py-2 text-sm font-medium capitalize tracking-wide text-neutral-600 transition hover:opacity-75 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-700 rounded-lg border" + diabled
	}
}

func (p *Props) size() string {
	if p.Variant == Icon {
		return "px-1 py-2 text-base"
	}
	switch p.Size {
	case Sm:
		return "px-3 py-2 text-xs"
	case Md:
		return "px-3 py-2 text-sm"
	default:
		return "px-3 py-2 text-base"
	}
}

func Button(props *Props) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{tws.TwMerge(base, " rounded-sm", props.variant(), props.size(), tws.NoRing, props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 1)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.Attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Click != "" {
			if props.Prevent {
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 2)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.Click)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/button/button.templ`, Line: 145, Col: 37}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 3)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 4)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.Click)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/button/button.templ`, Line: 147, Col: 29}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 5)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/button/button.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 7)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.BoundClass != "" {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 8)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.BoundClass)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/button/button.templ`, Line: 152, Col: 28}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 9)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.HX.HasMethod() {
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.HX.Attrbs())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Disabled {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 10)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 11)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = renderButtonContent(props).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func renderButtonContent(props *Props) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 13)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.IconLeft != nil {
			templ_7745c5c3_Err = props.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Content != nil {
			templ_7745c5c3_Err = props.Content.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/button/button.templ`, Line: 174, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 14)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.IconRight != nil {
			templ_7745c5c3_Err = props.IconRight.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 15)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
