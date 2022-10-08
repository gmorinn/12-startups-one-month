package utils

import (
	"fmt"
)

func ErrorResponse(err string, error_code error) error {
	return fmt.Errorf("%s: %s", err, error_code)
}
