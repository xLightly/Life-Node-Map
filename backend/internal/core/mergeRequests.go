package core

import "time"

type requestStatus string

type MergeRequest struct {
	Id          string        `db:"id"`
	CreatorId   string        `db:"creator_id"`
	SourceMapId string        `db:"source_map_id"`
	TargetMapId string        `db:"traget_map_id"`
	Status      requestStatus `db:"status"`
	CreatedAt   time.Time     `db:"created_at"`
}
