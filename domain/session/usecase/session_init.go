package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type SessionInitUsecase struct {
	sessionDataSource
	budgeterDataSource
	presenter sessionInitPresenter
}

type sessionDataSource interface {
	SaveSession(*model.Session) error
}

type budgeterDataSource interface {
	FindUserByEmail(string) (*model.User, error)
}

type sessionInitPresenter interface {
	PrepareResponse(*model.Session) *response.Response
}

func NewSessionInitUsecase(s sessionDataSource, b budgeterDataSource, p sessionInitPresenter) *SessionInitUsecase {
	return &SessionInitUsecase{
		sessionDataSource:  s,
		budgeterDataSource: b,
		presenter:          p,
	}
}

func (u *SessionInitUsecase) InitializeSession(email string) (*response.Response, error) {
	data := map[string]string{
		"Email":         email,
		"Authenticated": "0",
	}

	user, err := u.budgeterDataSource.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		data["UserID"] = user.ID.String()
	}

	sess := model.NewSession().SetData(data)
	err = u.sessionDataSource.SaveSession(sess)
	if err != nil {
		return nil, err
	}

	return u.presenter.PrepareResponse(sess), nil
}
