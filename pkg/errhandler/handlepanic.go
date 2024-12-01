package errhandler

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

func HandlePanic() {
	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println(color.Red("RECOVERING => "), a)
	}
}
