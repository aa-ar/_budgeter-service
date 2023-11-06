package budgeter

import (
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/aa-ar/budgeter-service/domain/model"
)

func (d BudgeterDataSource) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	sql, args, err := squirrel.
		Select(
			"id",
			"workspace_id",
			"email",
		).
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := d.DB.Queryx(sql, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	defer rows.Close()
	if user.ID.IsNil() {
		return nil, nil
	}
	return &user, nil
}
