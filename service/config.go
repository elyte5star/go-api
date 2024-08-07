package service

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Debug                 bool
	ClientOrigin          string ` validate:"required"`
	SmtpServer            string ` validate:"required"`
	SmtpUsername          string ` validate:"required"`
	SmtpPort              int    ` validate:"required"`
	SmtpPassword          string ` validate:"required"`
	JwtSecretKey          string ` validate:"required"`
	JwtExpireMinutesCount int    ` validate:"required,gte=0"`
	ServiceName           string ` validate:"required"`
	SupportEmail          string ` validate:"required,email"`
	ReadTimeout           int    ` validate:"required,gte=0"`
	Version               string ` validate:"required"`
	Logger                *slog.Logger
	Validate              *validator.Validate
	ServicePort           string ` validate:"required"`
	CorsOrigins           string ` validate:"required"`
	Url                   string ` validate:"required"`
	Doc                   string ` validate:"required"`
	DbConfig              *DbConfig
}

type DbConfig struct {
	User     string `envconfig:"MYSQL_USER" split_words:"true" validate:"required"`
	Password string `envconfig:"MYSQL_PASSWORD" split_words:"true" validate:"required"`
	Host     string `envconfig:"MYSQL_HOST" split_words:"true" validate:"required"`
	Port     string `envconfig:"MYSQL_PORT" split_words:"true" validate:"required"`
	Database string `envconfig:"MYSQL_DATABASE" validate:"required"`
}

func (dbConfig *DbConfig) URL() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	return dsn
}

func ParseConfig(val *validator.Validate) (*AppConfig, error) {
	var config AppConfig
	// Load the environment vars from a .env file if present
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldnt load env files %p", err)
	}
	if err := envconfig.Process("myapp", &config); err != nil {
		log.Fatalf("Couldnt parse env variables to config struct %p", err)
	}
	if err := val.Struct(config); err != nil {
		log.Fatal(err)
	}
	return &config, nil

}
