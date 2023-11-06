package presenter

import (
	"net/http"

	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type SessionInitPresenter struct{}

func NewSessionInitPresenter() *SessionInitPresenter {
	return &SessionInitPresenter{}
}

type sessionInitResponse struct {
	OK bool `json:"ok"`
}

func (p *SessionInitPresenter) PrepareResponse(sess *model.Session) *response.Response {
	cookie := sess.ToCookie()
	return response.New(http.StatusOK, sessionInitResponse{true}, []*http.Cookie{cookie})
}
