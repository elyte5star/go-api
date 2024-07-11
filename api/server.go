package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/api/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartApi(a *fiber.App, cfg *service.AppConfig, db *sqlx.DB) {
	// Create channel for  connections.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		// Received an interrupt signal, shutdown.
		cfg.Logger.Info("Gracefully shutting down...")
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			cfg.Logger.Warn(fmt.Sprintf("Oops... Server is not shutting down! Reason: %v", err))
		}
		close(c)
	}()

	// ...
	address := fmt.Sprintf(":%v", cfg.ServicePort)
	if err := a.Listen(address); err != nil {
		log.Panic(err)
	}
	cfg.Logger.Warn("Running cleanup tasks...")
	
	// Your cleanup tasks go here
	db.Close()
}
