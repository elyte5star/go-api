package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
)

func main() {

	// Load the config struct with values from the environment
	conf := config.Config()

	// Set up the logger
	logger := middleware.DefaultLogger()
	if conf.Debug {
		logger = middleware.DebugLogger()
	}
	conf.Logger = logger

	// Output the config for debugging
	//fmt.Printf("%+v\n", conf.DbConfig)

	bootstrap := Handler(&conf)
	address := fmt.Sprintf(":%v", conf.ServicePort)

	logger.Info("Listening on " + address)
	// start server
	bootstrap.Listen(address)

}
