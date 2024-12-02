package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/gommon/color"

	"github.com/bartalcorn/goweb/pkg/config"
	"github.com/bartalcorn/goweb/pkg/middleware"
)

func Start(router *http.ServeMux) {
	// The HTTP Server
	address := fmt.Sprintf("%s:%s", config.Configs.ADDRESS, config.Configs.PORT)
	server := &http.Server{Addr: address, Handler: middleware.Logger(router)}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancelCtx := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancelCtx()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger a graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	fmt.Printf("Available at: %v\n", color.BlueBg("http://"+address, color.Yel, color.U))
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func cancelServer() {
	fmt.Println("Server was cancelled")
}
