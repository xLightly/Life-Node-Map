package core

import (
	"context"
	"errors"
)

type path string

type Node struct {
	Id         string `db:"id"`
	MapId      string `db:"map_id"`
	Ltree      path   `db:"path"`
	Title      string `db:"title"`
	Properties []byte `db:"properties"`
}

func NewNode(id string, mapId string, ltree path, title string, props []byte) (*Node, error) {
	if mapId == "" || ltree == "" || title == "" || props == nil {
		return nil, errors.New("empty node entity fields")
	}
	return &Node{
		Id:         id,
		MapId:      mapId,
		Ltree:      ltree,
		Title:      title,
		Properties: props,
	}, nil
}

type NodeRepository interface {
	Create(ctx context.Context, n *Node) error
	GetById(ctx context.Context, id string) (*Node, error)
	UpdateById(ctx context.Context, id string, newLtree path, newTitle string, newProperties []byte) error
	DeleteById(ctx context.Context, id string) error
}
