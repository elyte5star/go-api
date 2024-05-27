package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Username string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Host     string `required:"true" split_words:"true"`
	Port     string `required:"true" split_words:"true"`
	Database string `required:"true" split_words:"true"`
}

func (dbConfig *DbConfig) URL() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?timeout=30s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	return dsn
}

func getConfig(dbConfig DbConfig) (*mysql.Config, error) {
	config, err := mysql.ParseDSN(dbConfig.URL())
	if err != nil {
		return nil, err
	}
	config.ParseTime = true

	return config, nil
}

func ConnectToDB(dbConfig DbConfig) (*sql.DB, error) {
	config, err := getConfig(dbConfig)
	if err != nil {
		return nil, err
	}
	// Get a database handle.
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db, nil
}
