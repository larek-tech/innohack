package model

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// EmailUserData - данные для авторизации через email
type EmailUserData struct {
	UserID   int64
	Email    string
	Password string
}

type EmailRegisterData struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
}
