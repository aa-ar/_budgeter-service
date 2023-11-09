package datasource

import (
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/segmentio/ksuid"
)

type BudgeterDataSource interface {
	CompareHashesForUser(ksuid.KSUID, string) error
	InsertWorkspaceAndUser(*model.User) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
}
