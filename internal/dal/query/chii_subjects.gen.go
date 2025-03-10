// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newSubject(db *gorm.DB) subject {
	_subject := subject{}

	_subject.subjectDo.UseDB(db)
	_subject.subjectDo.UseModel(&dao.Subject{})

	tableName := _subject.subjectDo.TableName()
	_subject.ALL = field.NewAsterisk(tableName)
	_subject.ID = field.NewField(tableName, "subject_id")
	_subject.TypeID = field.NewUint8(tableName, "subject_type_id")
	_subject.Name = field.NewString(tableName, "subject_name")
	_subject.NameCN = field.NewString(tableName, "subject_name_cn")
	_subject.UID = field.NewString(tableName, "subject_uid")
	_subject.Creator = field.NewUint32(tableName, "subject_creator")
	_subject.Dateline = field.NewUint32(tableName, "subject_dateline")
	_subject.Image = field.NewString(tableName, "subject_image")
	_subject.Platform = field.NewUint16(tableName, "subject_platform")
	_subject.Infobox = field.NewString(tableName, "field_infobox")
	_subject.Summary = field.NewString(tableName, "field_summary")
	_subject.Field5 = field.NewString(tableName, "field_5")
	_subject.Volumes = field.NewUint32(tableName, "field_volumes")
	_subject.Eps = field.NewUint32(tableName, "field_eps")
	_subject.Wish = field.NewUint32(tableName, "subject_wish")
	_subject.Collect = field.NewUint32(tableName, "subject_collect")
	_subject.Doing = field.NewUint32(tableName, "subject_doing")
	_subject.OnHold = field.NewUint32(tableName, "subject_on_hold")
	_subject.Dropped = field.NewUint32(tableName, "subject_dropped")
	_subject.Series = field.NewBool(tableName, "subject_series")
	_subject.SeriesEntry = field.NewUint32(tableName, "subject_series_entry")
	_subject.IdxCn = field.NewString(tableName, "subject_idx_cn")
	_subject.Airtime = field.NewUint8(tableName, "subject_airtime")
	_subject.Nsfw = field.NewBool(tableName, "subject_nsfw")
	_subject.Ban = field.NewUint8(tableName, "subject_ban")
	_subject.Fields = subjectHasOneFields{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fields", "dao.SubjectField"),
	}

	_subject.fillFieldMap()

	return _subject
}

type subject struct {
	subjectDo subjectDo

	ALL         field.Asterisk
	ID          field.Field
	TypeID      field.Uint8
	Name        field.String
	NameCN      field.String
	UID         field.String
	Creator     field.Uint32
	Dateline    field.Uint32
	Image       field.String
	Platform    field.Uint16
	Infobox     field.String
	Summary     field.String
	Field5      field.String
	Volumes     field.Uint32
	Eps         field.Uint32
	Wish        field.Uint32
	Collect     field.Uint32
	Doing       field.Uint32
	OnHold      field.Uint32
	Dropped     field.Uint32
	Series      field.Bool
	SeriesEntry field.Uint32
	IdxCn       field.String
	Airtime     field.Uint8
	Nsfw        field.Bool
	Ban         field.Uint8
	Fields      subjectHasOneFields

	fieldMap map[string]field.Expr
}

func (s subject) Table(newTableName string) *subject {
	s.subjectDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subject) As(alias string) *subject {
	s.subjectDo.DO = *(s.subjectDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subject) updateTableName(table string) *subject {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewField(table, "subject_id")
	s.TypeID = field.NewUint8(table, "subject_type_id")
	s.Name = field.NewString(table, "subject_name")
	s.NameCN = field.NewString(table, "subject_name_cn")
	s.UID = field.NewString(table, "subject_uid")
	s.Creator = field.NewUint32(table, "subject_creator")
	s.Dateline = field.NewUint32(table, "subject_dateline")
	s.Image = field.NewString(table, "subject_image")
	s.Platform = field.NewUint16(table, "subject_platform")
	s.Infobox = field.NewString(table, "field_infobox")
	s.Summary = field.NewString(table, "field_summary")
	s.Field5 = field.NewString(table, "field_5")
	s.Volumes = field.NewUint32(table, "field_volumes")
	s.Eps = field.NewUint32(table, "field_eps")
	s.Wish = field.NewUint32(table, "subject_wish")
	s.Collect = field.NewUint32(table, "subject_collect")
	s.Doing = field.NewUint32(table, "subject_doing")
	s.OnHold = field.NewUint32(table, "subject_on_hold")
	s.Dropped = field.NewUint32(table, "subject_dropped")
	s.Series = field.NewBool(table, "subject_series")
	s.SeriesEntry = field.NewUint32(table, "subject_series_entry")
	s.IdxCn = field.NewString(table, "subject_idx_cn")
	s.Airtime = field.NewUint8(table, "subject_airtime")
	s.Nsfw = field.NewBool(table, "subject_nsfw")
	s.Ban = field.NewUint8(table, "subject_ban")

	s.fillFieldMap()

	return s
}

func (s *subject) WithContext(ctx context.Context) *subjectDo { return s.subjectDo.WithContext(ctx) }

func (s subject) TableName() string { return s.subjectDo.TableName() }

func (s subject) Alias() string { return s.subjectDo.Alias() }

func (s *subject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subject) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 26)
	s.fieldMap["subject_id"] = s.ID
	s.fieldMap["subject_type_id"] = s.TypeID
	s.fieldMap["subject_name"] = s.Name
	s.fieldMap["subject_name_cn"] = s.NameCN
	s.fieldMap["subject_uid"] = s.UID
	s.fieldMap["subject_creator"] = s.Creator
	s.fieldMap["subject_dateline"] = s.Dateline
	s.fieldMap["subject_image"] = s.Image
	s.fieldMap["subject_platform"] = s.Platform
	s.fieldMap["field_infobox"] = s.Infobox
	s.fieldMap["field_summary"] = s.Summary
	s.fieldMap["field_5"] = s.Field5
	s.fieldMap["field_volumes"] = s.Volumes
	s.fieldMap["field_eps"] = s.Eps
	s.fieldMap["subject_wish"] = s.Wish
	s.fieldMap["subject_collect"] = s.Collect
	s.fieldMap["subject_doing"] = s.Doing
	s.fieldMap["subject_on_hold"] = s.OnHold
	s.fieldMap["subject_dropped"] = s.Dropped
	s.fieldMap["subject_series"] = s.Series
	s.fieldMap["subject_series_entry"] = s.SeriesEntry
	s.fieldMap["subject_idx_cn"] = s.IdxCn
	s.fieldMap["subject_airtime"] = s.Airtime
	s.fieldMap["subject_nsfw"] = s.Nsfw
	s.fieldMap["subject_ban"] = s.Ban

}

func (s subject) clone(db *gorm.DB) subject {
	s.subjectDo.ReplaceDB(db)
	return s
}

type subjectHasOneFields struct {
	db *gorm.DB

	field.RelationField
}

func (a subjectHasOneFields) Where(conds ...field.Expr) *subjectHasOneFields {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a subjectHasOneFields) WithContext(ctx context.Context) *subjectHasOneFields {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a subjectHasOneFields) Model(m *dao.Subject) *subjectHasOneFieldsTx {
	return &subjectHasOneFieldsTx{a.db.Model(m).Association(a.Name())}
}

type subjectHasOneFieldsTx struct{ tx *gorm.Association }

func (a subjectHasOneFieldsTx) Find() (result *dao.SubjectField, err error) {
	return result, a.tx.Find(&result)
}

func (a subjectHasOneFieldsTx) Append(values ...*dao.SubjectField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a subjectHasOneFieldsTx) Replace(values ...*dao.SubjectField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a subjectHasOneFieldsTx) Delete(values ...*dao.SubjectField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a subjectHasOneFieldsTx) Clear() error {
	return a.tx.Clear()
}

func (a subjectHasOneFieldsTx) Count() int64 {
	return a.tx.Count()
}

type subjectDo struct{ gen.DO }

func (s subjectDo) Debug() *subjectDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectDo) WithContext(ctx context.Context) *subjectDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectDo) ReadDB() *subjectDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectDo) WriteDB() *subjectDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectDo) Clauses(conds ...clause.Expression) *subjectDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectDo) Returning(value interface{}, columns ...string) *subjectDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectDo) Not(conds ...gen.Condition) *subjectDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectDo) Or(conds ...gen.Condition) *subjectDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectDo) Select(conds ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectDo) Where(conds ...gen.Condition) *subjectDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectDo) Order(conds ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectDo) Distinct(cols ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectDo) Omit(cols ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectDo) Join(table schema.Tabler, on ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectDo) Group(cols ...field.Expr) *subjectDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectDo) Having(conds ...gen.Condition) *subjectDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectDo) Limit(limit int) *subjectDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectDo) Offset(offset int) *subjectDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectDo) Unscoped() *subjectDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectDo) Create(values ...*dao.Subject) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectDo) CreateInBatches(values []*dao.Subject, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectDo) Save(values ...*dao.Subject) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectDo) First() (*dao.Subject, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Subject), nil
	}
}

func (s subjectDo) Take() (*dao.Subject, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Subject), nil
	}
}

func (s subjectDo) Last() (*dao.Subject, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Subject), nil
	}
}

func (s subjectDo) Find() ([]*dao.Subject, error) {
	result, err := s.DO.Find()
	return result.([]*dao.Subject), err
}

func (s subjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Subject, err error) {
	buf := make([]*dao.Subject, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectDo) FindInBatches(result *[]*dao.Subject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectDo) Attrs(attrs ...field.AssignExpr) *subjectDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectDo) Assign(attrs ...field.AssignExpr) *subjectDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectDo) Joins(fields ...field.RelationField) *subjectDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectDo) Preload(fields ...field.RelationField) *subjectDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectDo) FirstOrInit() (*dao.Subject, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Subject), nil
	}
}

func (s subjectDo) FirstOrCreate() (*dao.Subject, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Subject), nil
	}
}

func (s subjectDo) FindByPage(offset int, limit int) (result []*dao.Subject, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectDo) Delete(models ...*dao.Subject) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectDo) withDO(do gen.Dao) *subjectDo {
	s.DO = *do.(*gen.DO)
	return s
}
