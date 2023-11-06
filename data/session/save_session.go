package session

import "github.com/aa-ar/budgeter-service/domain/model"

func (d SessionDataSource) SaveSession(sess *model.Session) error {
	err := d.Client.HSet(d.ctx, sess.ID.String(), sess.Data).Err()
	if err != nil {
		return err
	}
	d.Client.Expire(d.ctx, sess.ID.String(), sess.Expiry).Err()
	if err != nil {
		return err
	}
	return nil
}
