package main

import (
	"fmt"
	"log/slog"
	"os"
	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
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


	// Output the config for debugging
	fmt.Printf("%+v\n", conf.DbConfig)

	swaggerDocHandler := middleware.SwaggerHandler(conf)

	app.Get("/docs/*", swaggerDocHandler)


	// Routes
	router.RouteSetup(app)

	
	address := fmt.Sprintf(":%v", conf.ServicePort)

	log.Info("Listening on " + address)
	// start server
	app.Listen(address)

}
