package core

import (
	"context"
	"errors"
	"time"
)

type Map struct {
	Id            string    `db:"id"`
	CreatorId     string    `db:"creator_id"`
	OriginalMapId string    `db:"original_map_id"`
	Title         string    `db:"title"`
	Description   string    `db:"description"`
	isPublic      bool      `db:"is_public"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func NewMap(id string, creatorId string, originMapId string, title string, description string, isPublic bool, createdAt time.Time, updatedAt time.Time) (*Map, error) {
	if creatorId == "" || title == "" || createdAt.IsZero() {
		return nil, errors.New("empty map entity fields")
	}
	return &Map{
		Id:            id,
		CreatorId:     creatorId,
		OriginalMapId: originMapId,
		Title:         title,
		Description:   description,
		isPublic:      isPublic,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}

type MapRepository interface {
	Create(ctx context.Context, m *Map) error
	GetById(ctx context.Context, id string) (*Map, error)
	GetByCreatorId(ctx context.Context, creatodId string) (*Map, error)
	GetAll(ctx context.Context) ([]Map, error)
	UpdateById(ctx context.Context, id string, newMap *Map) error
	DeleteById(ctx context.Context, id string) error
	DeleteByCreatorId(ctx context.Context, creatorId string) error
}
