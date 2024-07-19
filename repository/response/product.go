package response

import (
	"time"

	"github.com/google/uuid"
)

type GetProductResponse struct {
	Pid             uuid.UUID   `json:"pid"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Category        string      `json:"category"`
	Price           float64     `json:"price"`
	StockQuantity   int         `json:"stockQuantity"`
	Image           string      `json:"image"`
	Details         string      `json:"details"`
	ProductDiscount float64     `json:"productDiscount"`
	Reviews         interface{} `json:"reviews"`
}

// Products struct
type GetProductsResponse struct {
	Products         []GetProductResponse `json:"products"`
	NumberOfElements int                  `json:"numberOfElements"`
}

type GetProductReviewResponse struct {
	Rid          *uuid.UUID `json:"rid,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	Rating       *int       `json:"rating,omitempty"`
	ReviewerName *string    `json:"reviewerName,omitempty"`
	Comment      *string    `json:"comment,omitempty"`
	Email        *string    `json:"email,omitempty"`
}

type GetProductReviewsResponse struct {
	Reviews []GetProductReviewResponse `json:"reviews"`
}
