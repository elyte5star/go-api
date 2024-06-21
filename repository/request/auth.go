package request

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=10"`
	Password string `json:"password"  validate:"min=5,max=30"`
}

type UserPrincipal struct {
	Userid                  uuid.UUID `json:"userid"`
	Username                string    `json:"username"`
	Password                []byte    `json:"password"`
	Email                   string    `json:"email"`
	IsEnabled               bool      `json:"isEnabled "`
	IsAccountNonLocked      bool      `json:"isAccountNonLocked"`
	IsCredentialsNonExpired bool      `json:"isCredentialsNonExpired"`
	IsAdmin                 bool      `json:"isAdmin"`
}
