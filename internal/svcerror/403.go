package svcerror

import "net/http"

type BadRequestError struct{}

func (err BadRequestError) Status() int {
	return http.StatusBadRequest
}

func (err BadRequestError) Error() string {
	return http.StatusText(err.Status())
}
