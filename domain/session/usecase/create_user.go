package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/errors"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

type CreateUserUsecase struct {
	findSessionDataSource
	createUserDataSource
	presenter createUserPresenter
}

type findSessionDataSource interface {
	FindSession(ksuid.KSUID) (*model.Session, error)
	AuthenticateSession(ksuid.KSUID, *model.Session) error
	ClearSessionToken(ksuid.KSUID) error
}

type createUserDataSource interface {
	InsertWorkspaceAndUser(*model.User) (*model.User, error)
}

type createUserPresenter interface {
	PrepareResponse(*model.User) *response.Response
}

func NewCreateUserUsecase(s findSessionDataSource, b createUserDataSource, p createUserPresenter) *CreateUserUsecase {
	return &CreateUserUsecase{
		findSessionDataSource: s,
		createUserDataSource:  b,
		presenter:             p,
	}
}

func (u *CreateUserUsecase) CreateUser(sess *model.Session, pwd string) (*response.Response, error) {
	sess, err := u.findSessionDataSource.FindSession(sess.ID)
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

	persistedUser, err := u.createUserDataSource.InsertWorkspaceAndUser(user)
	if err != nil {
		logrus.Error(err)
		return nil, svcerror.InternalServerError{}
	}

	if err := u.findSessionDataSource.AuthenticateSession(persistedUser.ID, sess); err != nil {
		u.findSessionDataSource.ClearSessionToken(sess.ID)
	}

	return u.presenter.PrepareResponse(persistedUser), nil
}
