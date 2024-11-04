package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

var (
	errMap = map[error]ErrorResponse{
		shared.ErrEmailAlreadyTaken: {
			Msg:    shared.ErrEmailAlreadyTaken.Error(),
			Status: fiber.StatusBadRequest,
		},
		shared.ErrInvalidCredentials: {
			Status: fiber.StatusUnauthorized,
		},
		shared.ErrStorageInternal: {
			Status: fiber.StatusInternalServerError,
		},
		shared.ErrMissingJwt: {
			Status: fiber.StatusUnauthorized,
		},
		shared.ErrInvalidJwt: {
			Status: fiber.StatusUnauthorized,
		},
		shared.ErrNoAccessToSession: {
			Status: fiber.StatusForbidden,
		},
	}
)
