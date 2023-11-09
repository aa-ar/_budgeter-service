package budgeter

import (
	"github.com/Masterminds/squirrel"
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/sirupsen/logrus"
)

func (d BudgeterDataSource) InsertWorkspaceAndUser(user *model.User) (*model.User, error) {
	// Prepare workspace insert
	Workspace := model.NewWorkspace()
	workspaceSql, workspaceArgs, workspaceErr := squirrel.
		Insert("workspaces").
		Columns("id").
		Values(Workspace.ID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if workspaceErr != nil {
		logrus.Error(workspaceErr)
		return nil, workspaceErr
	}

	// Prepare user insert
	userSql, userArgs, userErr := squirrel.
		Insert("users").
		Columns("id", "workspace_id", "email", "hash").
		Values(user.ID, Workspace.ID, user.Email, user.Hash()).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if userErr != nil {
		logrus.Error(userErr)
		return nil, userErr
	}

	// Commit/rollback
	tx := d.DB.MustBegin()
	_, err := tx.Exec(workspaceSql, workspaceArgs...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	_, err = tx.Exec(userSql, userArgs...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	user.WorkspaceID = Workspace.ID
	return user, nil
}
