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
	return http.StatusInternalServerError
}

func (err BadSessionTokenError) Error() string {
	return "Bad session token"
}
