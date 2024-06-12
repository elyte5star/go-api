package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/api/common/config"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
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

type TimestampTime struct {
	time.Time
}

// /implement encoding.JSON.Marshaler interface
func (t *TimestampTime) MarshalJSON() ([]byte, error) {
	bin := make([]byte, 16)
	bin = strconv.AppendInt(bin[:0], t.Time.Unix(), 10)
	return bin, nil
}

func (t *TimestampTime) UnmarshalJSON(bin []byte) error {
	v, err := strconv.ParseInt(string(bin), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(v, 0)
	return nil
}

///
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

func InitValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
	}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Tag()
	}

	return fields
}
