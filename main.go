package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
)

func main() {

	// Load the config struct with values from the environment
	conf := config.ParseConfig()

	// Set up the logger
	logger := middleware.DefaultLogger()
	if conf.Debug {
		middleware.DebugLogger()
	}
	conf.Logger = logger

	bootstrap := Handler(&conf)
	address := fmt.Sprintf(":%v", conf.ServicePort)

	logger.Info("Listening on " + address)
	// start server
	bootstrap.Listen(address)

}
