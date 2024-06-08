package main

import (
	"fmt"

	"github.com/api/common/config"
	"github.com/api/common/middleware"
	_ "github.com/api/docs"
	"github.com/api/util"
)

// @title Elyte Realm API
// @version 1.0.1
// @description Interactive Documentation for Elyte e-Market
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name Elyte Fiber Application.
// @contact.url https://github.com/elyte5star.
// @accept json
// @contact.email elyte5star@gmail.com
// @license.name Proprietary
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @externalDocs.url https://github.com/elyte5star/go-api
// @BasePath /
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
