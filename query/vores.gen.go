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

	"elsenova/models"
)

func newVore(db *gorm.DB, opts ...gen.DOOption) vore {
	_vore := vore{}

	_vore.voreDo.UseDB(db, opts...)
	_vore.voreDo.UseModel(&models.Vore{})

	tableName := _vore.voreDo.TableName()
	_vore.ALL = field.NewAsterisk(tableName)
	_vore.ID = field.NewUint(tableName, "id")
	_vore.CreatedAt = field.NewTime(tableName, "created_at")
	_vore.UpdatedAt = field.NewTime(tableName, "updated_at")
	_vore.DeletedAt = field.NewField(tableName, "deleted_at")
	_vore.UserID = field.NewString(tableName, "user_id")

	_vore.fillFieldMap()

	return _vore
}

type vore struct {
	voreDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.String

	fieldMap map[string]field.Expr
}

func (v vore) Table(newTableName string) *vore {
	v.voreDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vore) As(alias string) *vore {
	v.voreDo.DO = *(v.voreDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vore) updateTableName(table string) *vore {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewUint(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.UserID = field.NewString(table, "user_id")

	v.fillFieldMap()

	return v
}

func (v *vore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vore) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 5)
	v.fieldMap["id"] = v.ID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["user_id"] = v.UserID
}

func (v vore) clone(db *gorm.DB) vore {
	v.voreDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vore) replaceDB(db *gorm.DB) vore {
	v.voreDo.ReplaceDB(db)
	return v
}

type voreDo struct{ gen.DO }

type IVoreDo interface {
	gen.SubQuery
	Debug() IVoreDo
	WithContext(ctx context.Context) IVoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVoreDo
	WriteDB() IVoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVoreDo
	Not(conds ...gen.Condition) IVoreDo
	Or(conds ...gen.Condition) IVoreDo
	Select(conds ...field.Expr) IVoreDo
	Where(conds ...gen.Condition) IVoreDo
	Order(conds ...field.Expr) IVoreDo
	Distinct(cols ...field.Expr) IVoreDo
	Omit(cols ...field.Expr) IVoreDo
	Join(table schema.Tabler, on ...field.Expr) IVoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVoreDo
	Group(cols ...field.Expr) IVoreDo
	Having(conds ...gen.Condition) IVoreDo
	Limit(limit int) IVoreDo
	Offset(offset int) IVoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVoreDo
	Unscoped() IVoreDo
	Create(values ...*models.Vore) error
	CreateInBatches(values []*models.Vore, batchSize int) error
	Save(values ...*models.Vore) error
	First() (*models.Vore, error)
	Take() (*models.Vore, error)
	Last() (*models.Vore, error)
	Find() ([]*models.Vore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vore, err error)
	FindInBatches(result *[]*models.Vore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Vore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVoreDo
	Assign(attrs ...field.AssignExpr) IVoreDo
	Joins(fields ...field.RelationField) IVoreDo
	Preload(fields ...field.RelationField) IVoreDo
	FirstOrInit() (*models.Vore, error)
	FirstOrCreate() (*models.Vore, error)
	FindByPage(offset int, limit int) (result []*models.Vore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v voreDo) Debug() IVoreDo {
	return v.withDO(v.DO.Debug())
}

func (v voreDo) WithContext(ctx context.Context) IVoreDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v voreDo) ReadDB() IVoreDo {
	return v.Clauses(dbresolver.Read)
}

func (v voreDo) WriteDB() IVoreDo {
	return v.Clauses(dbresolver.Write)
}

func (v voreDo) Session(config *gorm.Session) IVoreDo {
	return v.withDO(v.DO.Session(config))
}

func (v voreDo) Clauses(conds ...clause.Expression) IVoreDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v voreDo) Returning(value interface{}, columns ...string) IVoreDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v voreDo) Not(conds ...gen.Condition) IVoreDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v voreDo) Or(conds ...gen.Condition) IVoreDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v voreDo) Select(conds ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v voreDo) Where(conds ...gen.Condition) IVoreDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v voreDo) Order(conds ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v voreDo) Distinct(cols ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v voreDo) Omit(cols ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v voreDo) Join(table schema.Tabler, on ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v voreDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVoreDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v voreDo) RightJoin(table schema.Tabler, on ...field.Expr) IVoreDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v voreDo) Group(cols ...field.Expr) IVoreDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v voreDo) Having(conds ...gen.Condition) IVoreDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v voreDo) Limit(limit int) IVoreDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v voreDo) Offset(offset int) IVoreDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v voreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVoreDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v voreDo) Unscoped() IVoreDo {
	return v.withDO(v.DO.Unscoped())
}

func (v voreDo) Create(values ...*models.Vore) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v voreDo) CreateInBatches(values []*models.Vore, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v voreDo) Save(values ...*models.Vore) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v voreDo) First() (*models.Vore, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vore), nil
	}
}

func (v voreDo) Take() (*models.Vore, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vore), nil
	}
}

func (v voreDo) Last() (*models.Vore, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vore), nil
	}
}

func (v voreDo) Find() ([]*models.Vore, error) {
	result, err := v.DO.Find()
	return result.([]*models.Vore), err
}

func (v voreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vore, err error) {
	buf := make([]*models.Vore, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v voreDo) FindInBatches(result *[]*models.Vore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v voreDo) Attrs(attrs ...field.AssignExpr) IVoreDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v voreDo) Assign(attrs ...field.AssignExpr) IVoreDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v voreDo) Joins(fields ...field.RelationField) IVoreDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v voreDo) Preload(fields ...field.RelationField) IVoreDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v voreDo) FirstOrInit() (*models.Vore, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vore), nil
	}
}

func (v voreDo) FirstOrCreate() (*models.Vore, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vore), nil
	}
}

func (v voreDo) FindByPage(offset int, limit int) (result []*models.Vore, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v voreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v voreDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v voreDo) Delete(models ...*models.Vore) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *voreDo) withDO(do gen.Dao) *voreDo {
	v.DO = *do.(*gen.DO)
	return v
}