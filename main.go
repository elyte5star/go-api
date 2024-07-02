package main

import (
	"fmt"

	"github.com/api/common/middleware"
	_ "github.com/api/docs"
	"github.com/api/service"
	"github.com/api/service/dbutils"
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

@accept json

@produce json

@schemes http https

@securityDefinitions.apikey BearerAuth
@in header
@name Authorization
@description Bearer Token

@externalDocs.description  elyte5star
@externalDocs.url          https://github.com/elyte5star/go-api
*/
func main() {

	// Load the config struct with values from the environment

	//Set up validation and attach to config
	validate := util.InitValidator()

	cfg, _ := service.ParseConfig(validate)

	// Set up the logger
	logger := middleware.DefaultLogger()

	//Set logging to DEBUG LEVEL in Development
	if cfg.Debug {
		middleware.DebugLogger()
	}

	cfg.Logger = logger

	cfg.Validate = validate

	h := Handler(cfg)

	if db, err := service.ConnectToMySQL(cfg); err == nil {
		dbutils.CreateTables(db)
		dbutils.CreateAdminAccount("elyte",cfg)
	}
	address := fmt.Sprintf(":%v", cfg.ServicePort)

	logger.Info("Listening on " + address)
	// start server
	h.Listen(address)

}
