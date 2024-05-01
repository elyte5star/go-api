package model

import (
	"github.com/gofrs/uuid"
)

type User struct {
	UserId       uuid.UUID `json:"userid"`
	UserName string    `json:"username"`
	Password  string    `json:"password"`
	Email  string    `json:"email"`
	AccountNonLocked bool `json:"account_not_locked"`
	Admin bool `json:"admin"`
	Enabled bool `json:"enabled"`
	Telephone string `json:"telephone"`
	Discount float64 `json:"discount"`
}