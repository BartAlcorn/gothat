package main

import (
	srvr "github.com/bartalcorn/goweb/pkg/server"

	"github.com/bartalcorn/gothat/pkg/server"
)

func main() {
	srvr.Start(server.Router())
}
