package model

type SignUpReq struct {
	FirstName       string `validate:"required"`
	LastName        string `validate:"required"`
	Email           string `validate:"required,email" form:"email"`
	Password        string `validate:"required" form:"password"`
	PasswordConfirm string `validate:"required" form:"confirm-password"`
	Errors          map[string]error
}

type LoginReq struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required" form:"password"`
}

type EmailLoginData struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
