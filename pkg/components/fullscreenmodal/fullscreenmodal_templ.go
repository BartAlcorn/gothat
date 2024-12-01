// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package fullscreenmodal

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/bartalcorn/lucide"

type Props struct {
	// ID of the Modal
	// default: ""
	ID string

	// Title to display
	// default: ""
	Title string

	// text to display on trigger button
	Label string
}

func FullScreenModal(props *Props) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.ID != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ID=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/fullscreenmodal/fullscreenmodal.templ`, Line: 21, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-data=\"{\n\t\t\tfullscreenModal: false,\n\t\t\ttitle: &#39;&#39;,\n\t\t\topen(ev) {\n\t\t\t\tthis.fullscreenModal = true;\n\t\t\t\tthis.title = ev.title;\n\t\t\t\thtmx.ajax(ev.method, ev.fetch, {\n\t\t\t\t\t\ttarget: &#39;#full-screen-modal-body&#39;,\n\t\t\t\t\t\tswap: &#39;innerHTML&#39;,\n\t\t\t\t\t})\n\t\t\t}\n\t\t }\" @keyup.escape.document=\"fullscreenModal = false;\" @spl:full-screen-modal-open.document=\"open($event.detail);\"><button @click=\"fullscreenModal=true\" class=\"got-component\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/fullscreenmodal/fullscreenmodal.templ`, Line: 38, Col: 75}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><template x-teleport=\"body\"><div x-show=\"fullscreenModal\" x-transition:enter=\"transition ease-out duration-100\" x-transition:enter-start=\"opacity-0\" x-transition:enter-end=\"opacity-100\" x-transition:leave=\"transition ease-in duration-100\" x-transition:leave-start=\"opacity-100\" x-transition:leave-end=\"opacity-0\" class=\"flex flex-col gap-2 fixed inset-0 z-50 w-screen top-0 bottom-0 bg-gray-100\"><div class=\"flex-grow min-h-0 overflow-y-hidden\"><div class=\"flex gap-2 h-12 w-full px-4 items-center justify-between bg-gray-200 border-b border-gray-400\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Title != "" {
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/fullscreenmodal/fullscreenmodal.templ`, Line: 54, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span x-text=\"title\" class=\"flex-nowrap w-full text-base font-semibold text-gray-100 \"></span></div><button @click=\"fullscreenModal=false\" class=\"flex items-center justify-center text-gray-800 hover:text-blue-800\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = lucide.SquareX(lucide.Props{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div><div id=\"full-screen-modal-body\" class=\"pt-4 px-4 overflow-y-auto max-h-[760px] inset-4 w-full\"></div></div></div></template></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate