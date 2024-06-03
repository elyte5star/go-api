package response

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

