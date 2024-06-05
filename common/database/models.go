package database

import (
	"time"

	"github.com/google/uuid"
)

type AuditEntity struct {
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
	CreatedBy      string    `json:"CreatedBy"`
}
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

type Booking struct {
	Oid             uuid.UUID    `json:"oid" validate:"required,uuid"`
	AuditInfo       AuditEntity  `json:"auditInfo"`
	ShippingDetails UserAddress  `json:"shippingDetails"`
	TotalPrice      float64      `json:"totalPrice"`
	Cart            []ItemIncart `json:"cart"`
}

type ItemIncart struct {
	Pid             uuid.UUID `json:"pid" validate:"required,uuid"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:"required,lte=555"`
	Category        string    `json:"category" validate:"required,lte=255"`
	Price           float32   `json:"price"`
	StockQuantity   int       `json:"stockQuantity"`
	Image           string    `json:"image" validate:"required,lte=255"`
	Details         string    `json:"details" validate:"required,lte=555"`
	CalculatedPrice float64   `json:"calculatedPrice"`
	Quantity        int       `json:"quantity"`
}
type User struct {
	UserId           uuid.UUID       `json:"userid" validate:"required,uuid"`
	AuditInfo        AuditEntity     `json:"auditInfo"`
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
	UserOtp          *Otp            `json:"otp"`
	Address          *UserAddress    `json:"address"`
	Locations        []*Userlocation `json:"locations"`
	Bookings         []*Booking      `json:"bookings"`
}

type Product struct {
	Pid             uuid.UUID   `json:"pid" validate:"required,uuid"`
	AuditInfo       AuditEntity `json:"auditInfo"`
	Name            string      `json:"name" validate:"required"`
	Description     string      `json:"description" validate:"required,lte=555"`
	Category        string      `json:"category" validate:"required,lte=255"`
	Price           float32     `json:"price"`
	StockQuantity   int         `json:"stockQuantity"`
	Image           string      `json:"image"`
	Details         string      `json:"details" validate:"required,lte=555"`
	Reviews         []*Review   `json:"reviews"`
	ProductDiscount float64     `json:"productDiscount"`
}

type Review struct {
	Rid          uuid.UUID   `json:"rid" validate:"required,uuid"`
	AuditInfo    AuditEntity `json:"auditInfo"`
	Rating       int         `json:"rating" validate:"min=1,max=5"`
	ReviewerName string      `json:"reviewerName" validate:"required"`
	Comment      string      `json:"comment"  validate:"required"`
	Email        string      `json:"email" validate:"required,email"`
}
