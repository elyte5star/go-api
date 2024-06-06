package config

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/go-playground/validator.v9"
)

type AppConfig struct {
	Debug                 bool         `required:"true"`
	ClientOrigin          string       `required:"true"`
	SmtpServer            string       `required:"true"`
	SmtpUsername          string       `required:"true"`
	SmtpPassword          string       `required:"true"`
	JwtSecretKey          string       `required:"true"`
	JwtExpireMinutesCount int          `required:"true"`
	ServiceName           string       `required:"true"`
	Version               string       `required:"true"`
	Logger                *slog.Logger `ignored:"true"`
	ServicePort           string       `required:"true"`
	CorsOrigins           string       `required:"true"`
	Url                   string       `required:"true"`
	Doc                   string       `required:"true"`
	DbConfig              *DbConfig
}

type DbConfig struct {
	User     string `envconfig:"MYSQL_USER" required:"true" split_words:"true"`
	Password string `envconfig:"MYSQL_PASSWORD" required:"true" split_words:"true"`
	Host     string `envconfig:"MYSQL_HOST" required:"true" split_words:"true"`
	Port     string `envconfig:"MYSQL_PORT" required:"true" split_words:"true"`
	Database string `envconfig:"MYSQL_DATABASE" required:"true" split_words:"true"`
}

func (dbConfig *DbConfig) URL() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?timeout=30s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	return dsn
}

func ParseConfig(val *validator.Validate) (config *AppConfig, err error) {

	log.Println("Parsing ENV vars...")
	defer log.Println("Done Parsing ENV vars")
	config = &AppConfig{}
	// Load the environment vars from a .env file if present
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err = envconfig.Process("myapp", config); err != nil {
		log.Fatal(err)
	}
	//err = val.Struct(config)

	return

}
