// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryInfoDao is the data access object for table category_info.
type CategoryInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns CategoryInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CategoryInfoColumns defines and stores column names for table category_info.
type CategoryInfoColumns struct {
	Id        string //
	ParentId  string // 父级id
	Name      string //
	PicUrl    string // icon
	DeletedAt string //
	CreatedAt string //
	UpdatedAt string //
	Level     string // 等级 默认1级分类
	Sort      string //
}

// categoryInfoColumns holds the columns for table category_info.
var categoryInfoColumns = CategoryInfoColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	PicUrl:    "pic_url",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Level:     "level",
	Sort:      "sort",
}

// NewCategoryInfoDao creates and returns a new DAO object for table data access.
func NewCategoryInfoDao() *CategoryInfoDao {
	return &CategoryInfoDao{
		group:   "default",
		table:   "category_info",
		columns: categoryInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CategoryInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CategoryInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CategoryInfoDao) Columns() CategoryInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CategoryInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CategoryInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CategoryInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
