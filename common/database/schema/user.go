package schema

import (
	"errors"
	"time"

	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type UserAddress struct {
	AddressId     uuid.UUID   `json:"addressId" validate:"required,uuid"`
	AuditInfo     AuditEntity `json:"auditInfo"`
	FullName      string      `json:"fullName"`
	StreetAddress string      `json:"streetAddress"`
	Country       string      `json:"country"`
	State         string      `json:"state"`
	Zip           string      `json:"zip"`
}

type Userlocation struct {
	LocationId uuid.UUID   `json:"locationId" validate:"required,uuid"`
	AuditInfo  AuditEntity `json:"auditInfo"`
	Country    string      `json:"country" validate:"required"`
	Enabled    bool        `json:"enabled"`
}

type Otp struct {
	OtpId      uuid.UUID   `json:"otpId" validate:"required,uuid"`
	AuditInfo  AuditEntity `json:"auditInfo"`
	Email      string      `json:"email"`
	OtpString  string      `json:"otpString" validate:"required"`
	ExpiryDate time.Time   `json:"expiryDate"`
}


type User struct {
	Userid           uuid.UUID       `json:"userid" validate:"required,uuid"`
	AuditInfo        AuditEntity     `json:"auditInfo" validate:"required,dive"`
	UserName         string          `json:"username" validate:"required,lte=255"`
	Password         string          `json:"password" validate:"required"`
	Email            string          `json:"email" validate:"required,email"`
	AccountNonLocked bool            `json:"account_not_locked"`
	Admin            bool            `json:"admin"`
	Enabled          bool            `json:"enabled"`
	Telephone        string          `json:"telephone" validate:"min=5,max=16"`
	Discount         float64         `json:"discount"`
	FailedAttempt    int             `json:"failedAttempt"`
	LockTime         time.Time       `json:"lockTime"`
	UserOtp          Otp             `json:"otp"`
	Address          UserAddress     `json:"address"`
	Locations        []*Userlocation `json:"locations"`
	Bookings         []*Booking      `json:"bookings"`
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
