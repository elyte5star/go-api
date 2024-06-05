package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/api/common/config"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// TimeElapsed measures the time it takes to execute a function.
// Use it as like this with defer:
//
//	defer TimeElapsed(time.Now(), "FunctionToTime")
//
// For more details see: https://coderwall.com/p/cp5fya
func TimeElapsed(start time.Time, name string) string {
	elapsed := time.Since(start)
	fmt.Println(name + " took " + elapsed.String())
	return elapsed.String()
}
func TimeNow() string {
	return time.Now().UTC().String()
}

// func ConnectionString() string {
// 	connStr, status := os.LookupEnv("CONN_STR")
// 	if !status {
// 		Logger().Error("Missing environment variable CONN_STR")
// 	}

// 	return connStr
// }

func SysRequirment(cfg *config.AppConfig) bool {
	defer TimeElapsed(time.Now(), "Checking System Information and Requirements")
	logger := cfg.Logger
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 8 {
		logger.Error("Need Go version 1.22 or higher!")
		return false
	}
	// logger.Info("You are using " + runtime.Compiler + " ")
	// logger.Info("on a" + runtime.GOARCH + "machine")
	// logger.Info("Using Go version " + runtime.Version())
	// logger.Info("Number of CPUs:" + strconv.Itoa(runtime.NumCPU()))
	// logger.Info("Number of Goroutines:" + strconv.Itoa(runtime.NumGoroutine()))
	return true

}

const JwtSecret = "7a3c54660456ff1137b652e498624dfa09a0ec12b4fc49d38b85465da15027a1"

var RequestLogConfig = logger.Config{
	Format:   "${pid} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
	TimeZone: "UTC",
}
