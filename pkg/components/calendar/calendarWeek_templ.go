// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package calendar

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CalendarWeek() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"lg:flex lg:h-full lg:flex-col\"><header class=\"flex items-center justify-between border-b border-gray-200 px-6 py-4 lg:flex-none\"><h1 class=\"text-base font-semibold leading-6 text-gray-900\"><time datetime=\"2022-01\">January 2022</time></h1><div class=\"flex items-center\"><div class=\"relative flex items-center rounded-md bg-white shadow-xs md:items-stretch\"><button type=\"button\" class=\"flex h-9 w-12 items-center justify-center rounded-l-md border-y border-l border-gray-300 pr-1 text-gray-400 hover:text-gray-500 focus:relative md:w-9 md:pr-0 md:hover:bg-gray-50\"><span class=\"sr-only\">Previous month</span> <svg class=\"h-5 w-5\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z\" clip-rule=\"evenodd\"></path></svg></button> <button type=\"button\" class=\"hidden border-y border-gray-300 px-3.5 text-sm font-semibold text-gray-900 hover:bg-gray-50 focus:relative md:block\">Today</button> <span class=\"relative -mx-px h-5 w-px bg-gray-300 md:hidden\"></span> <button type=\"button\" class=\"flex h-9 w-12 items-center justify-center rounded-r-md border-y border-r border-gray-300 pl-1 text-gray-400 hover:text-gray-500 focus:relative md:w-9 md:pl-0 md:hover:bg-gray-50\"><span class=\"sr-only\">Next month</span> <svg class=\"h-5 w-5\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z\" clip-rule=\"evenodd\"></path></svg></button></div><div class=\"hidden md:ml-4 md:flex md:items-center\"><div class=\"relative\"><button type=\"button\" class=\"flex items-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-xs ring-1 ring-inset ring-gray-300 hover:bg-gray-50\" id=\"menu-button\" aria-expanded=\"false\" aria-haspopup=\"true\">Month view <svg class=\"-mr-1 h-5 w-5 text-gray-400\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M5.22 8.22a.75.75 0 0 1 1.06 0L10 11.94l3.72-3.72a.75.75 0 1 1 1.06 1.06l-4.25 4.25a.75.75 0 0 1-1.06 0L5.22 9.28a.75.75 0 0 1 0-1.06Z\" clip-rule=\"evenodd\"></path></svg></button><!--\n            Dropdown menu, show/hide based on menu state.\n\n            Entering: \"transition ease-out duration-100\"\n              From: \"transform opacity-0 scale-95\"\n              To: \"transform opacity-100 scale-100\"\n            Leaving: \"transition ease-in duration-75\"\n              From: \"transform opacity-100 scale-100\"\n              To: \"transform opacity-0 scale-95\"\n          --><div class=\"absolute right-0 z-10 mt-3 w-36 origin-top-right overflow-hidden rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-hidden\" role=\"menu\" aria-orientation=\"vertical\" aria-labelledby=\"menu-button\" tabindex=\"-1\"><div class=\"py-1\" role=\"none\"><!-- Active: \"bg-gray-100 text-gray-900\", Not Active: \"text-gray-700\" --><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-item-0\">Day view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-item-1\">Week view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-item-2\">Month view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-item-3\">Year view</a></div></div></div><div class=\"ml-6 h-6 w-px bg-gray-300\"></div><button type=\"button\" class=\"ml-6 rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500\">Add event</button></div><div class=\"relative ml-6 md:hidden\"><button type=\"button\" class=\"-mx-2 flex items-center rounded-full border border-transparent p-2 text-gray-400 hover:text-gray-500\" id=\"menu-0-button\" aria-expanded=\"false\" aria-haspopup=\"true\"><span class=\"sr-only\">Open menu</span> <svg class=\"h-5 w-5\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path d=\"M3 10a1.5 1.5 0 1 1 3 0 1.5 1.5 0 0 1-3 0ZM8.5 10a1.5 1.5 0 1 1 3 0 1.5 1.5 0 0 1-3 0ZM15.5 8.5a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3Z\"></path></svg></button><!--\n          Dropdown menu, show/hide based on menu state.\n\n          Entering: \"transition ease-out duration-100\"\n            From: \"transform opacity-0 scale-95\"\n            To: \"transform opacity-100 scale-100\"\n          Leaving: \"transition ease-in duration-75\"\n            From: \"transform opacity-100 scale-100\"\n            To: \"transform opacity-0 scale-95\"\n        --><div class=\"absolute right-0 z-10 mt-3 w-36 origin-top-right divide-y divide-gray-100 overflow-hidden rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-hidden\" role=\"menu\" aria-orientation=\"vertical\" aria-labelledby=\"menu-0-button\" tabindex=\"-1\"><div class=\"py-1\" role=\"none\"><!-- Active: \"bg-gray-100 text-gray-900\", Not Active: \"text-gray-700\" --><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-0\">Create event</a></div><div class=\"py-1\" role=\"none\"><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-1\">Go to today</a></div><div class=\"py-1\" role=\"none\"><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-2\">Day view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-3\">Week view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-4\">Month view</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700\" role=\"menuitem\" tabindex=\"-1\" id=\"menu-0-item-5\">Year view</a></div></div></div></div></header><div class=\"shadow-sm ring-1 ring-black ring-opacity-5 lg:flex lg:flex-auto lg:flex-col\"><div class=\"grid grid-cols-7 gap-px border-b border-gray-300 bg-gray-200 text-center text-xs font-semibold leading-6 text-gray-700 lg:flex-none\"><div class=\"flex justify-center bg-white py-2\"><span>M</span> <span class=\"sr-only sm:not-sr-only\">on</span></div><div class=\"flex justify-center bg-white py-2\"><span>T</span> <span class=\"sr-only sm:not-sr-only\">ue</span></div><div class=\"flex justify-center bg-white py-2\"><span>W</span> <span class=\"sr-only sm:not-sr-only\">ed</span></div><div class=\"flex justify-center bg-white py-2\"><span>T</span> <span class=\"sr-only sm:not-sr-only\">hu</span></div><div class=\"flex justify-center bg-white py-2\"><span>F</span> <span class=\"sr-only sm:not-sr-only\">ri</span></div><div class=\"flex justify-center bg-white py-2\"><span>S</span> <span class=\"sr-only sm:not-sr-only\">at</span></div><div class=\"flex justify-center bg-white py-2\"><span>S</span> <span class=\"sr-only sm:not-sr-only\">un</span></div></div><div class=\"flex bg-gray-200 text-xs leading-6 text-gray-700 lg:flex-auto\"><div class=\"hidden w-full lg:grid lg:grid-cols-7 lg:grid-rows-6 lg:gap-px\"><!--\n          Always include: \"relative py-2 px-3\"\n          Is current month, include: \"bg-white\"\n          Is not current month, include: \"bg-gray-50 text-gray-500\"\n        --><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><!--\n            Is today, include: \"flex h-6 w-6 items-center justify-center rounded-full bg-indigo-600 font-semibold text-white\"\n          --><time datetime=\"2021-12-27\">27</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2021-12-28\">28</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2021-12-29\">29</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2021-12-30\">30</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2021-12-31\">31</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-01\">1</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-01\">2</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-03\">3</time><ol class=\"mt-2\"><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Design review</p><time datetime=\"2022-01-03T10:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">10AM</time></a></li><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Sales meeting</p><time datetime=\"2022-01-03T14:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">2PM</time></a></li></ol></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-04\">4</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-05\">5</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-06\">6</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-07\">7</time><ol class=\"mt-2\"><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Date night</p><time datetime=\"2022-01-08T18:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">6PM</time></a></li></ol></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-08\">8</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-09\">9</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-10\">10</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-11\">11</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-12\" class=\"flex h-6 w-6 items-center justify-center rounded-full bg-indigo-600 font-semibold text-white\">12</time><ol class=\"mt-2\"><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Sam's birthday party</p><time datetime=\"2022-01-25T14:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">2PM</time></a></li></ol></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-13\">13</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-14\">14</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-15\">15</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-16\">16</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-17\">17</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-18\">18</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-19\">19</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-20\">20</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-21\">21</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-22\">22</time><ol class=\"mt-2\"><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Maple syrup museum</p><time datetime=\"2022-01-22T15:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">3PM</time></a></li><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Hockey game</p><time datetime=\"2022-01-22T19:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">7PM</time></a></li></ol></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-23\">23</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-24\">24</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-25\">25</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-26\">26</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-27\">27</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-28\">28</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-29\">29</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-30\">30</time></div><div class=\"relative bg-white px-3 py-2\"><time datetime=\"2022-01-31\">31</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-01\">1</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-02\">2</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-03\">3</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-04\">4</time><ol class=\"mt-2\"><li><a href=\"#\" class=\"group flex\"><p class=\"flex-auto truncate font-medium text-gray-900 group-hover:text-indigo-600\">Cinema with friends</p><time datetime=\"2022-02-04T21:00\" class=\"ml-3 hidden flex-none text-gray-500 group-hover:text-indigo-600 xl:block\">9PM</time></a></li></ol></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-05\">5</time></div><div class=\"relative bg-gray-50 px-3 py-2 text-gray-500\"><time datetime=\"2022-02-06\">6</time></div></div><div class=\"isolate grid w-full grid-cols-7 grid-rows-6 gap-px lg:hidden\"><!--\n          Always include: \"flex h-14 flex-col py-2 px-3 hover:bg-gray-100 focus:z-10\"\n          Is current month, include: \"bg-white\"\n          Is not current month, include: \"bg-gray-50\"\n          Is selected or is today, include: \"font-semibold\"\n          Is selected, include: \"text-white\"\n          Is not selected and is today, include: \"text-indigo-600\"\n          Is not selected and is current month, and is not today, include: \"text-gray-900\"\n          Is not selected, is not current month, and is not today: \"text-gray-500\"\n        --><button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><!--\n            Always include: \"ml-auto\"\n            Is selected, include: \"flex h-6 w-6 items-center justify-center rounded-full\"\n            Is selected and is today, include: \"bg-indigo-600\"\n            Is selected and is not today, include: \"bg-gray-900\"\n          --><time datetime=\"2021-12-27\" class=\"ml-auto\">27</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2021-12-28\" class=\"ml-auto\">28</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2021-12-29\" class=\"ml-auto\">29</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2021-12-30\" class=\"ml-auto\">30</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2021-12-31\" class=\"ml-auto\">31</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-01\" class=\"ml-auto\">1</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-02\" class=\"ml-auto\">2</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-03\" class=\"ml-auto\">3</time> <span class=\"sr-only\">2 events</span> <span class=\"-mx-0.5 mt-auto flex flex-wrap-reverse\"><span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span> <span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span></span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-04\" class=\"ml-auto\">4</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-05\" class=\"ml-auto\">5</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-06\" class=\"ml-auto\">6</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-07\" class=\"ml-auto\">7</time> <span class=\"sr-only\">1 event</span> <span class=\"-mx-0.5 mt-auto flex flex-wrap-reverse\"><span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span></span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-08\" class=\"ml-auto\">8</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-09\" class=\"ml-auto\">9</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-10\" class=\"ml-auto\">10</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-11\" class=\"ml-auto\">11</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 font-semibold text-indigo-600 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-12\" class=\"ml-auto\">12</time> <span class=\"sr-only\">1 event</span> <span class=\"-mx-0.5 mt-auto flex flex-wrap-reverse\"><span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span></span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-13\" class=\"ml-auto\">13</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-14\" class=\"ml-auto\">14</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-15\" class=\"ml-auto\">15</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-16\" class=\"ml-auto\">16</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-17\" class=\"ml-auto\">17</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-18\" class=\"ml-auto\">18</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-19\" class=\"ml-auto\">19</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-20\" class=\"ml-auto\">20</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-21\" class=\"ml-auto\">21</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 font-semibold text-white hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-22\" class=\"ml-auto flex h-6 w-6 items-center justify-center rounded-full bg-gray-900\">22</time> <span class=\"sr-only\">2 events</span> <span class=\"-mx-0.5 mt-auto flex flex-wrap-reverse\"><span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span> <span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span></span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-23\" class=\"ml-auto\">23</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-24\" class=\"ml-auto\">24</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-25\" class=\"ml-auto\">25</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-26\" class=\"ml-auto\">26</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-27\" class=\"ml-auto\">27</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-28\" class=\"ml-auto\">28</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-29\" class=\"ml-auto\">29</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-30\" class=\"ml-auto\">30</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-white px-3 py-2 text-gray-900 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-01-31\" class=\"ml-auto\">31</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-01\" class=\"ml-auto\">1</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-02\" class=\"ml-auto\">2</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-03\" class=\"ml-auto\">3</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-04\" class=\"ml-auto\">4</time> <span class=\"sr-only\">1 event</span> <span class=\"-mx-0.5 mt-auto flex flex-wrap-reverse\"><span class=\"mx-0.5 mb-1 h-1.5 w-1.5 rounded-full bg-gray-400\"></span></span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-05\" class=\"ml-auto\">5</time> <span class=\"sr-only\">0 events</span></button> <button type=\"button\" class=\"flex h-14 flex-col bg-gray-50 px-3 py-2 text-gray-500 hover:bg-gray-100 focus:z-10\"><time datetime=\"2022-02-06\" class=\"ml-auto\">6</time> <span class=\"sr-only\">0 events</span></button></div></div></div><div class=\"px-4 py-10 sm:px-6 lg:hidden\"><ol class=\"divide-y divide-gray-100 overflow-hidden rounded-lg bg-white text-sm shadow-sm ring-1 ring-black ring-opacity-5\"><li class=\"group flex p-4 pr-6 focus-within:bg-gray-50 hover:bg-gray-50\"><div class=\"flex-auto\"><p class=\"font-semibold text-gray-900\">Maple syrup museum</p><time datetime=\"2022-01-15T09:00\" class=\"mt-2 flex items-center text-gray-700\"><svg class=\"mr-2 h-5 w-5 text-gray-400\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16Zm.75-13a.75.75 0 0 0-1.5 0v5c0 .414.336.75.75.75h4a.75.75 0 0 0 0-1.5h-3.25V5Z\" clip-rule=\"evenodd\"></path></svg> 3PM</time></div><a href=\"#\" class=\"ml-6 flex-none self-center rounded-md bg-white px-3 py-2 font-semibold text-gray-900 opacity-0 shadow-xs ring-1 ring-inset ring-gray-300 hover:ring-gray-400 focus:opacity-100 group-hover:opacity-100\">Edit<span class=\"sr-only\">, Maple syrup museum</span></a></li><li class=\"group flex p-4 pr-6 focus-within:bg-gray-50 hover:bg-gray-50\"><div class=\"flex-auto\"><p class=\"font-semibold text-gray-900\">Hockey game</p><time datetime=\"2022-01-22T19:00\" class=\"mt-2 flex items-center text-gray-700\"><svg class=\"mr-2 h-5 w-5 text-gray-400\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16Zm.75-13a.75.75 0 0 0-1.5 0v5c0 .414.336.75.75.75h4a.75.75 0 0 0 0-1.5h-3.25V5Z\" clip-rule=\"evenodd\"></path></svg> 7PM</time></div><a href=\"#\" class=\"ml-6 flex-none self-center rounded-md bg-white px-3 py-2 font-semibold text-gray-900 opacity-0 shadow-xs ring-1 ring-inset ring-gray-300 hover:ring-gray-400 focus:opacity-100 group-hover:opacity-100\">Edit<span class=\"sr-only\">, Hockey game</span></a></li></ol></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
