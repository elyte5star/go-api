package schema

import (
	"errors"
	"time"

	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type UserAddress struct {
	AddressId     uuid.UUID   `db:"addressId" json:"addressId" validate:"required,uuid"`
	AuditInfo     AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	FullName      string      `json:"fullName"`
	StreetAddress string      `db:"fullName" json:"streetAddress"`
	Country       string      `db:"country" json:"country"`
	State         string      `db:"state" json:"state"`
	Zip           string      `db:"zip" json:"zip"`
}

type Userlocation struct {
	LocationId uuid.UUID   `db:"locationId" json:"locationId" validate:"required,uuid"`
	AuditInfo  AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	Country    string      `db:"country" json:"country" validate:"required"`
	Enabled    bool        `db:"enabled" json:"enabled" validate:"required"`
}

type Otp struct {
	OtpId      uuid.UUID   `db:"otpId" json:"otpId" validate:"required,uuid"`
	AuditInfo  AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	Email      string      `db:"email" json:"email" validate:"required"`
	OtpString  string      `db:"otpString" json:"otpString" validate:"required"`
	ExpiryDate time.Time   `db:"expiryDate" json:"expiryDate" validate:"required"`
}

type User struct {
	Userid           uuid.UUID       `db:"userid" json:"userid" validate:"required,uuid"`
	AuditInfo        AuditEntity     `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	UserName         string          `db:"username" json:"username" validate:"required,lte=255"`
	Password         string          `db:"password" json:"password"  validate:"min=8,max=16"`
	Email            string          `db:"email" json:"email" validate:"required,email"`
	AccountNonLocked bool            `db:"accountNonLocked" json:"accountNonLocked"`
	Admin            bool            `db:"admin" json:"admin"`
	Enabled          bool            `db:"admin" json:"enabled"`
	Telephone        string          `db:"telephone" json:"telephone" validate:"min=5,max=16"`
	Discount         float64         `db:"discount" json:"discount"`
	FailedAttempt    int             `db:"failedAttempt" json:"failedAttempt"`
	LockTime         time.Time       `db:"lockTime" json:"lockTime"`
	UserOtp          *Otp            `db:"otp" json:"otp"`
	Address          *UserAddress    `db:"address" json:"address"`
	Locations        []*Userlocation `db:"locations" json:"locations"`
	Bookings         []*Booking      `db:"bookings" json:"bookings"`
}

// This method simply returns the JSON-encoded representation of the struct.
func (otp Otp) Value() (driver.Value, error) {
	return json.Marshal(otp)
}

// Scan make the Otp struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (otp *Otp) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &otp)
}
