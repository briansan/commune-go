package errors

import (
	"fmt"
	"net/http"
)

type UserExistsErr struct {
	Username string
}

func (err UserExistsErr) Error() string {
	return fmt.Sprintf("username %v already exists", err.Username)
}

func (err UserExistsErr) StatusCode() int {
	return http.StatusConflict
}

type MongoErr struct {
	Err error
}

func (err MongoErr) Error() string {
	return err.Err.Error()
}

func (err MongoErr) StatusCode() int {
	return http.StatusInternalServerError
}
