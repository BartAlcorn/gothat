package tws

/**
tws - Tailwind Strings

This is a helper / convenience package to standardize the otherwise VERY verbose TailwindCSS directives.
*/

const (
	Component string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-gray-600 bg-gray-100 px-3 py-2 text-base text-gray-800 shadow-xs ring-1 ring-inset ring-gray-300 hover:bg-gray-50 hover:text-blue-600 active:bg-blue-600/50 active:text-blue-600"
	Disabled  string = "disabled:text-neutral-500/50 disabled:border-gray-200/60 disabled:bg-gray-200/20 disabled:hover:border-gray-200/10 disabled:hover:bg-gray-200/10"
	FlexRow   string = "flex flex-row gap-2"
	FlexRowCB string = FlexRow + " items-center justify-between"
	Primary   string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-blue-600 bg-blue-100 px-3 py-2 text-base text-blue-800 shadow-xs ring-1 ring-inset ring-blue-300 hover:bg-blue-50 hover:text-blue-600 active:bg-blue-600/50 active:text-blue-600"
	Secondary string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-teal-600 bg-teal-100 px-3 py-2 text-base text-teal-800 shadow-xs ring-1 ring-inset ring-teal-300 hover:bg-teal-50 hover:text-teal-600 active:bg-teal-600/50 active:text-teal-600"
	Danger    string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-rose-600 bg-rose-100 px-3 py-2 text-base text-rose-800 shadow-xs ring-1 ring-inset ring-rose-300 hover:bg-rose-50 hover:text-rose-600 active:bg-rose-600/50 active:text-rose-600"
	Warning   string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-amber-600 bg-amber-100 px-3 py-2 text-base text-amber-800 shadow-xs ring-1 ring-inset ring-amber-300 hover:bg-amber-50 hover:text-amber-600 active:bg-amber-600/50 active:text-amber-600"
	Info      string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-blue-600 bg-blue-100 px-3 py-2 text-base text-blue-800 shadow-xs ring-1 ring-inset ring-blue-300 hover:bg-blue-50 hover:text-blue-600 active:bg-blue-600/50 active:text-blue-600"
	Success   string = "inline-flex items-center justify-center gap-2 gap-x-1.5 rounded-md border-green-600 bg-green-100 px-3 py-2 text-base text-green-800 shadow-xs ring-1 ring-inset ring-green-300 hover:bg-green-50 hover:text-green-600 active:bg-green-600/50 active:text-green-600"
	NoRing    string = "focus-visible:outline-none focus-visible:ring-0 focus-visible:ring-offset-2 focus:outline-none focus:ring-0 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed"
)
