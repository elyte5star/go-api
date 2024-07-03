package dbutils

import (
	"log"
	"strings"

	"github.com/api/service"
	"github.com/api/service/dbutils/schema"
	"github.com/api/util"
	"github.com/jmoiron/sqlx"
)

// const schema = `

// DROP TABLE IF EXISTS users,otp,address,user_locations,bookings;
// `
const users = `CREATE TABLE IF NOT EXISTS users (
	userid CHAR(36) PRIMARY KEY,
	username VARCHAR(64) NOT NULL UNIQUE,
	password VARBINARY(255) NOT NULL,
	email VARCHAR(64) NOT NULL UNIQUE,
	accountNonLocked BOOLEAN DEFAULT false NOT NULL,
	admin BOOLEAN DEFAULT false NOT NULL,
	enabled BOOLEAN DEFAULT false  NOT NULL,
	isUsing2FA BOOLEAN DEFAULT false NOT NULL,
	telephone VARCHAR(64) NOT NULL UNIQUE,
	discount DECIMAL(16,2) DEFAULT '0.00',
	failedAttempt INT UNSIGNED  DEFAULT '0000',
	lockTime TIMESTAMP(0),
	auditInfo JSON NOT NULL
	) ENGINE=INNODB;
`

const otp = `CREATE TABLE IF NOT EXISTS otp (
		userid CHAR(36) PRIMARY KEY,
		email VARCHAR(64) NOT NULL UNIQUE,
		otpString VARCHAR(64) NOT NULL,
		expiryDate TIMESTAMP(0),
		FOREIGN KEY(userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
		) ENGINE=INNODB;
`

const userAddress = `CREATE TABLE IF NOT EXISTS address (
	userid CHAR(36),
	fullName VARCHAR(64) NOT NULL UNIQUE,
	streetAddress VARCHAR(64) NOT NULL,
	country VARCHAR(64) NOT NULL,
	state VARCHAR(64) NOT NULL,
	zip VARCHAR(64) NOT NULL,
	PRIMARY KEY (userid),
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
	) ENGINE=INNODB;
`
const userLocations = `CREATE TABLE IF NOT EXISTS user_locations (
	locationId CHAR(36) PRIMARY KEY,
	country VARCHAR(64) NOT NULL UNIQUE,
	enabled BOOLEAN DEFAULT false,
	userid CHAR(36) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
	) ENGINE=INNODB;
`

const bookings = `CREATE TABLE IF NOT EXISTS bookings (
	oId CHAR(36) PRIMARY KEY,
	createdAt TIMESTAMP(0),
	userid CHAR(36) NOT NULL,
	cart JSON NOT NULL,
	shippingDetails JSON NOT NULL,
	totalPrice DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
	) ENGINE=INNODB;
`
const products = `CREATE TABLE IF NOT EXISTS products (
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
	) ENGINE=INNODB;
`
const productReview = `CREATE TABLE IF NOT EXISTS reviews (
	rid CHAR(36) PRIMARY KEY,
	createdAt TIMESTAMP(0),
	rating INT UNSIGNED  DEFAULT '0000' NOT NULL,
	reviewerName VARCHAR(264) NOT NULL,
	comment VARCHAR(500) NOT NULL,
	email VARCHAR(64) NOT NULL,
	pid CHAR(36) NOT NULL,
	FOREIGN KEY (pid) REFERENCES products(pid) ON DELETE CASCADE ON UPDATE CASCADE
		) ENGINE=INNODB;
	`

func CreateTables(dbDriver *sqlx.DB) {

	defer dbDriver.Close()

	statement, driverError := dbDriver.Prepare(users)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
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
	statement.Close()

	log.Println("All tables created/initialized successfully!")
}

func CreateAdminAccount(username string, cfg *service.AppConfig) {
	db, err := service.DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return
	}
	user := new(schema.User)
	user.Userid = util.Ident()
	user.UserName = username
	user.SetPassword("string")
	user.Email = "elyte5star@gmail.com"
	user.LockTime = util.TimeThen()
	user.Telephone = "234802394"
	user.AccountNonLocked = true
	user.Admin = true
	user.IsUsing2FA = true
	user.Enabled = true
	audit := &schema.AuditEntity{CreatedAt: util.TimeNow(), LastModifiedAt: util.NullTime(), LastModifiedBy: "none", CreatedBy: username}
	user.AuditInfo = *audit
	if err := db.CreateUser(user); err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			cfg.Logger.Warn("Admin user already exist")
			return
		}
		cfg.Logger.Error(err.Error())
		return
	}
	cfg.Logger.Info("Admin account created")

}

func CreateProducts(user *[]schema.Product) {

}
