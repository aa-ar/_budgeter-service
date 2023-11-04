package svcerror

import "net/http"

type NotFoundError struct{}

func (err NotFoundError) Status() int {
	return http.StatusNotFound
}

func (err NotFoundError) Error() string {
	return "Resource not found"
}
