package repository

import (
	"github.com/gofrs/uuid"
)

type GetUserResponse struct {
	UserId           uuid.UUID `json:"userid"`
	LastModifiedAt   string    `json:"lastModifiedAt"`
	CreatedAt        string    `json:"createdAt"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	AccountNonLocked bool      `json:"account_not_locked"`
	Admin            bool      `json:"admin"`
	Enabled          bool      `json:"enabled"`
	IsUsing2FA       bool      `json:"isUsing2FA"`
	Telephone        string    `json:"telephone"`
}

type TokenResponse struct{
	UserId       uuid.UUID `json:"userid"`
	UserName string    `json:"username"`
	Email  string    `json:"email"`
	Enabled bool `json:"enabled"`
	Admin bool `json:"admin"`
	AccessToken string    `json:"accessToken"`
	TokenType string    `json:"tokenType"`

}

type GetUsersResponse struct {
	Users []GetUsersResponse `json:"users"`
}
