package session

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

func (d SessionDataSource) FindSession(sessionID ksuid.KSUID) (*model.Session, error) {
	cmd := d.Client.HGetAll(d.ctx, sessionID.String())
	if err := cmd.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	res, err := cmd.Result()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &model.Session{
		ID:   sessionID,
		Data: res,
	}, nil
}
