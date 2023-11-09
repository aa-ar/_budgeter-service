package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/domain/session/datasource"
	"github.com/aa-ar/budgeter-service/errors"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
	"github.com/sirupsen/logrus"
)

type CreateUserUsecase struct {
	sessionDataSource  datasource.SessionDataSource
	budgeterDataSource datasource.BudgeterDataSource
	presenter          createUserPresenter
}

type createUserPresenter interface {
	PrepareResponse(*model.User) *response.Response
}

func NewCreateUserUsecase(s datasource.SessionDataSource, b datasource.BudgeterDataSource, p createUserPresenter) *CreateUserUsecase {
	return &CreateUserUsecase{
		sessionDataSource:  s,
		budgeterDataSource: b,
		presenter:          p,
	}
}

func (u *CreateUserUsecase) CreateUser(sess *model.Session, pwd string) (*response.Response, error) {
	sess, err := u.sessionDataSource.FindSession(sess.ID)
	if err != nil {
		return nil, svcerror.InternalServerError{}
	}

	if sess == nil {
		return nil, errors.EmptySessionError{}
	}

	userID := sess.GetDatum("UserID")
	if userID != "" {
		return nil, errors.CookieIsAssignedToUserError{}
	}

	email := sess.GetDatum("Email")
	user, err := model.NewUser(email)
	if err != nil {
		return nil, err
	}

	if err := user.HashPassword(pwd); err != nil {
		logrus.Error(err)
		return nil, svcerror.InternalServerError{}
	}

	persistedUser, err := u.budgeterDataSource.InsertWorkspaceAndUser(user)
	if err != nil {
		logrus.Error(err)
		return nil, svcerror.InternalServerError{}
	}

	if err := u.sessionDataSource.AuthenticateSession(persistedUser.ID, sess); err != nil {
		u.sessionDataSource.ClearSessionToken(sess.ID)
	}

	return u.presenter.PrepareResponse(persistedUser), nil
}
