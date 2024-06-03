package response

type NoContent struct {
}

type StatusResponse struct {
	Path      string      `json:"path"`
	Message   string      `json:"mesage"`
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	TimeStamp string      `json:"timeStamp"`
	Result    interface{} `json:"result"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}
