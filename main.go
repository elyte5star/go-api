package main

import (
	"fmt"
	"log/slog"
	"os"
	"github.com/api/common/config"
	router "github.com/api/router"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {

	//check if application meets requirments
	meetSysRequirment := util.SystemInfo()

	if !meetSysRequirment {
		os.Exit(500)
	}

	// Load the config struct with values from the environment
	conf := config.Config()

	// Set up the logger
	log := util.Logger()
	conf.Logger = log

	appInfo := fmt.Sprintf("%s:%s", conf.ServiceName, conf.Version)

	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName: appInfo,
	})

	if conf.Debug {
		config := slogfiber.Config{
			DefaultLevel:     slog.LevelDebug,
			ClientErrorLevel: slog.LevelWarn,
			ServerErrorLevel: slog.LevelError,
		}
		app.Use(slogfiber.NewWithConfig(log, config))

	} else {
		app.Use(slogfiber.New(log))

	}

	// Routes
	router.RouteSetup(app)
	
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	address := fmt.Sprintf(":%v", conf.Port)

	log.Info("Listening on " + address)
	// start server
	app.Listen(address)

}
