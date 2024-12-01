package tws

/**
tws - Tailwind Strings

This is a helper / convenience package to standardize the otherwise VERY verbose TailwindCSS directives.
*/

const (
	FlexRow   string = "flex flex-row gap-2"
	FlexRowCB string = FlexRow + " items-center justify-between"
	NoRing    string = "focus-visible:outline-none focus-visible:ring-0 focus-visible:ring-offset-2 focus:outline-none focus:ring-0 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed"
)
