package model

import (
	"net/http"
	"time"

	"github.com/aa-ar/budgeter-service/errors"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

const COOKIE_NAME = "session"

type Session struct {
	ID     ksuid.KSUID
	Data   map[string]string
	Expiry time.Duration
}

func NewSession() *Session {
	sessionID := ksuid.New()
	expiry := time.Until(time.Now().AddDate(0, 1, 0))
	return &Session{
		ID:     sessionID,
		Expiry: expiry,
	}
}

func (sess *Session) SetData(data map[string]string) *Session {
	sess.Data = data
	return sess
}

func (sess *Session) GetDatum(key string) string {
	return sess.Data[key]
}

func (sess *Session) FromRequest(req *http.Request) (*Session, error) {
	cookie, err := req.Cookie(COOKIE_NAME)
	if err != nil {
		logrus.Error(err)
		return nil, errors.NoSessionCookieError{}
	}
	sessionID, err := ksuid.Parse(cookie.Value)
	if err != nil {
		logrus.Error(err)
		return sess, errors.BadSessionTokenError{}
	}
	sess.ID = sessionID
	sess.Expiry = time.Until(time.Now())
	return sess, nil
}

func (sess *Session) ToCookie() *http.Cookie {
	return &http.Cookie{
		Name:   COOKIE_NAME,
		Value:  sess.ID.String(),
		MaxAge: int(sess.Expiry.Seconds()),
	}
}
