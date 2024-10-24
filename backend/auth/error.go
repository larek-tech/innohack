package auth

import "errors"

var (
	ErrPasswordMissMatch  = errors.New("passwords don't match")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrPasswordMismatch   = errors.New("passwords are mismatched")
	ErrEmailAlreadyTaken  = errors.New("email is already taken")
)
