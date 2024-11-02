package pkg

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/jackc/pgx/v5/pgconn"
)

// WrapErr обертка для ошибок и передачи к ним контекста.
func WrapErr(e error, desc ...string) error {
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
