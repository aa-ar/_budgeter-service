package budgeter

import (
	"github.com/Masterminds/squirrel"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userHash struct {
	Hash string `db:"hash"`
}

func (d BudgeterDataSource) CompareHashesForUser(userID ksuid.KSUID, pwd string) error {
	var hash userHash
	sql, args, err := squirrel.
		Select("hash").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logrus.Error(err)
		return err
	}
	rows, err := d.DB.Queryx(sql, args...)
	if err != nil {
		logrus.Error(err)
		return err
	}
	for rows.Next() {
		err := rows.StructScan(&hash)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	defer rows.Close()
	return bcrypt.CompareHashAndPassword([]byte(hash.Hash), []byte(pwd))
}
