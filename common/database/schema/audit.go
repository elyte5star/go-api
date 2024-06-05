package schema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type AuditEntity struct {
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
	CreatedBy      string    `json:"CreatedBy"`
}



// This method simply returns the JSON-encoded representation of the struct.
func (r AuditEntity) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan make the AuditEntity struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (r *AuditEntity) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &r)
}
