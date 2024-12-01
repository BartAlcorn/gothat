package alpinejs

import (
	"fmt"

	"github.com/a-h/templ"
)

type Direction string
type Duration string

type TranslateProps struct {
	Direction   Direction `default:"right"`
	DurationIn  Duration  `default:"duration-300"`
	DurationOut Duration  `default:"duration-300"`
}

const (
	Top    Direction = "top"
	Right  Direction = "right"
	Bottom Direction = "bottom"
	Left   Direction = "left"

	Duration0    Duration = "duration-0"    // transition-duration: 0s;
	Duration75   Duration = "duration-75"   // transition-duration: 75ms;
	Duration100  Duration = "duration-100"  // transition-duration: 100ms;
	Duration150  Duration = "duration-150"  // transition-duration: 150ms;
	Duration200  Duration = "duration-200"  // transition-duration: 200ms;
	Duration300  Duration = "duration-300"  // transition-duration: 300ms;
	Duration500  Duration = "duration-500"  // transition-duration: 500ms;
	Duration700  Duration = "duration-700"  // transition-duration: 700ms;
	Duration1000 Duration = "duration-1000" // transition-duration: 1000ms;
)

func Translate(props *TranslateProps) templ.Attributes {
	if props.Direction == "" {
		props.Direction = "right"
	}
	if props.DurationIn == "" {
		props.DurationIn = "duration-300"
	}
	if props.DurationOut == "" {
		props.DurationOut = "duration-300"
	}

	var dir templ.Attributes
	switch props.Direction {
	case Top:
		dir = TranslateTop
	case Right:
		dir = TranslateRight
	case Bottom:
		dir = TranslateBottom
	case Left:
		dir = TranslateLeft
	}
	if props.DurationIn != "" {
		dir["x-transition:enter"] = fmt.Sprintf("transform transition ease-in %s", props.DurationIn)
	}
	if props.DurationOut != "" {
		dir["x-transition:leave"] = fmt.Sprintf("transform transition ease-out %s", props.DurationOut)
	}
	return dir
}

var TranslateLeft = templ.Attributes{
	"x-transition:enter":       "transform transition ease-in duration-300",
	"x-transition:enter-start": "-translate-x-full",
	"x-transition:enter-end":   "translate-x-0",
	"x-transition:leave":       "transform transition ease-out duration-300",
	"x-transition:leave-start": "translate-x-0",
	"x-transition:leave-end":   "left-0 -translate-x-full",
}
var TranslateRight = templ.Attributes{
	"x-transition:enter":       "transform transition ease-in duration-300",
	"x-transition:enter-start": "translate-x-full",
	"x-transition:enter-end":   "translate-x-0",
	"x-transition:leave":       "transform transition ease-out duration-300",
	"x-transition:leave-start": "translate-x-0",
	"x-transition:leave-end":   "right-0 translate-x-full",
}
var TranslateTop = templ.Attributes{
	"x-transition:enter":       "transform transition ease-in duration-300",
	"x-transition:enter-start": "-translate-y-full",
	"x-transition:enter-end":   "transform translate-y-0",
	"x-transition:leave":       "transform transition ease-out duration-300",
	"x-transition:leave-start": "translate-y-0",
	"x-transition:leave-end":   "top-0 -translate-y-full",
}
var TranslateBottom = templ.Attributes{
	"x-transition:enter":       "transform transition ease-in duration-300",
	"x-transition:enter-start": "translate-y-full",
	"x-transition:enter-end":   "translate-y-0",
	"x-transition:leave":       "transform transition ease-out duration-300",
	"x-transition:leave-start": "translate-y-0",
	"x-transition:leave-end":   "bottom-0 translate-y-full",
}
