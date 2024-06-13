package request

type CreateUserRequest struct {
	UserName  string `db:"username" json:"username" validate:"required,lte=255"`
	Password  string `db:"password" json:"password"  validate:"min=5,max=30"`
	Email     string `db:"email" json:"email" validate:"required,email"`
	Telephone string `db:"telephone" json:"telephone" validate:"min=5,max=16"`
}
