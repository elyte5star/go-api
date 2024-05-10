
package repository
import (
	"github.com/gofrs/uuid"
)




type GetUserResponse struct {
	UserId       uuid.UUID `json:"userid"`
	Username string    `json:"username"`
	Password  string    `json:"password"`
}