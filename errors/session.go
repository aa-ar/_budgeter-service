package errors

import (
	"net/http"
)

type NoSessionCookieError struct{}

func (err NoSessionCookieError) Status() int {
	return http.StatusBadRequest
}

func (err NoSessionCookieError) Error() string {
	return "Missing session cookie"
}

type BadSessionTokenError struct{}

func (err BadSessionTokenError) Status() int {
	return http.StatusBadRequest
}

func (err BadSessionTokenError) Error() string {
	return "Bad session token"
}

type CookieIsAssignedToUserError struct{}

func (err CookieIsAssignedToUserError) Status() int {
	return http.StatusBadRequest
}

func (err CookieIsAssignedToUserError) Error() string {
	return "The provided session cookie is already assigned to a user"
}
