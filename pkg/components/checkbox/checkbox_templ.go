// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package checkbox

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type Props struct {
	Name        string
	Label       string
	LabelLeft   string
	Classes     string
	Checked     bool
	Description string
	OnChange    string
	Click       string
	XModel      string
}

func Checkbox(props *Props) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{props.Classes, "flex justify-start"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex items-baseline h-5\"><input id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 19, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 20, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"peer w-5 h-5 rounded-sm border-2 border-gray-300 bg-gray-700 hover:border-blue-500 focus:ring-3 focus:ring-blue-500\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Checked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.XModel != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-model=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.XModel)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 27, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Click != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" @click=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(props.Click)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 30, Col: 25}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("></div><div class=\"ml-2 text-base\"><label for=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 35, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-gray-200 peer-checked:text-blue-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 35, Col: 91}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Checkbox2(props *Props) templ.Component {
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
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-start\"><label for=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 43, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex cursor-pointer items-center gap-2 text-base font-medium text-neutral-600  [&amp;:has(input:checked)]:text-neutral-900 dark:[&amp;:has(input:checked)]:text-zinc-50 [&amp;:has(input:disabled)]:opacity-75 [&amp;:has(input:disabled)]:cursor-not-allowed\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.LabelLeft != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(props.LabelLeft)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 47, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative flex items-center\"><input id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 51, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"before:content[&#39;&#39;] peer relative size-5 cursor-pointer appearance-none overflow-hidden rounded border border-zinc-300 bg-zinc-100 before:absolute before:inset-0 checked:border-sky-700 checked:before:bg-sky-700 focus:outline focus:outline-2 focus:outline-offset-2 focus:outline-zinc-500 checked:focus:outline-sky-700 active:outline-offset-0 disabled:cursor-not-allowed dark:border-zinc-700 dark:bg-zinc-800 dark:checked:border-sky-600 dark:checked:before:bg-sky-600 dark:focus:outline-zinc-500 dark:checked:focus:outline-sky-600\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Checked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" aria-hidden=\"true\" stroke=\"currentColor\" fill=\"none\" stroke-width=\"4\" class=\"pointer-events-none invisible absolute left-1/2 top-1/2 size-4 -translate-x-1/2 -translate-y-1/2 text-white peer-checked:visible dark:text-white\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M4.5 12.75l6 6 9-13.5\"></path></svg></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Label != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 63, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Description != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span id=\"checkboxDescription\" class=\"ml-6 text-sm text-neutral-600 dark:text-zinc-200\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(props.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/checkbox/checkbox.templ`, Line: 71, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Containered(props *Props) templ.Component {
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
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label for=\"checkboxDefault\" class=\"inline-flex min-w-[14rem] cursor-pointer items-center justify-between rounded-lg gap-3 border border-zinc-300 bg-zinc-100 px-4 py-2 text-sm font-medium text-neutral-600 dark:border-zinc-700 dark:bg-zinc-800 dark:text-zinc-200 [&amp;:has(input:checked)]:text-neutral-900 dark:[&amp;:has(input:checked)]:text-zinc-50 [&amp;:has(input:disabled)]:opacity-75 [&amp;:has(input:disabled)]:cursor-not-allowed\"><span>Notifications</span><div class=\"relative flex items-center\"><input id=\"checkboxDefault\" type=\"checkbox\" class=\"before:content[&#39;&#39;] peer relative size-4 cursor-pointer appearance-none overflow-hidden rounded border border-zinc-300 bg-zinc-50 before:absolute before:inset-0 checked:border-sky-700 checked:before:bg-sky-700 focus:outline focus:outline-2 focus:outline-offset-2 focus:outline-zinc-500 checked:focus:outline-sky-700 active:outline-offset-0 disabled:cursor-not-allowed dark:border-zinc-700 dark:bg-zinc-900 dark:checked:border-sky-600 dark:checked:before:bg-sky-600 dark:focus:outline-zinc-500 dark:checked:focus:outline-sky-600\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Checked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" aria-hidden=\"true\" stroke=\"currentColor\" fill=\"none\" stroke-width=\"4\" class=\"pointer-events-none invisible absolute left-1/2 top-1/2 size-3 -translate-x-1/2 -translate-y-1/2 text-white peer-checked:visible dark:text-white\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M4.5 12.75l6 6 9-13.5\"></path></svg></div></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
