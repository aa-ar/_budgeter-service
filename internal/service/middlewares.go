package service

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func (svc *Service) setupMiddlewares() {
	svc.middleware.Use(cors.Default())
	svc.middleware.Use(negroni.HandlerFunc(jsonResponse))
	svc.middleware.UseHandler(svc.router)
}

func jsonResponse(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}
