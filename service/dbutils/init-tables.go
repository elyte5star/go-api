package dbutils

import (
	"log"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

const users = `CREATE TABLE IF NOT EXISTS users (
	userid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	username VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(64) NOT NULL,
	accountNonLocked BOOLEAN DEFAULT false,
	admin BOOLEAN DEFAULT false,
	enabled BOOLEAN DEFAULT false,
	telephone VARCHAR(64) NOT NULL,
	discount   DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	failedAttempt INT UNSIGNED  DEFAULT '0000' NOT NULL,
	LockTime TIMESTAMP(0),
	auditInfo JSON
	) ENGINE=INNODB
	`

func CreateTables(logger *slog.Logger, dbDriver *sqlx.DB) {
	statement, driverError := dbDriver.Prepare(users)
	if driverError != nil {
		log.Fatal(driverError.Error())

	}
	// Create table
	_, statementError := statement.Exec()
	if statementError != nil {
		
		logger.Warn("Table already exists!")
	}
	// statement, _ = dbDriver.Prepare(station)
	// statement.Exec()
	// statement, _ = dbDriver.Prepare(schedule)
	// statement.Exec()
	logger.Info("All tables created/initialized successfully!")
}
