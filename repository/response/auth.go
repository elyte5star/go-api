package response

import "github.com/google/uuid"

type TokenResponse struct {
	UserId           uuid.UUID `json:"userid"`
	UserName         string    `json:"username"`
	Email            string    `json:"email"`
	Enabled          bool      `json:"enabled"`
	AccountNonLocked bool      `json:"account_not_locked"`
	Admin            bool      `json:"admin"`
	AccessToken      string    `json:"accessToken"`
	TokenType        string    `json:"tokenType"`
}
