package config

import (
	"log"
	"log/slog"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type  AppConfig struct {
	Debug      bool `required:"true"`
	ServiceName string `required:"true"`
	Version     string `required:"true"`
	Logger *slog.Logger `ignored:"true"`
	Port string `required:"true"`
	CorsOrigins string `required:"true"`
	Url string `required:"true"`
	Doc string `required:"true"`
}

func Config() AppConfig{

	var s  AppConfig

	// Load the environment vars from a .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("myapp", &s)
    if err != nil {
		log.Fatal(err)
    }
	
	return s

}



