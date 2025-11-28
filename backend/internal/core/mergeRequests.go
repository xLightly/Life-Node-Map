package core

import (
	"errors"
	"time"
)

type requestStatus string

type MergeRequest struct {
	Id          string        `db:"id"`
	CreatorId   string        `db:"creator_id"`
	SourceMapId string        `db:"source_map_id"`
	TargetMapId string        `db:"traget_map_id"`
	Status      requestStatus `db:"status"`
	CreatedAt   time.Time     `db:"created_at"`
}

func NewMergeRequest(id string, creatorId string, sourceMapId string, targerMapId string, status requestStatus, createdAt time.Time) (*MergeRequest, error) {
	if creatorId == "" || sourceMapId == "" || targerMapId == "" || status == "" || createdAt.IsZero() {
		return nil, errors.New("empty merge request entity fields")
	}
	return &MergeRequest{
		Id:          id,
		CreatorId:   creatorId,
		SourceMapId: sourceMapId,
		TargetMapId: targerMapId,
		Status:      status,
		CreatedAt:   createdAt,
	}, nil
}
