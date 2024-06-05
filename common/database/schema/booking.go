package schema
import "github.com/google/uuid"

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