package dbutils

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

const users = `CREATE TABLE IF NOT EXISTS users (
	userid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	username VARCHAR(64) NOT NULL UNIQUE,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(64) NOT NULL UNIQUE,
	accountNonLocked BOOLEAN DEFAULT false,
	admin BOOLEAN DEFAULT false,
	enabled BOOLEAN DEFAULT false,
	telephone VARCHAR(64) NOT NULL UNIQUE,
	discount   DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	failedAttempt INT UNSIGNED  DEFAULT '0000' NOT NULL,
	LockTime TIMESTAMP(0),
	auditInfo JSON NOT NULL
	) 
	`
const otp = `CREATE TABLE IF NOT EXISTS otp (
		userid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
		email VARCHAR(64) NOT NULL UNIQUE,
		otpString VARCHAR(64) NOT NULL,
		expiryDate TIMESTAMP(0)
		FOREIGN KEY (USER_ID) REFERENCES users(userid),
		)
`

const userAddress = `CREATE TABLE IF NOT EXISTS user_address (
	userid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	fullName VARCHAR(64) NOT NULL UNIQUE,
	streetAddress VARCHAR(64) NOT NULL,
	country VARCHAR(64) NOT NULL,
	state VARCHAR(64) NOT NULL,
	zip VARCHAR(64) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid),
	)
`

const userLocations = `CREATE TABLE IF NOT EXISTS user_locations (
	locationId CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	country VARCHAR(64) NOT NULL UNIQUE,
	enabled BOOLEAN DEFAULT false,
	userid CHAR(36) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid),
	)
`

const bookings = `CREATE TABLE IF NOT EXISTS bookings (
	oId CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	createdAt TIMESTAMP(0),
	userid CHAR(36) NOT NULL,
	cart JSON NOT NULL,
	shippingDetails JSON NOT NULL,
	totalPrice DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid),
	)
`
const products = `CREATE TABLE IF NOT EXISTS products (
	pid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	description VARCHAR(264) NOT NULL,
	category VARCHAR(264) NOT NULL,
	price DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	stockQuantity INT UNSIGNED  DEFAULT '0000' NOT NULL,
	image VARCHAR(64) NOT NULL,
	details VARCHAR(664) NOT NULL,
	productDiscount DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	auditInfo JSON NOT NULL,
	)
`
const productReview = `CREATE TABLE IF NOT EXISTS reviews (
	rid CHAR(36) DEFAULT (UUID_TO_BIN(UUID())) PRIMARY KEY,
	createdAt TIMESTAMP(0),
	rating INT UNSIGNED  DEFAULT '0000' NOT NULL,
	reviewerName VARCHAR(264) NOT NULL,
	comment VARCHAR(500) NOT NULL,
	email VARCHAR(64) NOT NULL,
	pid CHAR(36) NOT NULL,
	FOREIGN KEY (pid) REFERENCES products(pid),
		)
	`

func CreateTables(logger *slog.Logger, dbDriver *sqlx.DB) {
	statement, driverError := dbDriver.Prepare(users)
	if driverError != nil {
		logger.Warn(driverError.Error())

	}
	// Create table
	_, statementError := statement.Exec()
	if statementError != nil {
		logger.Warn("Table already exists!")
	}
	statement, _ = dbDriver.Prepare(otp)
	statement.Exec()
	statement, _ = dbDriver.Prepare(userAddress)
	statement.Exec()
	statement, _ = dbDriver.Prepare(userLocations)
	statement.Exec()
	statement, _ = dbDriver.Prepare(bookings)
	statement.Exec()
	statement, _ = dbDriver.Prepare(products)
	statement.Exec()
	statement, _ = dbDriver.Prepare(productReview)
	statement.Exec()
	logger.Info("All tables created/initialized successfully!")
}
