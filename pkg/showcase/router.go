package showcase

import (
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handle.base)
	mux.HandleFunc("GET /{element}", handle.element)
	return mux
}
