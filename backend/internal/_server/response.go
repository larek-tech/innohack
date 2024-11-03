package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/rs/zerolog/log"
)

var (
	healtCheckHandler = func(nodeID string) fiber.Handler {
		return func(c *fiber.Ctx) error {
			return c.JSON(healthCheckInfo{nodeID})
		}
	}
)

type healthCheckInfo struct {
	NodeID string `json:"node_id"`
}

// ErrorResponse is a struct that holds the error message and status code.
type ErrorResponse struct {
	Msg    string `json:"msg"`
	Status int    `json:"-"`
}

// ErrorHandler is a struct that holds the error status map.
type ErrorHandler struct {
	status map[error]ErrorResponse
}

// NewErrorHandler creates a new ErrorHandler instance with the given error status map.
func NewErrorHandler(errStatus map[error]ErrorResponse) ErrorHandler {
	status := map[error]ErrorResponse{
		pgx.ErrNoRows: {
			Msg:    "no rows found",
			Status: http.StatusNotFound,
		},
		fiber.ErrUnprocessableEntity: {
			Msg:    "validation error",
			Status: http.StatusUnprocessableEntity,
		},
	}

	for k, v := range errStatus {
		status[k] = v
	}

	return ErrorHandler{
		status: status,
	}
}

// Handler is a method that handles the error and returns a JSON response.
// Should be used as a fiber.Config.ErrorHandler.
func (h ErrorHandler) Handler(ctx *fiber.Ctx, err error) error {
	e := h.getErrorResponse(err)
	log.Err(err).Msg(e.Msg)
	return ctx.Status(e.Status).JSON(e) //nolint:wrapcheck // no need to wrap
}

func (h ErrorHandler) getErrorResponse(err error) ErrorResponse {
	var (
		ok bool
		e  ErrorResponse
	)

	if pkg.CheckPageNotFound(err) {
		return ErrorResponse{
			Msg:    "page not found",
			Status: fiber.StatusNotFound,
		}
	}

	if pkg.CheckDuplicateKey(err) {
		return ErrorResponse{
			Msg:    "duplicate key",
			Status: fiber.StatusBadRequest,
		}
	}

	if pkg.CheckValidationError(err) {
		return ErrorResponse{
			Msg:    err.Error(),
			Status: fiber.StatusUnprocessableEntity,
		}
	}

	e, ok = h.status[err]
	if !ok {
		e = ErrorResponse{
			Msg:    "unknown error",
			Status: fiber.StatusInternalServerError,
		}
	}
	if e.Msg == "" {
		e.Msg = err.Error()
	} else {
		e.Msg = fmt.Sprintf("%s %v", e.Msg, err)
	}

	return e
}
