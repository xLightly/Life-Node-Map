package core

type role string

type MapContributors struct {
	MapId  string `db:"map_id"`
	UserId string `db:"user_id"`
	Role   role   `db:"role"`
}
