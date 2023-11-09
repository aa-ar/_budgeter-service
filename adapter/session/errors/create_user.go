package errors

import (
	"net/http"
)

type MissingOrInvalidSessionError struct{}

func (err MissingOrInvalidSessionError) Status() int {
	return http.StatusBadRequest
}

func (err MissingOrInvalidSessionError) Error() string {
	return "Missing or invalid session"
}

type MissingPasswordError struct{}

func (err MissingPasswordError) Status() int {
	return http.StatusBadRequest
}

func (err MissingPasswordError) Error() string {
	return "Missing password"
}
