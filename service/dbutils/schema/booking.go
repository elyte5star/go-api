package schema

import (
	"time"
	"github.com/google/uuid"
)

type Booking struct {
	Oid             uuid.UUID    `db:"oid" json:"oid" validate:"required,uuid"`
	ShippingDetails UserAddress  `db:"shippingDetails" json:"shippingDetails" validate:"required,dive"`
	CreatedAt       *time.Time   `db:"createdAt" json:"createdAt"  validate:"required"`
	TotalPrice      float64      `db:"totalPrice" json:"totalPrice"  validate:"required"`
	Cart            []ItemIncart `db:"cart" json:"cart" validate:"required,dive"`
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
	CalculatedPrice float64   `json:"calculatedPrice" validate:"required"`
	Quantity        int       `json:"quantity"  validate:"min=1"`
}
