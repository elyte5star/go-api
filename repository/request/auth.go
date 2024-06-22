package request



type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=10"`
	Password string `json:"password"  validate:"min=5,max=30"`
}

