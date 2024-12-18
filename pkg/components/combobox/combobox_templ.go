// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package combobox

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"encoding/json"
	"github.com/bartalcorn/gothat/pkg/errhandler"
)

type Props struct {
	ID                string  // form ID
	Name              string  // form name
	Value             *Item   // in single mode, the value of the pre-selected item
	Selected          []*Item // in Mulitple Mode, the pre-selected item(s)
	SelectedPlacement string  // where the selected chips appears in relation to the trigger
	Multiple          bool    // support multiple selections
	By                string  // compare the objects by a particular property instead of by their identity
	Items             []*Item // slice of Items, the options to be displayed
}

type Item struct {
	Key      string `json:"label"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

func options(items []*Item) string {
	itm, err := json.Marshal(items)
	errhandler.OnErr("error parsing", err)
	return string(itm)
}

func (props *Props) selected() string {
	if props.Multiple {
		if len(props.Selected) == 0 {
			return "[]"
		}
		return options(props.Selected)
	}
	if props.Value != nil {
		itm, err := json.Marshal(props.Value)
		errhandler.OnErr("error parsing", err)
		return string(itm)
	}
	return "''"
}

func ComboBox(props *Props) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{\n        options: [\n            {\n                value: &#39;Agriculture&#39;,\n                label: &#39;Agriculture&#39;,\n            },\n            {\n                value: &#39;Construction&#39;,\n                label: &#39;Construction&#39;,\n            },\n            {\n                value: &#39;Education&#39;,\n                label: &#39;Education&#39;,\n            },\n            {\n                value: &#39;Entertainment&#39;,\n                label: &#39;Entertainment&#39;,\n            },\n            {\n                value: &#39;Finance&#39;,\n                label: &#39;Finance&#39;,\n            },\n            {\n                value: &#39;Healthcare&#39;,\n                label: &#39;Healthcare&#39;,\n            },\n            {\n                value: &#39;Hospitality&#39;,\n                label: &#39;Hospitality&#39;,\n            },\n            {\n                value: &#39;IT&#39;,\n                label: &#39;IT&#39;,\n            },\n            {\n                value: &#39;Manufacturing&#39;,\n                label: &#39;Manufacturing&#39;,\n            },\n            {\n                value: &#39;Marketing&#39;,\n                label: &#39;Marketing&#39;,\n            },\n            {\n                value: &#39;Real Estate&#39;,\n                label: &#39;Real Estate&#39;,\n            },\n            {\n                value: &#39;Retail&#39;,\n                label: &#39;Retail&#39;,\n            },\n            {\n                value: &#39;Transportation&#39;,\n                label: &#39;Transportation&#39;,\n            },\n        ],\n        isOpen: false,\n        openedWithKeyboard: false,\n        selectedOption: null,\n        setSelectedOption(option) {\n            this.selectedOption = option\n            this.isOpen = false\n            this.openedWithKeyboard = false\n            this.$refs.hiddenTextField.value = option.value\n        },\n        highlightFirstMatchingOption(pressedKey) {\n            const option = this.options.find((item) =&gt;\n                item.label.toLowerCase().startsWith(pressedKey.toLowerCase()),\n            )\n            if (option) {\n                const index = this.options.indexOf(option)\n                const allOptions = document.querySelectorAll(&#39;.combobox-option&#39;)\n                if (allOptions[index]) {\n                    allOptions[index].focus()\n                }\n            }\n        },\n    }\" class=\"w-full max-w-xs flex flex-col gap-1\" x-on:keydown=\"highlightFirstMatchingOption($event.key)\" x-on:keydown.esc.window=\"isOpen = false, openedWithKeyboard = false\"><label for=\"industry\" class=\"w-fit pl-0.5 text-sm text-neutral-600 dark:text-zinc-200\">Industry</label><div class=\"relative\"><!-- trigger button  --><button type=\"button\" role=\"combobox\" class=\"inline-flex w-full items-center justify-between gap-2 whitespace-nowrap border-zinc-300 bg-zinc-100 px-4 py-2 text-sm font-medium capitalize tracking-wide text-neutral-600 transition hover:opacity-75 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-700 dark:border-zinc-700 dark:bg-zinc-800/50 dark:text-zinc-200 dark:focus-visible:outline-sky-600 rounded-lg border\" aria-haspopup=\"listbox\" aria-controls=\"industriesList\" x-on:click=\"isOpen = ! isOpen\" x-on:keydown.down.prevent=\"openedWithKeyboard = true\" x-on:keydown.enter.prevent=\"openedWithKeyboard = true\" x-on:keydown.space.prevent=\"openedWithKeyboard = true\" x-bind:aria-label=\"selectedOption ? selectedOption.value : &#39;Please Select&#39;\" x-bind:aria-expanded=\"isOpen || openedWithKeyboard\"><span class=\"text-sm font-normal\" x-text=\"selectedOption ? selectedOption.value : &#39;Please Select&#39;\"></span><!-- Chevron  --><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" class=\"size-5\"><path fill-rule=\"evenodd\" d=\"M5.22 8.22a.75.75 0 0 1 1.06 0L10 11.94l3.72-3.72a.75.75 0 1 1 1.06 1.06l-4.25 4.25a.75.75 0 0 1-1.06 0L5.22 9.28a.75.75 0 0 1 0-1.06Z\" clip-rule=\"evenodd\"></path></svg></button><!-- hidden input to grab the selected value  --><input id=\"industry\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/combobox/combobox.templ`, Line: 140, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"text\" x-ref=\"hiddenTextField\" hidden><ul x-cloak x-show=\"isOpen || openedWithKeyboard\" id=\"industriesList\" class=\"absolute z-10 left-0 top-11 flex max-h-44 w-full flex-col overflow-hidden overflow-y-auto border-zinc-300 bg-zinc-100 py-1.5 dark:border-zinc-700 dark:bg-zinc-800 rounded-lg border\" role=\"listbox\" aria-label=\"industries list\" x-on:click.outside=\"isOpen = false, openedWithKeyboard = false\" x-on:keydown.down.prevent=\"$focus.wrap().next()\" x-on:keydown.up.prevent=\"$focus.wrap().previous()\" x-transition x-trap=\"openedWithKeyboard\"><template x-for=\"(item, index) in options\" x-bind:key=\"item.value\"><li class=\"combobox-option inline-flex cursor-pointer justify-between gap-6 bg-zinc-100 px-4 py-2 text-sm text-neutral-600 hover:bg-zinc-800/5 hover:text-neutral-900 focus-visible:bg-zinc-800/5 focus-visible:text-neutral-900 focus-visible:outline-none dark:bg-zinc-800 dark:text-zinc-200 dark:hover:bg-zinc-100/5 dark:hover:text-zinc-50 dark:focus-visible:bg-zinc-100/10 dark:focus-visible:text-zinc-50\" role=\"option\" x-on:click=\"setSelectedOption(item)\" x-on:keydown.enter=\"setSelectedOption(item)\" x-bind:id=\"&#39;option-&#39; + index\" tabindex=\"0\"><!-- Label  --><span x-bind:class=\"selectedOption == item ? &#39;font-bold&#39; : null\" x-text=\"item.label\"></span><!-- Screen reader 'selected' indicator  --><span class=\"sr-only\" x-text=\"selectedOption == item ? &#39;selected&#39; : null\"></span><!-- Checkmark  --><svg x-cloak x-show=\"selectedOption == item\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" fill=\"none\" stroke-width=\"2\" class=\"size-4\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m4.5 12.75 6 6 9-13.5\"></path></svg></li></template></ul></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
