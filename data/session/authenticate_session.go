package session

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

func (d SessionDataSource) AuthenticateSession(userID ksuid.KSUID, sess *model.Session) error {
	if sess.Data["UserID"] == "" {
		sess.Data["UserID"] = userID.String()
	}
	sess.Data["Authenticated"] = "1"
	if err := d.Client.HSet(d.ctx, sess.ID.String(), sess.Data).Err(); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
