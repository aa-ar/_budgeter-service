package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aa-ar/budgeter-service/adapter/session/errors"
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
	"github.com/sirupsen/logrus"
)

type SessionAuthHandler struct {
	usecase sessionAuthUsecase
}

type sessionAuthUsecase interface {
	AuthenticateSession(*model.Session, string) (*response.Response, error)
}

func NewSessionAuthHandler(usecase sessionAuthUsecase) *SessionAuthHandler {
	return &SessionAuthHandler{
		usecase: usecase,
	}
}

func (h *SessionAuthHandler) Path() string {
	return "/session/auth"
}

func (h *SessionAuthHandler) Methods() []string {
	return []string{http.MethodPost}
}

type sessionAuthRequest struct {
	Password string
}

func (h *SessionAuthHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	sess, err := model.NewSession().FromRequest(r)
	if err != nil {
		return err
	}

	body := json.NewDecoder(r.Body)
	var req *sessionAuthRequest
	err = body.Decode(&req)
	if err != nil {
		logrus.Error(err)
		return svcerror.BadRequestError{}
	}

	if req.Password == "" {
		return errors.MissingPasswordError{}
	}

	res, err := h.usecase.AuthenticateSession(sess, req.Password)
	if err != nil {
		return err
	}

	res.WriteTo(w)
	return nil
}
