package errors

import "net/http"

type UserEmailCannotBeEmptyError struct{}

func (err UserEmailCannotBeEmptyError) Status() int {
	return http.StatusBadRequest
}

func (err UserEmailCannotBeEmptyError) Error() string {
	return "User email cannot be empty"
}
