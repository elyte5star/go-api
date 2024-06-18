package schema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Pid             uuid.UUID   `db:"pid" json:"pid" validate:"required,uuid"`
	AuditInfo       AuditEntity `db:"auditInfo" json:"auditInfo" validate:"required,dive"`
	Name            string      `db:"name" json:"name" validate:"required"`
	Description     string      `db:"description" json:"description" validate:"required,lte=555"`
	Category        string      `db:"category" json:"category" validate:"required,lte=255"`
	Price           float32     `db:"price" json:"price" validate:"required"`
	StockQuantity   int         `db:"stockQuantity" json:"stockQuantity" validate:"required"`
	Image           string      `db:"image" json:"image" validate:"required"`
	Details         string      `db:"details" json:"details" validate:"required,lte=555"`
	ProductDiscount float64     `db:"productDiscount" json:"productDiscount"`
}

type Review struct {
	Rid          uuid.UUID  `db:"rid" json:"rid" validate:"required,uuid"`
	CreatedAt    *time.Time `db:"createdAt" json:"createdAt"`
	Rating       int        `db:"rating" json:"rating" validate:"min=1,max=5"`
	ReviewerName string     `db:"reviewerName" json:"reviewerName" validate:"required"`
	Comment      string     `db:"comment" json:"comment"  validate:"required"`
	Email        string     `db:"email" json:"email" validate:"required,email"`
}

// This method simply returns the JSON-encoded representation of the struct.
func (r *Review) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan make the Review struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (r *Review) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &r)
}
