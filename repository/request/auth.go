package request

import "github.com/google/uuid"

// swagger:parameters LoginRequest
type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password"  validate:"min=5,max=30"`
}

type UserCredentials struct {
	Userid                  uuid.UUID `json:"userid"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	IsEnabled               bool      `json:"isEnabled "`
	IsAccountNonLocked      bool      `json:"isAccountNonLocked"`
	IsCredentialsNonExpired bool      `json:"isCredentialsNonExpired"`
	IsAdmin                 bool      `json:"isAdmin"`
	TokenId                 string    `json:"tokenId"`
}
