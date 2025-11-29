package core

import (
	"context"
	"errors"
)

type contributorRole string

type MapContributors struct {
	MapId  string          `db:"map_id"`
	UserId string          `db:"user_id"`
	Role   contributorRole `db:"role"`
}

func NewMapContibutors(mapId string, userId string, role contributorRole) (*MapContributors, error) {
	if mapId == "" || userId == "" || role == "" {
		return nil, errors.New("empty map contributors entity fields")
	}
	return &MapContributors{
		MapId:  mapId,
		UserId: userId,
		Role:   role,
	}, nil
}

type MapContributorsRepository interface {
	Create(ctx context.Context, mp *MapContributors) error
	GetById(ctx context.Context, id string) (*MapContributors, error)
	UpdateById(ctx context.Context, role contributorRole) error
	DeleteById(ctx context.Context, id string) error
}
