package service

import (
	"net/http"

	"github.com/aa-ar/budgeter-service/internal/handler"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
)

type Handler interface {
	Path() string
	Methods() []string
	Handler(w http.ResponseWriter, r *http.Request) error
}

func (svc *Service) setupHandlers() {
	svc.router.MethodNotAllowedHandler = handler.Handler(methodNotAllowedHandler)
	svc.router.NotFoundHandler = handler.Handler(notFoundHandler)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) error {
	return svcerror.MethodNotAllowedError{}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) error {
	return svcerror.NotFoundError{}
}
