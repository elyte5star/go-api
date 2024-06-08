package response

import (
	"fmt"
	"time"
)

type NoContent struct {
}

type RequestResponse struct {
	Path      string      `default:"0" json:"path"`
	Message   string      `default:"Operation Successful!" json:"message"`
	Success   bool        `default:"true" json:"success"`
	Code      int         `default:"200" json:"code"`
	TimeStamp time.Time   `json:"timeStamp"`
	Result    interface{} `default:"nil" json:"result"`
}

type ErrorResponse struct {
	Code    int    `default:"500" json:"code"`
	Message string `default:"Something went wrong" json:"cause"`
	Success bool   `default:"false" json:"success"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("statusCode %d: message %v: success %v", e.Code, e.Message, e.Success)
}
