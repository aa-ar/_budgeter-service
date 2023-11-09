package datasource

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/segmentio/ksuid"
)

type SessionDataSource interface {
	FindSession(ksuid.KSUID) (*model.Session, error)
	AuthenticateSession(ksuid.KSUID, *model.Session) error
	ClearSessionToken(ksuid.KSUID) error
	SaveTempSession(*model.Session) error
}
