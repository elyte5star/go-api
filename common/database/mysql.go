package database

import (
	"database/sql"
	"time"

	"github.com/api/common/config"
	"github.com/go-sql-driver/mysql"
)

func getConfig(dbConfig config.DbConfig) (*mysql.Config, error) {
	config, err := mysql.ParseDSN(dbConfig.URL())
	if err != nil {
		return nil, err
	}
	config.ParseTime = true

	return config, nil
}

func ConnectToDB(cfg config.AppConfig) (*sql.DB, error) {
	config, err := getConfig(*cfg.DbConfig)
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
