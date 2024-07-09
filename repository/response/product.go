package response

import (
	"github.com/api/service/dbutils/schema"
	"github.com/google/uuid"
)

type GetProductResponse struct {
	Pid             uuid.UUID        `json:"pid"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Category        string           `json:"category"`
	Price           float64          `json:"price"`
	StockQuantity   int              `json:"stockQuantity"`
	Image           string           `json:"image"`
	Details         string           `json:"details"`
	ProductDiscount float64          `json:"productDiscount"`
	Reviews         []*schema.Review `json:"reviews"`
}

// Products struct
type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}
