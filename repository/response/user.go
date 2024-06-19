package response

import (
	"time"

	"github.com/api/service/dbutils/schema"
	"github.com/google/uuid"
	
)


type GetUserResponse struct {
	Userid           uuid.UUID           `json:"userid"`
	LastModifiedAt   time.Time           `json:"lastModifiedAt"`
	CreatedAt        time.Time           `json:"createdAt"`
	Username         string              `json:"username"`
	Email            string              `json:"email"`
	AccountNonLocked bool                `json:"account_not_locked"`
	Admin            bool                `json:"admin"`
	Enabled          bool                `json:"enabled"`
	IsUsing2FA       bool                `json:"isUsing2FA"`
	Telephone        string              `json:"telephone"`
	LockTime         time.Time           `json:"lockTime"`
	Address          *schema.UserAddress `json:"address"`
	Bookings         []*schema.Booking   `json:"bookings"`
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}