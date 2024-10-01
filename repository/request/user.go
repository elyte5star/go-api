package request

type CreateUserRequest struct {
	Username        string   `json:"username" validate:"min=5,max=30"`
	Password        string   `json:"password" validate:"eqfield=ConfirmPassword,min=5,max=30"`
	ConfirmPassword string   `json:"confirmPassword" validate:"min=5,max=30"`
	Email           string   `json:"email" validate:"required,email"`
	Telephone       string   `json:"telephone" validate:"required,tel"`
	Discount        *float64 `json:"discount,omitempty"`
}

type ModifyUser struct {
	Username  string            `json:"username" validate:"min=5,max=30"`
	Telephone string            `json:"telephone" validate:"tel"`
	Address   *CreateAddressReq `json:"address,omitempty"`
}

type ModifyUserPassword struct {
	OldPassword        string `json:"oldPassword" validate:"nefield=newPassword,min=5,max=30"`
	NewPassword        string `json:"newPassword" validate:"eqfield=ConfirmNewPassword,min=5,max=30"`
	ConfirmNewPassword string `json:"confirmNewPassword"  validate:"min=5,max=30"`
}

type CreateAddressReq struct {
	FullName      string `json:"fullName" validate:"required"`
	StreetAddress string `json:"streetAddress" validate:"required"`
	Country       string `json:"country" validate:"required"`
	State         string `json:"state" validate:"required"`
	Zip           string `json:"zip" validate:"required"`
}
