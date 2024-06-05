package repository

import "github.com/google/uuid"

type GetProductResponse struct {
	Pid           uuid.UUID `json:"pid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	Price         float32   `json:"price"`
	StockQuantity int       `json:"stockQuantity"`
	Image         string    `json:"image"`
	Details       string    `json:"details"`
	Reviews       []Review  `json:"reviews"`
}

// Products struct
type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}
