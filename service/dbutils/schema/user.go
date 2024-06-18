package schema

import (
	"time"

	"github.com/google/uuid"
)

type UserAddress struct {
	AddressId     uuid.UUID   `db:"addressId" json:"addressId" validate:"required,uuid"`
	FullName      string      `json:"fullName"`
	StreetAddress string      `db:"fullName" json:"streetAddress"`
	Country       string      `db:"country" json:"country"`
	State         string      `db:"state" json:"state"`
	Zip           string      `db:"zip" json:"zip"`
}

type Userlocations struct {
	LocationId uuid.UUID   `db:"locationId" json:"locationId" validate:"required,uuid"`
	AuditInfo  AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	Country    string      `db:"country" json:"country" validate:"required"`
	Enabled    bool        `db:"enabled" json:"enabled" validate:"required"`
}

type Otp struct {
	OtpId      uuid.UUID   `db:"otpId" json:"otpId" validate:"required,uuid"`
	Email      string      `db:"email" json:"email" validate:"required"`
	OtpString  string      `db:"otpString" json:"otpString" validate:"required"`
	ExpiryDate time.Time   `db:"expiryDate" json:"expiryDate" validate:"required"`
}

type User struct {
	Userid           uuid.UUID   `db:"userid" json:"userid" validate:"required,uuid"`
	UserName         string      `db:"username" json:"username" validate:"required,lte=255"`
	Password         string      `db:"password" json:"password"  validate:"min=8,max=16"`
	Email            string      `db:"email" json:"email" validate:"required,email"`
	AccountNonLocked bool        `db:"accountNonLocked" json:"accountNonLocked"`
	Admin            bool        `db:"admin" json:"admin"`
	Enabled          bool        `db:"enabled" json:"enabled"`
	Telephone        string      `db:"telephone" json:"telephone" validate:"required"`
	Discount         float64     `db:"discount" json:"discount"`
	FailedAttempt    int         `db:"failedAttempt" json:"failedAttempt"`
	LockTime         *time.Time  `db:"lockTime" json:"lockTime"`
	AuditInfo        AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	//UserOtp          *Otp            `db:"otp" json:"otp"`
	// Address          *UserAddress    `db:"address" json:"address"`
	// Locations        []*Userlocation `db:"locations" json:"locations"`
	// Bookings         []*Booking      `db:"bookings" json:"bookings"`
}
