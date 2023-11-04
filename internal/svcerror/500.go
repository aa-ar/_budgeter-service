package svcerror

import "net/http"

type InternalServerError struct {
	details string
}

func (err InternalServerError) Status() int {
	return http.StatusInternalServerError
}

func (err InternalServerError) Error() string {
	if err.details != "" {
		return err.details
	}
	return "Internal server error"
}
