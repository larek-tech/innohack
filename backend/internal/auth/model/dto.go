package model

type SignUpReq struct {
	Email           string `validate:"required,email" form:"email"`
	Password        string `validate:"required" form:"password"`
	PasswordConfirm string `validate:"required" form:"confirm-password"`
}

type LoginReq struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required" form:"password"`
}
