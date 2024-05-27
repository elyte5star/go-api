
package repository
import (
	"github.com/gofrs/uuid"
)




type GetUserResponse struct {
	UserId       uuid.UUID `json:"userid"`
	LastModifiedAt string `json:"lastModifiedAt"`
	Username string    `json:"username"`
	Email  string    `json:"email"`
	AccountNonLocked bool `json:"account_not_locked"`
	Admin bool `json:"admin"`
	Enabled bool `json:"enabled"`
	Telephone string `json:"telephone"`
}