package request

import "github.com/google/uuid"

type CreateProductRequest struct {
	Name            string   `json:"name" validate:"required,min=3,max=20"`
	StockQuantity   int      `json:"stockQuantity" validate:"gte=0,lte=1000"`
	Image           string   `json:"image" validate:"required"`
	Details         string   `json:"details" validate:"lte=1500"`
	Category        string   `json:"category" validate:"lte=255"`
	ProductDiscount *float64 `json:"productDiscount,omitempty"`
	Price           float64  `json:"price" validate:"required"`
	Description     string   `json:"description" validate:"max=555"`
}

type CreateProductReviewRequest struct {
	Pid          uuid.UUID `json:"pid" validate:"required,uuid"`
	Rating       int       `json:"rating" validate:"min=1,max=5"`
	ReviewerName string    `json:"reviewerName" validate:"required"`
	Comment      string    `json:"comment"  validate:"required"`
	Email        string    `json:"email" validate:"email"`
}

type CreateProductsRequest struct {
	Products []CreateProductRequest `json:"products" validate:"required,dive,required"`
}

type GetproductsQuery struct {
	Page int    `query:"page"`
	Size int    `query:"size"`
	Sort string `query:"sort"`
}
