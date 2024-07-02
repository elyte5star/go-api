package request

import "github.com/google/uuid"

type CreateProductRequest struct {
	Name            string  `json:"name" validate:"required"`
	StockQuantity   int     `json:"stockQuantity" validate:"required"`
	Image           string  `json:"image" validate:"required"`
	Details         string  `json:"details" validate:"required,lte=555"`
	Category        string  `json:"category" validate:"required,lte=255"`
	ProductDiscount float64 `json:"productDiscount" validate:"required,omitempty"`
	Price           float32 `json:"price" validate:"required"`
	Description     string  `json:"description,omitempty" validate:"required,lte=555"`
}

type CreateProductReviewRequest struct {
	Pid          uuid.UUID `json:"pid" validate:"required,uuid"`
	Rating       int       `json:"rating" validate:"min=1,max=5"`
	ReviewerName string    `json:"reviewerName" validate:"required"`
	Comment      string    `json:"comment"  validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
}
