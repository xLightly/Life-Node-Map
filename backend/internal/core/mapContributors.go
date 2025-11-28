package core

import "errors"

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
