package database

import (
	"fmt"
	"time"

	"github.com/api/common/config"
	"github.com/api/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*repository.UserQueries
}

func getDbConfig(dbConfig *config.DbConfig) (*mysql.Config, error) {
	config, err := mysql.ParseDSN(dbConfig.URL())
	if err != nil {
		return nil, err
	}
	config.ParseTime = true

	return config, nil
}

func ConnectToMySQL(cfg *config.AppConfig) (*sqlx.DB, error) {
	config, err := getDbConfig(cfg.DbConfig)
	if err != nil {
		return nil, fmt.Errorf("error, Getting Database configurations, %w", err)
	}
	// Get a database handle.
	db, err := sqlx.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, cant ping database, %w", err)
	}

	return db, nil
}

func DbWithQueries(cfg *config.AppConfig) (*Queries, error) {
	db, err := ConnectToMySQL(cfg)
	if err != nil {
		return nil, err
	}
	Initialize(db, cfg)
	return &Queries{
		UserQueries: &repository.UserQueries{DB: db},
	}, nil

}
