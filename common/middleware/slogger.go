package middleware

import (
	"log/slog"
	"os"
	//"github.com/gofiber/fiber/v2/middleware/logger"
)

//	var RequestLogConfig = logger.Config{
//		Format:   "${pid} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
//		TimeZone: "UTC",
//	}
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
