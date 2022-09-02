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

func newNotificationField(db *gorm.DB) notificationField {
	_notificationField := notificationField{}

	_notificationField.notificationFieldDo.UseDB(db)
	_notificationField.notificationFieldDo.UseModel(&dao.NotificationField{})

	tableName := _notificationField.notificationFieldDo.TableName()
	_notificationField.ALL = field.NewField(tableName, "*")
	_notificationField.ID = field.NewField(tableName, "ntf_id")
	_notificationField.RelatedType = field.NewUint8(tableName, "ntf_hash")
	_notificationField.RelatedID = field.NewUint32(tableName, "ntf_rid")
	_notificationField.Title = field.NewString(tableName, "ntf_title")

	_notificationField.fillFieldMap()

	return _notificationField
}

type notificationField struct {
	notificationFieldDo notificationFieldDo

	ALL         field.Field
	ID          field.Field
	RelatedType field.Uint8
	RelatedID   field.Uint32
	Title       field.String

	fieldMap map[string]field.Expr
}

func (n notificationField) Table(newTableName string) *notificationField {
	n.notificationFieldDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notificationField) As(alias string) *notificationField {
	n.notificationFieldDo.DO = *(n.notificationFieldDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notificationField) updateTableName(table string) *notificationField {
	n.ALL = field.NewField(table, "*")
	n.ID = field.NewField(table, "ntf_id")
	n.RelatedType = field.NewUint8(table, "ntf_hash")
	n.RelatedID = field.NewUint32(table, "ntf_rid")
	n.Title = field.NewString(table, "ntf_title")

	n.fillFieldMap()

	return n
}

func (n *notificationField) WithContext(ctx context.Context) *notificationFieldDo {
	return n.notificationFieldDo.WithContext(ctx)
}

func (n notificationField) TableName() string { return n.notificationFieldDo.TableName() }

func (n notificationField) Alias() string { return n.notificationFieldDo.Alias() }

func (n *notificationField) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notificationField) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 4)
	n.fieldMap["ntf_id"] = n.ID
	n.fieldMap["ntf_hash"] = n.RelatedType
	n.fieldMap["ntf_rid"] = n.RelatedID
	n.fieldMap["ntf_title"] = n.Title
}

func (n notificationField) clone(db *gorm.DB) notificationField {
	n.notificationFieldDo.ReplaceDB(db)
	return n
}

type notificationFieldDo struct{ gen.DO }

func (n notificationFieldDo) Debug() *notificationFieldDo {
	return n.withDO(n.DO.Debug())
}

func (n notificationFieldDo) WithContext(ctx context.Context) *notificationFieldDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n notificationFieldDo) ReadDB() *notificationFieldDo {
	return n.Clauses(dbresolver.Read)
}

func (n notificationFieldDo) WriteDB() *notificationFieldDo {
	return n.Clauses(dbresolver.Write)
}

func (n notificationFieldDo) Clauses(conds ...clause.Expression) *notificationFieldDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n notificationFieldDo) Returning(value interface{}, columns ...string) *notificationFieldDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n notificationFieldDo) Not(conds ...gen.Condition) *notificationFieldDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n notificationFieldDo) Or(conds ...gen.Condition) *notificationFieldDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n notificationFieldDo) Select(conds ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n notificationFieldDo) Where(conds ...gen.Condition) *notificationFieldDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n notificationFieldDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *notificationFieldDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n notificationFieldDo) Order(conds ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n notificationFieldDo) Distinct(cols ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n notificationFieldDo) Omit(cols ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n notificationFieldDo) Join(table schema.Tabler, on ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n notificationFieldDo) LeftJoin(table schema.Tabler, on ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n notificationFieldDo) RightJoin(table schema.Tabler, on ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n notificationFieldDo) Group(cols ...field.Expr) *notificationFieldDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n notificationFieldDo) Having(conds ...gen.Condition) *notificationFieldDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n notificationFieldDo) Limit(limit int) *notificationFieldDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n notificationFieldDo) Offset(offset int) *notificationFieldDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n notificationFieldDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *notificationFieldDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n notificationFieldDo) Unscoped() *notificationFieldDo {
	return n.withDO(n.DO.Unscoped())
}

func (n notificationFieldDo) Create(values ...*dao.NotificationField) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n notificationFieldDo) CreateInBatches(values []*dao.NotificationField, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n notificationFieldDo) Save(values ...*dao.NotificationField) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n notificationFieldDo) First() (*dao.NotificationField, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.NotificationField), nil
	}
}

func (n notificationFieldDo) Take() (*dao.NotificationField, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.NotificationField), nil
	}
}

func (n notificationFieldDo) Last() (*dao.NotificationField, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.NotificationField), nil
	}
}

func (n notificationFieldDo) Find() ([]*dao.NotificationField, error) {
	result, err := n.DO.Find()
	return result.([]*dao.NotificationField), err
}

func (n notificationFieldDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.NotificationField, err error) {
	buf := make([]*dao.NotificationField, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n notificationFieldDo) FindInBatches(result *[]*dao.NotificationField, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n notificationFieldDo) Attrs(attrs ...field.AssignExpr) *notificationFieldDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n notificationFieldDo) Assign(attrs ...field.AssignExpr) *notificationFieldDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n notificationFieldDo) Joins(fields ...field.RelationField) *notificationFieldDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n notificationFieldDo) Preload(fields ...field.RelationField) *notificationFieldDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n notificationFieldDo) FirstOrInit() (*dao.NotificationField, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.NotificationField), nil
	}
}

func (n notificationFieldDo) FirstOrCreate() (*dao.NotificationField, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.NotificationField), nil
	}
}

func (n notificationFieldDo) FindByPage(offset int, limit int) (result []*dao.NotificationField, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n notificationFieldDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n notificationFieldDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n notificationFieldDo) Delete(models ...*dao.NotificationField) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *notificationFieldDo) withDO(do gen.Dao) *notificationFieldDo {
	n.DO = *do.(*gen.DO)
	return n
}
