package revision

import (
	"context"
	"time"

	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/dal/dao"
	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/model"
	"go.uber.org/zap"
)

type mysqlRepo struct {
	q   *query.Query
	log *zap.Logger
}

func NewMysqlRepo(q *query.Query, log *zap.Logger) domain.RevisionRepo {
	return mysqlRepo{
		q,
		log,
	}
}

func (r mysqlRepo) CountPersonRelated(ctx context.Context, id domain.PersonIDType) (int64, error) {
	c, err := r.q.RevisionHistory.WithContext(ctx).Where(r.q.RevisionHistory.Type.In(model.PersonRevisionTypes...)).Count()
	if err != nil {
		return 0, errgo.Wrap(err, "dal")
	}
	return c, nil
}

func (r mysqlRepo) ListPersonRelated(
	ctx context.Context, personID domain.PersonIDType, limit int, offset int,
) ([]model.Revision, error) {
	revisions, err := r.q.RevisionHistory.WithContext(ctx).
		Preload(r.q.RevisionHistory.Creator).
		Where(r.q.RevisionHistory.Type.In(model.PersonRevisionTypes...)).
		Limit(limit).
		Offset(offset).Find()
	if err != nil {
		return nil, errgo.Wrap(err, "dal")
	}

	result := make([]model.Revision, len(revisions))
	for i, revision := range revisions {
		result[i] = convertRevisionDao(revision)
	}
	return result, nil
}

func (r mysqlRepo) GetPersonRelated(ctx context.Context, id domain.IDType) (model.Revision, error) {
	revision, err := r.q.RevisionHistory.WithContext(ctx).
		Preload(r.q.RevisionHistory.Data).
		Preload(r.q.RevisionHistory.Creator).
		Where(r.q.RevisionHistory.ID.Eq(id),
			r.q.RevisionHistory.Type.In(model.PersonRevisionTypes...)).
		First()
	if err != nil {
		r.log.Error("unexpected error happened", zap.Error(err))

		return model.Revision{}, errgo.Wrap(err, "dal")
	}
	return convertRevisionDao(revision), nil
}

func convertRevisionDao(r *dao.RevisionHistory) model.Revision {
	var data []interface{}
	if r.Data != nil {
		data = r.Data.Text
	}
	return model.Revision{
		ID:      r.ID,
		Type:    r.Type,
		Summary: r.Summary,
		Creator: model.Creator{
			Username: r.Creator.Username,
			Nickname: r.Creator.Nickname,
		},
		CreatedAt: time.Unix(int64(r.CreatedAt), 0),
		Data:      data,
	}
}
