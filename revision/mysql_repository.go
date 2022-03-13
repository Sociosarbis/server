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

func NewMysqlRepo(q *query.Query, log *zap.Logger) (domain.RevisionRepo, error) {
	return mysqlRepo{
		q,
		log,
	}, nil
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
) ([]*model.Revision, error) {
	revisions, err := r.q.RevisionHistory.WithContext(ctx).
		Where(r.q.RevisionHistory.Type.In(model.PersonRevisionTypes...)).
		Limit(limit).
		Offset(offset).Find()
	if err != nil {
		return nil, errgo.Wrap(err, "dal")
	}

	result := make([]*model.Revision, len(revisions))
	for i, revision := range revisions {
		result[i] = convertRevisionDao(revision, nil)
	}
	return result, nil
}

func (r mysqlRepo) GetPersonRelated(ctx context.Context, id domain.IDType) (*model.Revision, error) {
	var exitError error
	revision, err := r.q.RevisionHistory.WithContext(ctx).
		Where(r.q.RevisionHistory.ID.Eq(id),
			r.q.RevisionHistory.Type.In(model.PersonRevisionTypes...)).
		First()

	if err == nil {
		data, err := r.q.RevisionText.WithContext(ctx).
			Where(r.q.RevisionText.TextID.Eq(revision.TextID)).First()
		if err == nil {
			return convertRevisionDao(revision, data), nil
		}
		exitError = err
	} else {
		exitError = err
	}
	r.log.Error("unexpected error happened", zap.Error(exitError))
	return &model.Revision{}, errgo.Wrap(exitError, "dal")
}

func convertRevisionDao(r *dao.RevisionHistory, data *dao.RevisionText) *model.Revision {
	var text dao.GzipPhpSerializedBlob
	if data != nil {
		text = data.Text
	}

	return &model.Revision{
		ID:        r.ID,
		Type:      r.Type,
		Summary:   r.Summary,
		CreatorID: r.CreatorID,
		CreatedAt: time.Unix(int64(r.CreatedAt), 0),
		Data:      text,
	}
}
