package response

import (
	"fmt"
	"time"

	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
)

type NoContent struct {
}
type RecordNotFoundError struct{}

func (e *RecordNotFoundError) Error() string {
	return "record not found"
}

type RequestResponse struct {
	Path      string      `json:"path"`
	Message   string      `json:"message"`
	Success   bool        `default:"true" json:"success"`
	Code      int         `json:"code"`
	TimeStamp time.Time   `json:"timeStamp"`
	Result    interface{} `json:"result"`
}

func NewResponse(c *fiber.Ctx) *RequestResponse {
	return &RequestResponse{
		Path:      c.Route().Path,
		Message:   "Operation Successful!",
		Code:      200,
		TimeStamp: util.TimeNow(),
		Success:   true,
	}

}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `default:"false" json:"success"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{Code: 503, Message: "Service unavailable", Success: false}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("statusCode %d: message %v: success %v", e.Code, e.Message, e.Success)
}
