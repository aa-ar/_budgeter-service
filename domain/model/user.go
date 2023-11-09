package model

import (
	"github.com/aa-ar/budgeter-service/errors"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const USER_PASSWORD_HASH_COST = 12

type User struct {
	ID          ksuid.KSUID `json:"id" db:"id"`
	WorkspaceID ksuid.KSUID `json:"workspace_id" db:"workspace_id"`
	Email       string      `json:"email" db:"email"`
	hash        string      `json:"-" db:"hash"`
}

func NewUser(email string) (*User, error) {
	if email == "" {
		return nil, errors.UserEmailCannotBeEmptyError{}
	}
	return &User{ID: ksuid.New(), Email: email}, nil
}

func (user *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), USER_PASSWORD_HASH_COST)
	if err == nil {
		user.hash = string(hash)
	}
	return err
}

func (user *User) Hash() string {
	return user.hash
}
