package server

import (
	"net/http"

	"github.com/bartalcorn/weather"

	"github.com/bartalcorn/gothat/pkg/config"
	"github.com/bartalcorn/gothat/pkg/embedded"
	"github.com/bartalcorn/gothat/pkg/home"
	"github.com/bartalcorn/gothat/pkg/showcase"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", home.Router())
	mux.Handle("/home/", http.StripPrefix("/home", home.Router()))
	mux.Handle("/showcase/", http.StripPrefix("/showcase", showcase.Router()))
	mux.Handle("/weather/", http.StripPrefix("/weather", weather.Router()))

	if config.Configs.ENV == "local" {
		mux.Handle("GET /assets/{path...}", http.StripPrefix("/assets/", http.FileServer(http.Dir("./pkg/embedded/assets"))))
	} else {
		mux.Handle("GET /assets/path...}", http.StripPrefix("/assets/", http.FileServer(http.FS(embedded.Assets))))
	}

	return mux
}
