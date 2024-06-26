package request

import "github.com/api/service/dbutils/schema"

type CreateUserRequest struct {
	Username        string `json:"username" validate:"min=5,max=30"`
	Password        string `json:"password"  validate:"eqfield=ConfirmPassword,min=5,max=30"`
	ConfirmPassword string `json:"confirmPassword"  validate:"min=5,max=30"`
	Email           string `json:"email" validate:"required,email"`
	Telephone       string `json:"telephone" validate:"required,tel"`
}

type ModifyUser struct {
	Username  string             `json:"username" validate:"min=5,max=30"`
	Telephone string             `json:"telephone" validate:"min=5,max=16"`
	Address   schema.UserAddress `json:"address,omitempty"  validate:"required,dive,omitempty"`
}

type ModifyUserPassword struct {
	OldPassword        string `json:"oldPassword" validate:"nefield=newPassword,min=5,max=30"`
	NewPassword        string `json:"newPassword" validate:"eqfield=ConfirmNewPassword,min=5,max=30"`
	ConfirmNewPassword string `json:"confirmNewPassword"  validate:"min=5,max=30"`
}
