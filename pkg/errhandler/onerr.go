package errhandler

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/gommon/color"
)

func OnErr(message string, err error) {
	if err != nil {
		fmt.Printf("%v %v %v\n", time.Now().Format("2006/01/02 15:04:05"), color.Red("ERROR "+message, color.B), color.Cyan(err))

		_, file, no, ok := runtime.Caller(1)
		fp := strings.Split(file, "/")
		if ok {
			fmt.Printf("%v %v", fmt.Sprintf(color.Yellow("at line =>")), fmt.Sprintf(color.Yellow(fmt.Sprintf("%s:%d \n", strings.Join(fp[len(fp)-4:], "/"), no-1), color.B)))
		}
	}
}
