package dbutils

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const schema = `

DROP TABLE IF EXISTS users,otp,address,user_locations,bookings;

CREATE TABLE IF NOT EXISTS users (
	userid CHAR(36) PRIMARY KEY,
	username VARCHAR(64) NOT NULL UNIQUE,
	password VARBINARY(255) NOT NULL,
	email VARCHAR(64) NOT NULL UNIQUE,
	accountNonLocked BOOLEAN DEFAULT false,
	admin BOOLEAN DEFAULT false,
	enabled BOOLEAN DEFAULT false ,
	isUsing2FA BOOLEAN DEFAULT false ,
	telephone VARCHAR(64) NOT NULL UNIQUE,
	discount  DECIMAL(16,2) DEFAULT '0.00',
	failedAttempt INT UNSIGNED  DEFAULT '0000',
	lockTime TIMESTAMP(0),
	auditInfo JSON NOT NULL
	);
	
	
CREATE TABLE IF NOT EXISTS otp (
		userid CHAR(36) PRIMARY KEY,
		email VARCHAR(64) NOT NULL UNIQUE,
		otpString VARCHAR(64) NOT NULL,
		expiryDate TIMESTAMP(0),
		FOREIGN KEY(userid) REFERENCES users(userid)
		);


CREATE TABLE IF NOT EXISTS address (
	addressId CHAR(36).
	ownerId CHAR(36),
	fullName VARCHAR(64) NOT NULL,
	streetAddress VARCHAR(64) NOT NULL,
	country VARCHAR(64) NOT NULL,
	state VARCHAR(64) NOT NULL,
	zip VARCHAR(64) NOT NULL,
	PRIMARY KEY(ownerId,addressId),
	FOREIGN KEY (ownerId) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE NO ACTION,
	);

CREATE TABLE IF NOT EXISTS user_locations (
	locationId CHAR(36) PRIMARY KEY,
	country VARCHAR(64) NOT NULL UNIQUE,
	enabled BOOLEAN DEFAULT false,
	userid CHAR(36) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE
	);


CREATE TABLE IF NOT EXISTS bookings (
	owerId CHAR(36) NOT NULL,
	oId CHAR(36),
	createdAt TIMESTAMP(0),
	owerId CHAR(36) NOT NULL,
	cart JSON NOT NULL,
	shippingDetails JSON NOT NULL,
	totalPrice DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	PRIMARY KEY(oId),
	FOREIGN KEY (ownerId) REFERENCES users(userid) ON DELETE CASCADE
	);

CREATE TABLE IF NOT EXISTS products (
	pid CHAR(36) PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	description VARCHAR(264) NOT NULL,
	category VARCHAR(264) NOT NULL,
	price DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	stockQuantity INT UNSIGNED  DEFAULT '0000' NOT NULL,
	image VARCHAR(64) NOT NULL,
	details VARCHAR(664) NOT NULL,
	productDiscount DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	auditInfo JSON NOT NULL
	);
	
CREATE TABLE IF NOT EXISTS reviews (
	rid CHAR(36) PRIMARY KEY,
	createdAt TIMESTAMP(0),
	rating INT UNSIGNED  DEFAULT '0000' NOT NULL,
	reviewerName VARCHAR(264) NOT NULL,
	comment VARCHAR(500) NOT NULL,
	email VARCHAR(64) NOT NULL,
	prodId CHAR(36) NOT NULL,
	FOREIGN KEY (pid) REFERENCES products(prodId) ON DELETE CASCADE
		) 
	`

func CreateTables(db *sqlx.DB) {

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(schema)
	// //DropTable(dbDriver)

	// statement, driverError := dbDriver.Prepare(schema)
	// if driverError != nil {
	// 	log.Println(driverError)
	// }
	// // Create tables
	// _, statementError := statement.Exec()
	// if statementError != nil {
	// 	log.Println("Table already exists!")
	// }
	// statement.Close()

	log.Println("All tables created/initialized successfully!")
}
