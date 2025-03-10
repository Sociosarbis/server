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

func newGroupTopicComment(db *gorm.DB) groupTopicComment {
	_groupTopicComment := groupTopicComment{}

	_groupTopicComment.groupTopicCommentDo.UseDB(db)
	_groupTopicComment.groupTopicCommentDo.UseModel(&dao.GroupTopicComment{})

	tableName := _groupTopicComment.groupTopicCommentDo.TableName()
	_groupTopicComment.ALL = field.NewAsterisk(tableName)
	_groupTopicComment.ID = field.NewUint32(tableName, "grp_pst_id")
	_groupTopicComment.TopicID = field.NewUint32(tableName, "grp_pst_mid")
	_groupTopicComment.UID = field.NewUint32(tableName, "grp_pst_uid")
	_groupTopicComment.Related = field.NewUint32(tableName, "grp_pst_related")
	_groupTopicComment.Content = field.NewString(tableName, "grp_pst_content")
	_groupTopicComment.State = field.NewUint8(tableName, "grp_pst_state")
	_groupTopicComment.CreatedTime = field.NewUint32(tableName, "grp_pst_dateline")

	_groupTopicComment.fillFieldMap()

	return _groupTopicComment
}

type groupTopicComment struct {
	groupTopicCommentDo groupTopicCommentDo

	ALL         field.Asterisk
	ID          field.Uint32
	TopicID     field.Uint32
	UID         field.Uint32
	Related     field.Uint32
	Content     field.String
	State       field.Uint8
	CreatedTime field.Uint32

	fieldMap map[string]field.Expr
}

func (g groupTopicComment) Table(newTableName string) *groupTopicComment {
	g.groupTopicCommentDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupTopicComment) As(alias string) *groupTopicComment {
	g.groupTopicCommentDo.DO = *(g.groupTopicCommentDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupTopicComment) updateTableName(table string) *groupTopicComment {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewUint32(table, "grp_pst_id")
	g.TopicID = field.NewUint32(table, "grp_pst_mid")
	g.UID = field.NewUint32(table, "grp_pst_uid")
	g.Related = field.NewUint32(table, "grp_pst_related")
	g.Content = field.NewString(table, "grp_pst_content")
	g.State = field.NewUint8(table, "grp_pst_state")
	g.CreatedTime = field.NewUint32(table, "grp_pst_dateline")

	g.fillFieldMap()

	return g
}

func (g *groupTopicComment) WithContext(ctx context.Context) *groupTopicCommentDo {
	return g.groupTopicCommentDo.WithContext(ctx)
}

func (g groupTopicComment) TableName() string { return g.groupTopicCommentDo.TableName() }

func (g groupTopicComment) Alias() string { return g.groupTopicCommentDo.Alias() }

func (g *groupTopicComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupTopicComment) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["grp_pst_id"] = g.ID
	g.fieldMap["grp_pst_mid"] = g.TopicID
	g.fieldMap["grp_pst_uid"] = g.UID
	g.fieldMap["grp_pst_related"] = g.Related
	g.fieldMap["grp_pst_content"] = g.Content
	g.fieldMap["grp_pst_state"] = g.State
	g.fieldMap["grp_pst_dateline"] = g.CreatedTime
}

func (g groupTopicComment) clone(db *gorm.DB) groupTopicComment {
	g.groupTopicCommentDo.ReplaceDB(db)
	return g
}

type groupTopicCommentDo struct{ gen.DO }

func (g groupTopicCommentDo) Debug() *groupTopicCommentDo {
	return g.withDO(g.DO.Debug())
}

func (g groupTopicCommentDo) WithContext(ctx context.Context) *groupTopicCommentDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupTopicCommentDo) ReadDB() *groupTopicCommentDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupTopicCommentDo) WriteDB() *groupTopicCommentDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupTopicCommentDo) Clauses(conds ...clause.Expression) *groupTopicCommentDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupTopicCommentDo) Returning(value interface{}, columns ...string) *groupTopicCommentDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupTopicCommentDo) Not(conds ...gen.Condition) *groupTopicCommentDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupTopicCommentDo) Or(conds ...gen.Condition) *groupTopicCommentDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupTopicCommentDo) Select(conds ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupTopicCommentDo) Where(conds ...gen.Condition) *groupTopicCommentDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupTopicCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *groupTopicCommentDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g groupTopicCommentDo) Order(conds ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupTopicCommentDo) Distinct(cols ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupTopicCommentDo) Omit(cols ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupTopicCommentDo) Join(table schema.Tabler, on ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupTopicCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupTopicCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupTopicCommentDo) Group(cols ...field.Expr) *groupTopicCommentDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupTopicCommentDo) Having(conds ...gen.Condition) *groupTopicCommentDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupTopicCommentDo) Limit(limit int) *groupTopicCommentDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupTopicCommentDo) Offset(offset int) *groupTopicCommentDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupTopicCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *groupTopicCommentDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupTopicCommentDo) Unscoped() *groupTopicCommentDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupTopicCommentDo) Create(values ...*dao.GroupTopicComment) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupTopicCommentDo) CreateInBatches(values []*dao.GroupTopicComment, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupTopicCommentDo) Save(values ...*dao.GroupTopicComment) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupTopicCommentDo) First() (*dao.GroupTopicComment, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.GroupTopicComment), nil
	}
}

func (g groupTopicCommentDo) Take() (*dao.GroupTopicComment, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.GroupTopicComment), nil
	}
}

func (g groupTopicCommentDo) Last() (*dao.GroupTopicComment, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.GroupTopicComment), nil
	}
}

func (g groupTopicCommentDo) Find() ([]*dao.GroupTopicComment, error) {
	result, err := g.DO.Find()
	return result.([]*dao.GroupTopicComment), err
}

func (g groupTopicCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.GroupTopicComment, err error) {
	buf := make([]*dao.GroupTopicComment, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupTopicCommentDo) FindInBatches(result *[]*dao.GroupTopicComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupTopicCommentDo) Attrs(attrs ...field.AssignExpr) *groupTopicCommentDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupTopicCommentDo) Assign(attrs ...field.AssignExpr) *groupTopicCommentDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupTopicCommentDo) Joins(fields ...field.RelationField) *groupTopicCommentDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupTopicCommentDo) Preload(fields ...field.RelationField) *groupTopicCommentDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupTopicCommentDo) FirstOrInit() (*dao.GroupTopicComment, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.GroupTopicComment), nil
	}
}

func (g groupTopicCommentDo) FirstOrCreate() (*dao.GroupTopicComment, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.GroupTopicComment), nil
	}
}

func (g groupTopicCommentDo) FindByPage(offset int, limit int) (result []*dao.GroupTopicComment, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g groupTopicCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupTopicCommentDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupTopicCommentDo) Delete(models ...*dao.GroupTopicComment) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupTopicCommentDo) withDO(do gen.Dao) *groupTopicCommentDo {
	g.DO = *do.(*gen.DO)
	return g
}
