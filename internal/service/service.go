package service

import (
	"fmt"
	"net/http"

	"github.com/aa-ar/budgeter-service/internal/logger"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Service struct {
	port       int
	router     *mux.Router
	middleware *negroni.Negroni
	logger     *logger.Logger
}

func NewService(port int) *Service {
	return &Service{
		port:       port,
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		logger:     logger.New(),
	}
}

func (svc *Service) Port() string {
	return fmt.Sprintf(":%d", svc.port)
}

func (svc *Service) Setup() {
	svc.setupHandlers()
	svc.setupMiddlewares()
}

func (svc *Service) Start() {
	svc.Setup()
	http.ListenAndServe(svc.Port(), svc.middleware)
}
