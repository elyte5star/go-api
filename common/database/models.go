package database

import (
	"github.com/gofrs/uuid"
)

type User struct {
	UserId           uuid.UUID `json:"userid"`
	UserName         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	AccountNonLocked bool      `json:"account_not_locked"`
	Admin            bool      `json:"admin"`
	Enabled          bool      `json:"enabled"`
	Telephone        string    `json:"telephone"`
	Discount         float64   `json:"discount"`
}

type Product struct {
	Pid           uuid.UUID `json:"pid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	Price         float32   `json:"price"`
	StockQuantity int       `json:"stockQuantity"`
	Image         string    `json:"image"`
	Details       string    `json:"details"`
}


type Review struct {
	Rid          int    `json:"rid"`
	Rating       int    `json:"rating"`
	ReviewerName string `json:"reviewerName"`
	Comment      string `json:"comment"`
	Email        string `json:"email"`
}
