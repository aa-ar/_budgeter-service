package handler

import (
	"net/http"

	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn != nil {
		var err error = fn(w, r)
		if err == nil {
			return
		}

		var e svcerror.Error = convertError(err)
		response.New(e.Status(), e, nil).WriteTo(w)
		return
	}
}
