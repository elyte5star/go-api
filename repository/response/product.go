package response

import (
	"time"

	"github.com/google/uuid"
)

type GetProductResponse struct {
	Pid             uuid.UUID                   `json:"pid"`
	Name            string                      `json:"name"`
	Description     string                      `json:"description"`
	Category        string                      `json:"category"`
	Price           float64                     `json:"price"`
	StockQuantity   int                         `json:"stockQuantity"`
	Image           string                      `json:"image"`
	Details         string                      `json:"details"`
	ProductDiscount float64                     `json:"productDiscount"`
	Reviews         []*GetProductReviewResponse `json:"reviews"`
}

// Products struct
type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
	Count    int               `json:"count"`
}

type GetProductReviewResponse struct {
	Rid          uuid.UUID `json:"rid"`
	CreatedAt    time.Time `json:"createdAt"`
	Rating       int       `json:"rating"`
	ReviewerName string    `json:"reviewerName"`
	Comment      string    `json:"comment"`
	Email        string    `json:"email"`
}

type GetProductReviewsResponse struct {
	Reviews []GetProductReviewResponse `json:"reviews"`
}
