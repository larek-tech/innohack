package model

type SignupReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

type TokenResp struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}
