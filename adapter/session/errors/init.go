package errors

import (
	"net/http"
)

type MissingEmailAddressError struct{}

func (err MissingEmailAddressError) Status() int {
	return http.StatusBadRequest
}

func (err MissingEmailAddressError) Error() string {
	return "Missing email"
}

type InvalidEmailAddressError struct{}

func (err InvalidEmailAddressError) Status() int {
	return http.StatusBadRequest
}

func (err InvalidEmailAddressError) Error() string {
	return "Invalid email address"
}
