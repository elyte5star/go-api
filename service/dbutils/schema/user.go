package schema

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserAddress struct {
	AddressId     uuid.UUID `db:"addressId" json:"addressId" validate:"required,uuid"`
	FullName      string    `json:"fullName"`
	StreetAddress string    `db:"fullName" json:"streetAddress"`
	Country       string    `db:"country" json:"country"`
	State         string    `db:"state" json:"state"`
	Zip           string    `db:"zip" json:"zip"`
}

type Userlocations struct {
	LocationId uuid.UUID   `db:"locationId" json:"locationId" validate:"required,uuid"`
	AuditInfo  AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	Country    string      `db:"country" json:"country" validate:"required"`
	Enabled    bool        `db:"enabled" json:"enabled" validate:"required"`
}

type Otp struct {
	OtpId      uuid.UUID `db:"otpId" json:"otpId" validate:"required,uuid"`
	Email      string    `db:"email" json:"email" validate:"required"`
	OtpString  string    `db:"otpString" json:"otpString" validate:"required"`
	ExpiryDate time.Time `db:"expiryDate" json:"expiryDate" validate:"required"`
}

type User struct {
	Userid           uuid.UUID   `db:"userid" json:"userid" validate:"required,uuid"`
	UserName         string      `db:"username" json:"username" validate:"required,lte=255"`
	Password         []byte      `db:"password" json:"password"  validate:"required"`
	Email            string      `db:"email" json:"email" validate:"required,email"`
	AccountNonLocked bool        `db:"accountNonLocked" json:"accountNonLocked"`
	Admin            bool        `db:"admin" json:"admin"`
	Enabled          bool        `db:"enabled" json:"enabled"`
	IsUsing2FA       bool        `db:"isUsing2FA" json:"isUsing2FA"`
	Telephone        string      `db:"telephone" json:"telephone" validate:"required"`
	Discount         float64     `db:"discount" json:"discount"`
	FailedAttempt    int         `db:"failedAttempt" json:"failedAttempt"`
	LockTime         time.Time   `db:"lockTime" json:"lockTime"`
	AuditInfo        AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
}

// SetPassword: sets the hashed password to the user struct defined above
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), 12)
	user.Password = hashedPassword
}

// ComparePassword: Used to compare user stored password and  login  password
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(strings.TrimSpace(password)))
}
