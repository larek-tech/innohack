package model

type SignupReq struct {
	Email           string `validate:"required,email" form:"email"`
	Password        string `validate:"required" form:"password"`
	PasswordConfirm string `validate:"required" form:"confirm-password"`
	Errors          map[string]error
}

type LoginReq struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required" form:"password"`
}
