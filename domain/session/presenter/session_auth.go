package presenter

import (
	"net/http"

	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type SessionAuthPresenter struct{}

func NewSessionAuthPresenter() *SessionAuthPresenter {
	return &SessionAuthPresenter{}
}

type sessionAuthResponse struct {
	OK bool `json:"ok"`
}

func (p *SessionAuthPresenter) PrepareResponse(sess *model.Session) *response.Response {
	return response.New(http.StatusOK, sessionAuthResponse{true}, nil)
}
