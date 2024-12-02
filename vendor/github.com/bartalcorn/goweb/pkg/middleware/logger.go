package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/gommon/color"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(color.Cyan(time.Since(start)), color.Yellow(r.Method), color.Green(r.URL.Path))
			if r.Header.Get("Hx-Request") != "" {
				log.Println(color.BlueBg("Hx-Request", r.Header.Get("HX-Request")))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
