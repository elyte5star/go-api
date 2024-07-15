package response

import (
	"fmt"
	"time"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	d "github.com/mcuadros/go-defaults"
)

type NoContent struct {
}
type RecordNotFoundError struct{}

func (e *RecordNotFoundError) Error() string {
	return "record not found"
}

type RequestResponse struct {
	Path      string      `default:"0" json:"path"`
	Message   string      `default:"Operation Successful!" json:"message"`
	Success   bool        `default:"true" json:"success"`
	Code      int         `default:"200" json:"code"`
	TimeStamp time.Time   `json:"timeStamp"`
	Result    interface{} `json:"result"`
}

func NewResponse(c *fiber.Ctx) *RequestResponse {
	response := new(RequestResponse)
	d.SetDefaults(response)
	response.TimeStamp = util.TimeNow()
	response.Path = c.Route().Path
	return response

}

type ErrorResponse struct {
	Code    int    `default:"503" json:"code"`
	Message string `default:"Service unavailable" json:"message"`
	Success bool   `default:"false" json:"success"`
}

func NewErrorResponse() *ErrorResponse {
	err := new(ErrorResponse)
	d.SetDefaults(err)
	return err
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("statusCode %d: message %v: success %v", e.Code, e.Message, e.Success)
}
