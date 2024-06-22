package api

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/api/service"
	"github.com/gofiber/fiber/v2"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App, cfg *service.AppConfig) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})
	log := cfg.Logger
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Warn(fmt.Sprintf("Oops... Server is not shutting down! Reason: %v", err))
		}

		close(idleConnsClosed)
	}()
	address := fmt.Sprintf(":%v", cfg.ServicePort)
	// Run server.
	if err := a.Listen(address); err != nil {
		log.Warn(fmt.Sprintf("Oops... Server is not running! Reason: %v", err))
	}

	<-idleConnsClosed
}

