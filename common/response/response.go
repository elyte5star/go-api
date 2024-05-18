package response
import (
	"github.com/gofrs/uuid"
)


type NoContent struct {
}

type StatusResponse struct {
	Status string `json:"status"`
	Message string `json:"mesage"`
	Success bool `json:"success"`
	Code  int    `json:"code"`
	TimeStamp  string `json:"timeStamp"`
	Result string  `json:"result"`

}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
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