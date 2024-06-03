package middleware

import (
	"os"
	"log/slog"
	"time"
	slogformatter "github.com/samber/slog-formatter"
)



func Logger() *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	timeFormmatter := slogformatter.NewFormatterHandler(slogformatter.TimezoneConverter(time.UTC), slogformatter.TimeFormatter(time.DateTime, time.UTC))
	logger := slog.New(timeFormmatter(logHandler))
	return logger
}