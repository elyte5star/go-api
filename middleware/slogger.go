package middleware

import (
	"log/slog"
	"os"
	
)

var programLevel = new(slog.LevelVar) // Info by default

func DefaultLogger() *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: programLevel})
	//timeFormmatter := slogformatter.NewFormatterHandler(slogformatter.TimezoneConverter(time.UTC), slogformatter.TimeFormatter(time.DateTime, time.UTC))
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return logger
}

func DebugLogger() {
	programLevel.Set(slog.LevelDebug)
}
