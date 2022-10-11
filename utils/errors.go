package utils

import (
	"fmt"

	"github.com/vektah/gqlparser/gqlerror"
)

func ErrorResponse(err string, error_code error) error {
	return fmt.Errorf("%s: %s", err, error_code)
}

// Gqlerror return error for graphql
func Gqlerror(err string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: err,
	}
}
