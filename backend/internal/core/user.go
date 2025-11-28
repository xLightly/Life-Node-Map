package core

import (
	"errors"
	"time"
)

type User struct {
	Id           string    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func NewUser(id string, username string, passwordHash string, createdAt time.Time) (*User, error) {
	if username == "" || passwordHash == "" || createdAt.IsZero() {
		return nil, errors.New("empty user entity fields")
	}
	return &User{
		Id:           id,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}, nil
}
