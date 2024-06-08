package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	"github.com/api/util"
)

func main() {
	validate := util.InitValidator()
	// Load the config struct with values from the environment
	conf, _ := config.ParseConfig(validate)

	// Set up the logger
	logger := middleware.DefaultLogger()
	if conf.Debug {
		middleware.DebugLogger()
	}
	conf.Logger = logger

	// Output the config for debugging
	//fmt.Printf("%+v\n", conf)

	bootstrap := Handler(conf)
	address := fmt.Sprintf(":%v", conf.ServicePort)

	logger.Info("Listening on " + address)
	// start server
	bootstrap.Listen(address)

}
