package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
	"github.com/api/util"
)

/*
@title Elyte Realm API
@version 1.0.1
@description Interactive Documentation for Elyte-Realm API
@termsOfService http://swagger.io/terms/

@contact.name Elyte Fiber Application.
@contact.url https://github.com/elyte5star.
@contact.email elyte5star@gmail.com

@license.name Proprietary
@license.url http://www.apache.org/licenses/LICENSE-2.0.html

@host localhost:8080
@BasePath /

@securityDefinitions.apikey ApiKeyAuth
@in header
@name Authorization
@description Jwt Bearer Token

@accept json
*/
func main() {

	// Load the config struct with values from the environment
	conf, _ := config.ParseConfig()

	// Set up the logger
	logger := middleware.DefaultLogger()

	//Set logging to DEBUG LEVEL in Development
	if conf.Debug {
		middleware.DebugLogger()
	}

	conf.Logger = logger

	//Set up validation and attach to config
	validate := util.InitValidator()

	conf.Validate = validate

	// Output the config for debugging
	//fmt.Printf("%+v\n", conf)

	bootstrap := Handler(conf)

	address := fmt.Sprintf(":%v", conf.ServicePort)

	logger.Info("Listening on " + address)
	// start server
	bootstrap.Listen(address)

}
