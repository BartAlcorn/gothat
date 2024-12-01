package formats

import (
	"time"
)

func Now() string {
	return time.Now().Format("Jan 02 2006 15:04")
}
