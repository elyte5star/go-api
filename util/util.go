package util

import (
	"time"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// TimeElapsed measures the time it takes to execute a function.
// Use it as like this with defer:
//     defer TimeElapsed(time.Now(), "FunctionToTime")
//
// For more details see: https://coderwall.com/p/cp5fya
func TimeElapsed(start time.Time, name string) string{
	elapsed := time.Since(start)
	Logger().Info(name +" took " + elapsed.String())
	return elapsed.String()
}


func ConnectionString() string {
    connStr, status := os.LookupEnv("CONN_STR")
    if !status {
		Logger().Error("Missing environment variable CONN_STR")
    }

    return connStr
}


func Logger () *slog.Logger{
	logHandler:= slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	logger:= slog.New(logHandler)
	return logger
}


func SystemInfo() {
	defer TimeElapsed(time.Now(), "System Information")
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 8 {
		Logger().Info("Need Go version 1.22 or higher!")
		return
	}
	Logger().Info("You are using " + runtime.Compiler + " ")
	Logger().Info("on a" + runtime.GOARCH + "machine")
	Logger().Info("Using Go version " + runtime.Version())
	Logger().Info("Number of CPUs:" + strconv.Itoa(runtime.NumCPU()))
	Logger().Info("Number of Goroutines:" + strconv.Itoa(runtime.NumGoroutine()))
	
}

const JwtSecret = "7a3c54660456ff1137b652e498624dfa09a0ec12b4fc49d38b85465da15027a1"


var RequestLogConfig = logger.Config{
    Format:        "${pid} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
    TimeZone:      "UTC",
   
}