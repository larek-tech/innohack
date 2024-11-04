package shared

import "errors"

// 400
var (
	ErrEmailAlreadyTaken = errors.New("email is already taken")
)

// 401
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidJwt         = errors.New("invalid jwt token")
	ErrMissingJwt         = errors.New("missing jwt")
)

// 403
var (
	ErrNoAccessToSession = errors.New("no access to session")
)

// 500
var (
	ErrStorageInternal = errors.New("internal storage error")
)

// system
var (
	// ErrDuplicateKey is an error for postgres unique key violation.
	ErrDuplicateKey = errors.New("duplicate key")
)
