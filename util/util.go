package util

import (
	"bytes"
	"fmt"
	"log/slog"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
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
func TimeNow() time.Time {
	return time.Now().UTC()
}

func TimeThen() time.Time {
	return time.Date(
		1970, 11, 17, 20, 34, 58, 651387237, time.UTC)
}

func NullTime() time.Time {
	var t time.Time
	return t
}

func Ident() uuid.UUID {
	return uuid.New()
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

func SysRequirment(logger *slog.Logger) bool {
	defer TimeElapsed(time.Now(), "Checking your Go environment")
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 8 {
		logger.Error("Need Go version 1.22 or higher!")
		return false
	}
	logger.Debug("You are using " + runtime.Compiler + " ")
	logger.Debug("on a" + runtime.GOARCH + "machine")
	logger.Debug("Using Go version " + runtime.Version())
	logger.Debug("Number of CPUs:" + strconv.Itoa(runtime.NumCPU()))
	logger.Debug("Number of Goroutines:" + strconv.Itoa(runtime.NumGoroutine()))
	return true

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

	// Validate the phone number using a regular expression
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	_ = validate.RegisterValidation("tel", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		return re.MatchString(field)
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) string {
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
		fields[err.Field()] = err.Param()
	}

	return CreateKeyValuePairs(fields)
}

func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}


