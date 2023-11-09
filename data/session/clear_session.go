package session

import (
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

func (d SessionDataSource) ClearSessionToken(sessionID ksuid.KSUID) error {
	if err := d.Client.Del(d.ctx, sessionID.String()).Err(); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
