package service

import (
	"fmt"
	"net/http"

	"github.com/aa-ar/budgeter-service/internal/handler"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Service struct {
	port       int
	router     *mux.Router
	middleware *negroni.Negroni
}

func NewService(port int) *Service {
	return &Service{
		port:       port,
		router:     mux.NewRouter(),
		middleware: negroni.New(),
	}
}

func (svc *Service) Port() string {
	return fmt.Sprintf(":%d", svc.port)
}

func (svc *Service) AttachRegistries(registries []Registry) *Service {
	for _, registry := range registries {
		for _, h := range registry.Handlers() {
			svc.router.Handle(h.Path(), handler.Handler(h.Handler)).Methods(h.Methods()...)
		}
	}
	return svc
}

func (svc *Service) setup() {
	svc.setupHandlers()
	svc.setupMiddlewares()
}

func (svc *Service) Start() {
	svc.setup()
	http.ListenAndServe(svc.Port(), svc.middleware)
}
