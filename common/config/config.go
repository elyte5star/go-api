package config

import (
	"log"
	"log/slog"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Debug       bool         `required:"true"`
	ServiceName string       `required:"true"`
	Version     string       `required:"true"`
	Logger      *slog.Logger `ignored:"true"`
	ServicePort        string       `required:"true"`
	CorsOrigins string       `required:"true"`
	Url         string       `required:"true"`
	Doc         string       `required:"true"`
	DbConfig  *DbSpec
	
}


type DbSpec struct {
	MysqlUser string `envconfig:"MYSQL_USER" required:"true" split_words:"true"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD" required:"true" split_words:"true"`
	MysqlHost     string `envconfig:"MYSQL_HOST" required:"true" split_words:"true"`
	MysqlPort     string `envconfig:"MYSQL_PORT" required:"true" split_words:"true"`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE" required:"true" split_words:"true"`
}

func Config() AppConfig {

	var config AppConfig

	// Load the environment vars from a .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("myapp", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config

}
