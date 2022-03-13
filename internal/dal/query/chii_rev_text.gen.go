// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/bangumi/server/internal/dal/dao"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newRevisionText(db *gorm.DB) revisionText {
	_revisionText := revisionText{}

	_revisionText.revisionTextDo.UseDB(db)
	_revisionText.revisionTextDo.UseModel(&dao.RevisionText{})

	tableName := _revisionText.revisionTextDo.TableName()
	_revisionText.ALL = field.NewField(tableName, "*")
	_revisionText.TextID = field.NewUint32(tableName, "rev_text_id")
	_revisionText.Text = field.NewField(tableName, "rev_text")

	_revisionText.fillFieldMap()

	return _revisionText
}

type revisionText struct {
	revisionTextDo revisionTextDo

	ALL    field.Field
	TextID field.Uint32
	Text   field.Field

	fieldMap map[string]field.Expr
}

func (r revisionText) Table(newTableName string) *revisionText {
	r.revisionTextDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r revisionText) As(alias string) *revisionText {
	r.revisionTextDo.DO = *(r.revisionTextDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *revisionText) updateTableName(table string) *revisionText {
	r.ALL = field.NewField(table, "*")
	r.TextID = field.NewUint32(table, "rev_text_id")
	r.Text = field.NewField(table, "rev_text")

	r.fillFieldMap()

	return r
}

func (r *revisionText) WithContext(ctx context.Context) *revisionTextDo {
	return r.revisionTextDo.WithContext(ctx)
}

func (r revisionText) TableName() string { return r.revisionTextDo.TableName() }

func (r *revisionText) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *revisionText) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["rev_text_id"] = r.TextID
	r.fieldMap["rev_text"] = r.Text
}

func (r revisionText) clone(db *gorm.DB) revisionText {
	r.revisionTextDo.ReplaceDB(db)
	return r
}

type revisionTextDo struct{ gen.DO }

func (r revisionTextDo) Debug() *revisionTextDo {
	return r.withDO(r.DO.Debug())
}

func (r revisionTextDo) WithContext(ctx context.Context) *revisionTextDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r revisionTextDo) Clauses(conds ...clause.Expression) *revisionTextDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r revisionTextDo) Returning(value interface{}, columns ...string) *revisionTextDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r revisionTextDo) Not(conds ...gen.Condition) *revisionTextDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r revisionTextDo) Or(conds ...gen.Condition) *revisionTextDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r revisionTextDo) Select(conds ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r revisionTextDo) Where(conds ...gen.Condition) *revisionTextDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r revisionTextDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *revisionTextDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r revisionTextDo) Order(conds ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r revisionTextDo) Distinct(cols ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r revisionTextDo) Omit(cols ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r revisionTextDo) Join(table schema.Tabler, on ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r revisionTextDo) LeftJoin(table schema.Tabler, on ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r revisionTextDo) RightJoin(table schema.Tabler, on ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r revisionTextDo) Group(cols ...field.Expr) *revisionTextDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r revisionTextDo) Having(conds ...gen.Condition) *revisionTextDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r revisionTextDo) Limit(limit int) *revisionTextDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r revisionTextDo) Offset(offset int) *revisionTextDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r revisionTextDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *revisionTextDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r revisionTextDo) Unscoped() *revisionTextDo {
	return r.withDO(r.DO.Unscoped())
}

func (r revisionTextDo) Create(values ...*dao.RevisionText) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r revisionTextDo) CreateInBatches(values []*dao.RevisionText, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r revisionTextDo) Save(values ...*dao.RevisionText) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r revisionTextDo) First() (*dao.RevisionText, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionText), nil
	}
}

func (r revisionTextDo) Take() (*dao.RevisionText, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionText), nil
	}
}

func (r revisionTextDo) Last() (*dao.RevisionText, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionText), nil
	}
}

func (r revisionTextDo) Find() ([]*dao.RevisionText, error) {
	result, err := r.DO.Find()
	return result.([]*dao.RevisionText), err
}

func (r revisionTextDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.RevisionText, err error) {
	buf := make([]*dao.RevisionText, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r revisionTextDo) FindInBatches(result *[]*dao.RevisionText, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r revisionTextDo) Attrs(attrs ...field.AssignExpr) *revisionTextDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r revisionTextDo) Assign(attrs ...field.AssignExpr) *revisionTextDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r revisionTextDo) Joins(field field.RelationField) *revisionTextDo {
	return r.withDO(r.DO.Joins(field))
}

func (r revisionTextDo) Preload(field field.RelationField) *revisionTextDo {
	return r.withDO(r.DO.Preload(field))
}

func (r revisionTextDo) FirstOrInit() (*dao.RevisionText, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionText), nil
	}
}

func (r revisionTextDo) FirstOrCreate() (*dao.RevisionText, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionText), nil
	}
}

func (r revisionTextDo) FindByPage(offset int, limit int) (result []*dao.RevisionText, count int64, err error) {
	if limit <= 0 {
		count, err = r.Count()
		return
	}

	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r revisionTextDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *revisionTextDo) withDO(do gen.Dao) *revisionTextDo {
	r.DO = *do.(*gen.DO)
	return r
}
