package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/domain/session/datasource"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type SessionInitUsecase struct {
	sessionDataSource datasource.SessionDataSource
	budgeterDatSource datasource.BudgeterDataSource
	presenter         sessionInitPresenter
}

type sessionInitPresenter interface {
	PrepareResponse(*model.Session) *response.Response
}

func NewSessionInitUsecase(s datasource.SessionDataSource, b datasource.BudgeterDataSource, p sessionInitPresenter) *SessionInitUsecase {
	return &SessionInitUsecase{
		sessionDataSource: s,
		budgeterDatSource: b,
		presenter:         p,
	}
}

func (u *SessionInitUsecase) InitializeSession(email string) (*response.Response, error) {
	data := map[string]string{
		"Email":         email,
		"Authenticated": "0",
	}

	user, err := u.budgeterDatSource.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		data["UserID"] = user.ID.String()
	}

	sess := model.NewSession().SetData(data)
	err = u.sessionDataSource.SaveTempSession(sess)
	if err != nil {
		return nil, err
	}

	return u.presenter.PrepareResponse(sess), nil
}
