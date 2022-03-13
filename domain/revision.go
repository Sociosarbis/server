package domain

import (
	"context"

	"github.com/bangumi/server/model"
)

type RevisionRepo interface {
	CountPersonRelated(ctx context.Context, personID PersonIDType) (int64, error)

	ListPersonRelated(
		ctx context.Context, personID PersonIDType, limit int, offset int,
	) ([]*model.Revision, error)

	GetPersonRelated(ctx context.Context, id IDType) (*model.Revision, error)
}
