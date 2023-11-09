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

type EmptySessionError struct{}

func (err EmptySessionError) Status() int {
	return http.StatusUnauthorized
}

func (err EmptySessionError) Error() string {
	return "Empty session"
}

type CookieIsAssignedToUserError struct{}

func (err CookieIsAssignedToUserError) Status() int {
	return http.StatusBadRequest
}

func (err CookieIsAssignedToUserError) Error() string {
	return "The provided session cookie is already assigned to a user"
}

type InvalidSessionProvidedError struct{}

func (err InvalidSessionProvidedError) Status() int {
	return http.StatusBadRequest
}

func (err InvalidSessionProvidedError) Error() string {
	return "The provided session cookie is invalid"
}

type InvalidPasswordError struct{}

func (err InvalidPasswordError) Status() int {
	return http.StatusBadRequest
}

func (err InvalidPasswordError) Error() string {
	return "The provided password is invalid for the given session"
}
