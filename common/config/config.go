package config

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/api/util"
)


func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		util.Logger().Error("Error loading .env file")

	}
	return os.Getenv(key)

}


