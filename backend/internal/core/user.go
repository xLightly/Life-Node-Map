package core

import (
	"context"
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

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	UpdateById(ctx context.Context, id string, newUser *User) error
	DeleteById(ctx context.Context, id string) error
	DeleteByUsername(ctx context.Context, username string) error
}
