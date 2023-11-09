package usecase

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/errors"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

type SessionAuthUsecase struct {
	sessionAuthenticatorDataSource
	hashComparerDataSource
	presenter sessionAuthPresenter
}

type sessionAuthenticatorDataSource interface {
	FindSession(ksuid.KSUID) (*model.Session, error)
	AuthenticateSession(ksuid.KSUID, *model.Session) error
}

type hashComparerDataSource interface {
	CompareHashesForUser(ksuid.KSUID, string) error
}

type sessionAuthPresenter interface {
	PrepareResponse(*model.Session) *response.Response
}

func NewSessionAuthUsecase(s sessionAuthenticatorDataSource, b hashComparerDataSource, p sessionAuthPresenter) *SessionAuthUsecase {
	return &SessionAuthUsecase{
		sessionAuthenticatorDataSource: s,
		hashComparerDataSource:         b,
		presenter:                      p,
	}
}

func (u *SessionAuthUsecase) AuthenticateSession(sess *model.Session, pwd string) (*response.Response, error) {
	sess, err := u.sessionAuthenticatorDataSource.FindSession(sess.ID)
	if err != nil {
		return nil, svcerror.InternalServerError{}
	}

	if sess == nil {
		return nil, errors.EmptySessionError{}
	}

	rawUserID := sess.Data["UserID"]
	if rawUserID == "" {
		return nil, errors.InvalidSessionProvidedError{}
	}

	userID, err := ksuid.Parse(rawUserID)
	if err != nil {
		return nil, svcerror.InternalServerError{}
	}

	if err = u.hashComparerDataSource.CompareHashesForUser(userID, pwd); err != nil {
		return nil, errors.InvalidPasswordError{}
	}

	if err := u.sessionAuthenticatorDataSource.AuthenticateSession(userID, sess); err != nil {
		logrus.Error(err)
		return nil, svcerror.InternalServerError{}
	}

	return u.presenter.PrepareResponse(sess), nil
}
