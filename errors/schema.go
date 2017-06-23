package errors

import (
	"fmt"
	"net/http"
)

type ValidationErr struct {
	HTTPError
	Field string
	Type  string
}

func (err ValidationErr) Error() string {
	return fmt.Sprintf("%v field is required as %v", err.Field, err.Type)
}

func (err ValidationErr) StatusCode() int {
	return http.StatusBadRequest
}
