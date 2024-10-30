package pkg

import (
	"fmt"
	"runtime"
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
