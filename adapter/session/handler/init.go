package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/aa-ar/budgeter-service/adapter/session/errors"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
)

type SessionInitHandler struct {
	usecase sessionInitUsecase
}

type sessionInitUsecase interface {
	InitializeSession(string) (*response.Response, error)
}

func NewSessionInitHandler(usecase sessionInitUsecase) *SessionInitHandler {
	return &SessionInitHandler{
		usecase: usecase,
	}
}

func (h *SessionInitHandler) Path() string {
	return "/session/init"
}

func (h *SessionInitHandler) Methods() []string {
	return []string{http.MethodPost}
}

type sessionInitRequest struct {
	Email string
}

func (h *SessionInitHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	body := json.NewDecoder(r.Body)
	var req *sessionInitRequest
	err := body.Decode(&req)
	if err != nil {
		log.Fatal(err)
		return svcerror.BadRequestError{}
	}

	if req.Email == "" {
		return errors.MissingEmailAddressError{}
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.InvalidEmailAddressError{}
	}

	res, err := h.usecase.InitializeSession(req.Email)
	if err != nil {
		return err
	}

	res.WriteTo(w)
	return nil
}
