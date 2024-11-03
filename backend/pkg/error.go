package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

// WrapErr обертка для ошибок и передачи к ним контекста.
func WrapErr(e error, desc ...string) error {
	if e == nil {
		return nil
	}
	var d string
	if len(desc) > 0 {
		d = desc[0]
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("undefined call %s -> %w", d, e)
	}
	return fmt.Errorf("%s:%d %s -> %w", file, line, d, e)
}

// CheckDuplicateKey checks if the error is a postgres duplicate key violation.
func CheckDuplicateKey(err error) bool {
	var pgError *pgconn.PgError
	return errors.As(err, &pgError) && pgError.Code == "23505"
}

// CheckPageNotFound checks if the error is a fiber page not found error.
func CheckPageNotFound(err error) bool {
	var fiberError *fiber.Error
	return errors.As(err, &fiberError) && fiberError.Code == http.StatusNotFound
}

// CheckValidationError checks if the error is a validation error.
func CheckValidationError(err error) bool {
	var validationErrors validator.ValidationErrors
	return errors.As(err, &validationErrors)
}
