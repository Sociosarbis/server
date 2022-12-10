// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package pm

import (
	"context"
	"database/sql/driver"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gen"

	"github.com/bangumi/server/internal/dal/dao"
	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/pkg/generic/slice"
)

const recentContactLimit = 15

var errUserIrrelevantMessageType = errors.New("has user irrelevant message")
var errRelatedPrivateMessageNotExists = errors.New("related private message not exists")
var errInvalidOperation = errors.New("invalid operation")

type mysqlRepo struct {
	q   *query.Query
	log *zap.Logger
}

func NewMysqlRepo(q *query.Query, log *zap.Logger) (domain.PrivateMessageRepo, error) {
	return mysqlRepo{q: q, log: log.Named("pm.mysqlRepo")}, nil
}

func (r mysqlRepo) List(
	ctx context.Context,
	userID model.UserID,
	folder model.PrivateMessageFolderType,
	offset int,
	limit int,
) ([]model.PrivateMessageListItem, error) {
	var conds []gen.Condition
	do := r.q.PrivateMessage.WithContext(ctx)
	if folder == model.PrivateMessageFolderTypeInbox {
		conds = []gen.Condition{
			r.q.PrivateMessage.ReceiverID.Eq(userID),
			r.q.PrivateMessage.DeletedByReceiver.Is(false),
		}
	} else {
		conds = []gen.Condition{
			r.q.PrivateMessage.SenderID.Eq(userID),
			r.q.PrivateMessage.DeletedBySender.Is(false),
		}
	}
	ret, err := do.
		Where(
			conds...,
		).Order(r.q.PrivateMessage.ID.Desc()).Offset(offset).Limit(limit).Find()

	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return make([]model.PrivateMessageListItem, 0), errgo.Wrap(err, "dal")
	}

	mainIDs := slice.Unique(slice.Map(ret, func(v *dao.PrivateMessage) model.PrivateMessageID {
		return v.RelatedMessageID
	}))

	mainMsgList, err := do.Where(r.q.PrivateMessage.ID.In(slice.ToValuer(mainIDs)...)).Find()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return make([]model.PrivateMessageListItem, 0), errgo.Wrap(err, "dal")
	}
	mainMsgs := slice.ToMap(mainMsgList, func(v *dao.PrivateMessage) model.PrivateMessageID {
		return v.ID
	})
	return slice.Map(ret, func(v *dao.PrivateMessage) model.PrivateMessageListItem {
		return model.PrivateMessageListItem{
			Main: convertDaoToModel(mainMsgs[v.RelatedMessageID]),
			Self: convertDaoToModel(v),
		}
	}), nil
}

func (r mysqlRepo) CountByFolder(ctx context.Context,
	userID model.UserID,
	folder model.PrivateMessageFolderType) (int64, error) {
	var conds []gen.Condition
	do := r.q.PrivateMessage.WithContext(ctx)
	if folder == model.PrivateMessageFolderTypeInbox {
		conds = []gen.Condition{
			r.q.PrivateMessage.ReceiverID.Eq(userID),
			r.q.PrivateMessage.DeletedByReceiver.Is(false),
		}
	} else {
		conds = []gen.Condition{
			r.q.PrivateMessage.SenderID.Eq(userID),
			r.q.PrivateMessage.DeletedBySender.Is(false),
		}
	}
	count, err := do.Where(conds...).Count()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return 0, errgo.Wrap(err, "dal")
	}
	return count, nil
}

func (r mysqlRepo) ListRelated(
	ctx context.Context,
	userID model.UserID,
	id model.PrivateMessageID,
) ([]model.PrivateMessage, error) {
	do := r.q.PrivateMessage.WithContext(ctx)
	firstMsg, err := do.Where(r.q.PrivateMessage.ID.Eq(id)).First()
	var emptyMsgList = make([]model.PrivateMessage, 0)
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return emptyMsgList, errgo.Wrap(err, "dal")
	}
	if firstMsg == nil {
		return emptyMsgList, domain.ErrNotFound
	}
	if firstMsg.SenderID != userID && firstMsg.ReceiverID != userID {
		return emptyMsgList, domain.ErrPrivateMessageNotOwned
	}
	if (firstMsg.SenderID == userID && firstMsg.DeletedBySender) ||
		(firstMsg.ReceiverID == userID && firstMsg.DeletedByReceiver) {
		return emptyMsgList, domain.ErrPrivateMessageDeleted
	}
	res, err := do.Where(
		r.q.PrivateMessage.RelatedMessageID.Eq(id),
		do.Where(
			do.
				Where(
					r.q.PrivateMessage.SenderID.Eq(userID),
					r.q.PrivateMessage.DeletedBySender.Is(false),
				),
		).
			Or(
				r.q.PrivateMessage.ReceiverID.Eq(userID),
				r.q.PrivateMessage.DeletedByReceiver.Is(false),
			),
	).
		Find()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return emptyMsgList, errgo.Wrap(err, "dal")
	}
	return slice.Map(res, convertDaoToModel), nil
}

func (r mysqlRepo) CountTypes(
	ctx context.Context,
	userID model.UserID,
) (model.PrivateMessageTypeCounts, error) {
	do := r.q.PrivateMessage.WithContext(ctx)
	res := model.PrivateMessageTypeCounts{}
	c1, err := r.CountByFolder(ctx, userID, model.PrivateMessageFolderTypeOutbox)
	if err != nil {
		return res, err
	}
	c2, err := r.CountByFolder(ctx, userID, model.PrivateMessageFolderTypeInbox)
	if err != nil {
		return res, err
	}
	c3, err := do.Where(
		r.q.PrivateMessage.ReceiverID.Eq(userID),
		r.q.PrivateMessage.DeletedBySender.Is(false),
		r.q.PrivateMessage.New.Is(true)).
		Count()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return res, errgo.Wrap(err, "dal")
	}
	res.Outbox = c1
	res.Inbox = c2
	res.Unread = c3
	return res, nil
}

func (r mysqlRepo) ListRecentContact(
	ctx context.Context,
	userID model.UserID,
) ([]model.UserID, error) {
	res, err := r.q.PrivateMessage.
		WithContext(ctx).
		Select(r.q.PrivateMessage.ReceiverID).
		Where(r.q.PrivateMessage.SenderID.Eq(userID)).
		Order(r.q.PrivateMessage.CreatedTime.Desc()).
		Group(r.q.PrivateMessage.ReceiverID).
		Limit(recentContactLimit).
		Find()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return make([]model.UserID, 0), errgo.Wrap(err, "dal")
	}
	return slice.Map(res, func(v *dao.PrivateMessage) model.UserID {
		return v.ReceiverID
	}), nil
}

func (r mysqlRepo) MarkRead(ctx context.Context, userID model.UserID, relatedID model.PrivateMessageID) error {
	rows, err := r.q.PrivateMessage.WithContext(ctx).
		Where(
			r.q.PrivateMessage.RelatedMessageID.Eq(relatedID),
			r.q.PrivateMessage.ReceiverID.Eq(userID),
			r.q.PrivateMessage.New.Is(true)).
		Update(r.q.PrivateMessage.New, false)

	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return errgo.Wrap(err, "dal")
	}
	if rows.RowsAffected == 0 {
		return errInvalidOperation
	}
	return nil
}

func (r mysqlRepo) Create(
	ctx context.Context,
	senderID model.UserID,
	receiverIDs []model.UserID,
	relatedIDFilter domain.PrivateMessageIDFilter,
	title string,
	content string,
) ([]model.PrivateMessage, error) {
	emptyList := make([]model.PrivateMessage, 0)
	if relatedIDFilter.Type.Set {
		if len(receiverIDs) > 1 {
			return emptyList, errRelatedPrivateMessageNotExists
		}
		msg, err := r.q.WithContext(ctx).
			PrivateMessage.
			Where(r.q.PrivateMessage.ID.Eq(relatedIDFilter.Type.Value)).First()
		if err != nil ||
			(msg.SenderID != senderID && msg.SenderID != receiverIDs[0]) ||
			(msg.ReceiverID != senderID && msg.ReceiverID != receiverIDs[0]) {
			r.log.Error("unexpected error", zap.Error(err))
			return emptyList, errRelatedPrivateMessageNotExists
		}
	}
	msgs := make([]*dao.PrivateMessage, len(receiverIDs))
	for i := range msgs {
		msgs[i] = &dao.PrivateMessage{
			SenderID:    senderID,
			ReceiverID:  receiverIDs[i],
			Title:       title,
			Content:     content,
			New:         true,
			CreatedTime: uint32(time.Now().Unix()),
		}
		if relatedIDFilter.Type.Set {
			msgs[i].RelatedMessageID = relatedIDFilter.Type.Value
		}
	}
	res := emptyList
	err := r.q.Transaction(func(tx *query.Query) error {
		txCtx := tx.WithContext(ctx)
		err := txCtx.PrivateMessage.Create(msgs...)
		if err != nil {
			r.log.Error("unexpected error", zap.Error(err))
			return errgo.Wrap(err, "dal")
		}
		if !relatedIDFilter.Type.Set {
			for i := range msgs {
				msgs[i].MainMessageID = msgs[i].ID
				msgs[i].RelatedMessageID = msgs[i].ID
			}
			err = txCtx.PrivateMessage.Save(msgs...)
			if err != nil {
				r.log.Error("unexpected error", zap.Error(err))
				return errgo.Wrap(err, "dal")
			}
		}
		res = slice.Map(msgs, convertDaoToModel)
		return nil
	})
	return res, err
}

func (r mysqlRepo) Delete(
	ctx context.Context,
	userID model.UserID,
	ids []model.PrivateMessageID,
) error {
	do := r.q.PrivateMessage.WithContext(ctx)
	pms, err := do.
		Where(
			r.q.PrivateMessage.ID.In(slice.ToValuer(ids)...),
			do.Where(
				r.q.PrivateMessage.SenderID.Eq(userID)).Or(r.q.PrivateMessage.ReceiverID.Eq(userID)),
		).Find()
	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return errgo.Wrap(err, "dal")
	}
	if len(pms) != len(ids) {
		return errUserIrrelevantMessageType
	}
	err = r.q.Transaction(func(tx *query.Query) error {
		err = handleReplyDeletes(ctx, tx, pms, userID)
		if err != nil {
			r.log.Error("unexpected error", zap.Error(err))
			return err
		}
		err = handleMainDeletes(ctx, tx, pms, userID)
		if err != nil {
			r.log.Error("unexpected error", zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		r.log.Error("unexpected error", zap.Error(err))
		return errgo.Wrap(err, "dal")
	}

	return nil
}

func handleReplyDeletes(ctx context.Context, tx *query.Query, pms []*dao.PrivateMessage, userID model.UserID) error {
	senderDeletes := slice.MapFilter(pms, func(v *dao.PrivateMessage) (driver.Valuer, bool) {
		ok := !v.DeletedBySender && v.MainMessageID == 0 && v.SenderID == userID
		if ok {
			return driver.Valuer(v.ID), ok
		}
		return nil, false
	})
	receiverDeletes := slice.MapFilter(pms, func(v *dao.PrivateMessage) (driver.Valuer, bool) {
		ok := !v.DeletedByReceiver && v.MainMessageID == 0 && v.ReceiverID == userID
		if ok {
			return driver.Valuer(v.ID), ok
		}
		return nil, false
	})
	txCtx := tx.WithContext(ctx)
	if len(senderDeletes) != 0 {
		_, err := txCtx.PrivateMessage.Where(
			tx.PrivateMessage.ID.In(senderDeletes...),
		).Update(tx.PrivateMessage.DeletedBySender, true)
		if err != nil {
			return errgo.Wrap(err, "dal")
		}
	}
	if len(receiverDeletes) != 0 {
		_, err := txCtx.PrivateMessage.Where(
			tx.PrivateMessage.ID.In(receiverDeletes...),
		).Update(tx.PrivateMessage.DeletedByReceiver, true)
		if err != nil {
			return errgo.Wrap(err, "dal")
		}
	}
	return nil
}

func handleMainDeletes(ctx context.Context, tx *query.Query, pms []*dao.PrivateMessage, userID model.UserID) error {
	senderDeletes := slice.MapFilter(pms, func(v *dao.PrivateMessage) (driver.Valuer, bool) {
		ok := v.MainMessageID != 0 && v.SenderID == userID
		if ok {
			return driver.Valuer(v.ID), ok
		}
		return nil, false
	})
	receiverDeletes := slice.MapFilter(pms, func(v *dao.PrivateMessage) (driver.Valuer, bool) {
		ok := v.MainMessageID != 0 && v.ReceiverID == userID
		if ok {
			return driver.Valuer(v.ID), ok
		}
		return nil, false
	})
	txCtx := tx.WithContext(ctx)
	if len(senderDeletes) != 0 {
		_, err := txCtx.PrivateMessage.Where(
			tx.PrivateMessage.RelatedMessageID.In(senderDeletes...),
			tx.PrivateMessage.DeletedBySender.Is(false),
		).Update(tx.PrivateMessage.DeletedBySender, true)
		if err != nil {
			return errgo.Wrap(err, "dal")
		}
	}
	if len(receiverDeletes) != 0 {
		_, err := txCtx.PrivateMessage.Where(
			tx.PrivateMessage.RelatedMessageID.In(receiverDeletes...),
			tx.PrivateMessage.DeletedByReceiver.Is(false),
		).Update(tx.PrivateMessage.DeletedByReceiver, true)
		if err != nil {
			return errgo.Wrap(err, "dal")
		}
	}
	return nil
}

func convertDaoToModel(d *dao.PrivateMessage) model.PrivateMessage {
	if d == nil {
		return model.PrivateMessage{}
	}
	return model.PrivateMessage{
		CreatedTime:       time.Unix(int64(d.CreatedTime), 0),
		Title:             d.Title,
		Content:           d.Content,
		Folder:            d.Folder,
		SenderID:          d.SenderID,
		ReceiverID:        d.ReceiverID,
		ID:                d.ID,
		MainMessageID:     d.MainMessageID,
		RelatedMessageID:  d.RelatedMessageID,
		New:               d.New,
		DeletedBySender:   d.DeletedBySender,
		DeletedByReceiver: d.DeletedByReceiver,
	}
}
