package auth

import "errors"

type EmailUserData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegistrationInput struct {
	Email           string `validate:"required,email" form:"email"`
	Password        string `validate:"required" form:"password"`
	PasswordConfirm string `validate:"required" form:"confirm-password"`
	Errors          map[string]error
}

type LoginInput struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required" form:"password"`
}

type FormData struct {
	Data  RegistrationInput
	Error map[string]error
}

func newFormData(data RegistrationInput, errs []error) FormData {
	errorsMap := make(map[string]error)
	for _, err := range errs {
		if errors.Is(err, ErrPasswordMissMatch) {
			errorsMap["password"] = err
		}
	}

	return FormData{
		Data:  data,
		Error: errorsMap,
	}
}
