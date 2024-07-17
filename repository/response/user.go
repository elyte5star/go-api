package response

import (
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	Userid           uuid.UUID              `json:"userid"`
	LastModifiedAt   *time.Time             `json:"lastModifiedAt"`
	CreatedAt        time.Time              `json:"createdAt"`
	Username         string                 `json:"username"`
	Email            string                 `json:"email"`
	AccountNonLocked bool                   `json:"accountNonLocked"`
	Admin            bool                   `json:"admin"`
	Enabled          bool                   `json:"enabled"`
	IsUsing2FA       bool                   `json:"isUsing2FA"`
	Telephone        string                 `json:"telephone"`
	LockTime         *time.Time             `json:"lockTime"`
	Address          *GetUserAdressResponse `json:"address"`
	Bookings         []*GetBookingResponse  `json:"bookings"`
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

type GetUserAdressResponse struct {
	FullName      string ` json:"fullName"`
	StreetAddress string `json:"streetAddress"`
	Country       string `json:"country"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
}

type GetBookingResponse struct {
}
