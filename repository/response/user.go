package response

import (
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	Userid           uuid.UUID              `json:"userid"`
	Username         string                 `json:"username"`
	Email            string                 `json:"email"`
	AccountNonLocked bool                   `json:"accountNonLocked"`
	Admin            bool                   `json:"admin"`
	Enabled          bool                   `json:"enabled"`
	IsUsing2FA       bool                   `json:"isUsing2FA"`
	Telephone        string                 `json:"telephone"`
	LockTime         *time.Time             `json:"lockTime"`
	LastModifiedAt   *time.Time             `json:"lastModifiedAt"`
	CreatedAt        time.Time              `json:"createdAt"`
	Discount         *float64               `json:"discount"`
	Address          *GetUserAdressResponse `json:"address"`
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

type GetUserAdressResponse struct {
	FullName      *string ` json:"fullName,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Country       *string `json:"country,omitempty"`
	State         *string `json:"state,omitempty"`
	Zip           *string `json:"zip,omitempty"`
}

type GetBookingResponse struct {
}
