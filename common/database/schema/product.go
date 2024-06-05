package schema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

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

// This method simply returns the JSON-encoded representation of the struct.
func (r Review) Value() (driver.Value, error) {
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
