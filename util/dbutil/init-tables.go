package dbutil

import (
	"log/slog"
	"strings"

	"github.com/api/repository/schema"
	"github.com/api/service"
	"github.com/api/util"
	"github.com/jmoiron/sqlx"
)

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
	lockTime DATETIME NULL,
	auditInfo LONGTEXT NOT NULL
	) ENGINE=INNODB DEFAULT CHARSET=utf8;
`

const otp = `CREATE TABLE IF NOT EXISTS otp (
		userid CHAR(36) PRIMARY KEY,
		email VARCHAR(64) NOT NULL UNIQUE,
		otpString VARCHAR(64) NOT NULL,
		expiryDate DATETIME NOT NULL,
		FOREIGN KEY(userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
		) ENGINE=INNODB DEFAULT CHARSET=utf8;
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
	) ENGINE=INNODB DEFAULT CHARSET=utf8;
`
const userLocations = `CREATE TABLE IF NOT EXISTS user_locations (
	locationId CHAR(36) PRIMARY KEY,
	country VARCHAR(64) NOT NULL UNIQUE,
	enabled BOOLEAN DEFAULT false,
	userid CHAR(36) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
	) ENGINE=INNODB DEFAULT CHARSET=utf8;
`

const bookings = `CREATE TABLE IF NOT EXISTS bookings (
	oId CHAR(36) PRIMARY KEY,
	createdAt DATETIME NOT NULL,
	userid CHAR(36) NOT NULL,
	cart JSON NOT NULL,
	shippingDetails JSON NOT NULL,
	totalPrice DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
	) ENGINE=INNODB DEFAULT CHARSET=utf8;
`
const products = `CREATE TABLE IF NOT EXISTS products (
	pid CHAR(36) PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	description VARCHAR(264) NOT NULL,
	category VARCHAR(264) NOT NULL,
	price DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	stockQuantity INT UNSIGNED  DEFAULT '0000' NOT NULL,
	image VARCHAR(64) NOT NULL UNIQUE,
	details VARCHAR(1500) NOT NULL,
	productDiscount DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
	auditInfo LONGTEXT NOT NULL
	) ENGINE=INNODB DEFAULT CHARSET=utf8;
`
const productReview = `CREATE TABLE IF NOT EXISTS reviews (
	rid CHAR(36) PRIMARY KEY,
	createdAt DATETIME NOT NULL,
	rating INT UNSIGNED  DEFAULT '0000' NOT NULL,
	reviewerName VARCHAR(264) NOT NULL,
	comment VARCHAR(500) NOT NULL,
	email VARCHAR(64) NOT NULL,
	pid CHAR(36) NOT NULL,
	FOREIGN KEY (pid) REFERENCES products(pid) ON DELETE CASCADE ON UPDATE CASCADE
		) ENGINE=INNODB DEFAULT CHARSET=utf8;
	`

const dropOtpTable = `DROP TABLE IF EXISTS otp;`
const dropAddressTable = `DROP TABLE IF EXISTS address;`
const dropBookingsTable = `DROP TABLE IF EXISTS bookings;`
const dropUserLocationTable = `DROP TABLE IF EXISTS user_locations;`
const dropUserTable = `DROP TABLE IF EXISTS users;`
const dropReviewTable = `DROP TABLE IF EXISTS reviews;`

//const dropProductsTable = `DROP TABLE IF EXISTS products;`

func LoadDatabase(dbDriver *sqlx.DB, cfg *service.AppConfig) {
	log := cfg.Logger

	defer dbDriver.Close()

	//Droptables(cfg.Logger, dbDriver)

	statement, _ := dbDriver.Prepare(users)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(otp)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(userAddress)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(userLocations)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(bookings)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(products)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(productReview)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())

	}

	log.Info("All tables created/initialized successfully!")
	CreateAdminAccount(cfg)
}
func Droptables(log *slog.Logger, dbDriver *sqlx.DB) {
	statement, _ := dbDriver.Prepare(dropOtpTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(dropAddressTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(dropBookingsTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(dropUserLocationTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(dropUserTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	statement, _ = dbDriver.Prepare(dropReviewTable)
	if _, err := statement.Exec(); err != nil {
		log.Error(err.Error())
	}
	// statement, _ = dbDriver.Prepare(dropProductsTable)
	// if _, err := statement.Exec(); err != nil {
	// 	log.Error(err.Error())
	// }

}
func CreateAdminAccount(cfg *service.AppConfig) {
	db, err := service.DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return
	}
	user := new(schema.User)
	user.Userid = util.Ident()
	user.Username = cfg.SmtpUsername
	user.SetPassword("string")
	user.Email = cfg.SupportEmail
	user.Telephone = "234802394"
	user.AccountNonLocked = true
	user.FailedAttempt = 0
	user.Discount = 0.0
	user.Admin = true
	user.IsUsing2FA = true
	user.Enabled = true
	audit := &schema.AuditEntity{CreatedAt: util.TimeNow(), LastModifiedBy: "none", CreatedBy: user.Userid.String()}
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

