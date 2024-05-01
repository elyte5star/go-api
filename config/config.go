package config

import (
	"log/slog"
	"os"
	"github.com/joho/godotenv"
)

var logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})

func Config(key string) string {
	// load .env file
	logger := slog.New(logHandler)
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")

	}
	return os.Getenv(key)

}
