package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type SessionInitUsecase struct {
	saveSessionDataSource
	findUserDataSource
	presenter sessionInitPresenter
}

type saveSessionDataSource interface {
	SaveTempSession(*model.Session) error
}

type findUserDataSource interface {
	FindUserByEmail(string) (*model.User, error)
}

type sessionInitPresenter interface {
	PrepareResponse(*model.Session) *response.Response
}

func NewSessionInitUsecase(s saveSessionDataSource, b findUserDataSource, p sessionInitPresenter) *SessionInitUsecase {
	return &SessionInitUsecase{
		saveSessionDataSource: s,
		findUserDataSource:    b,
		presenter:             p,
	}
}

func (u *SessionInitUsecase) InitializeSession(email string) (*response.Response, error) {
	data := map[string]string{
		"Email":         email,
		"Authenticated": "0",
	}

	user, err := u.findUserDataSource.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		data["UserID"] = user.ID.String()
	}

	sess := model.NewSession().SetData(data)
	err = u.saveSessionDataSource.SaveTempSession(sess)
	if err != nil {
		return nil, err
	}

	return u.presenter.PrepareResponse(sess), nil
}
