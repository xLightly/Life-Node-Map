package core

type path string

type Node struct {
	Id         string `db:"id"`
	MapId      string `db:"map_id"`
	Ltree      path   `db:"path"`
	Title      string `db:"title"`
	Properties []byte `db:"properties"`
}
