package handler

import "github.com/aa-ar/budgeter-service/internal/svcerror"

func convertError(err error) svcerror.Error {
	e, ok := err.(svcerror.Error)
	if !ok {
		return svcerror.InternalServerError{}
	}
	return e
}
