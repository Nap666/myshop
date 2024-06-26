// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsInfoDao is the data access object for table goods_info.
type GoodsInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GoodsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// GoodsInfoColumns defines and stores column names for table goods_info.
type GoodsInfoColumns struct {
	Id               string //
	PicUrl           string // 图片
	Name             string // 商品名称
	Price            string // 价格 单位分
	Level1CategoryId string // 1级分类id
	Level2CategoryId string // 2级分类id
	Level3CategoryId string // 3级分类id
	Brand            string // 品牌
	CouponId         string // 优惠券id
	Stock            string // 库存
	Sale             string // 销量
	Tags             string // 标签
	DetailInfo       string // 商品详情
	CreatedAt        string //
	UpdatedAt        string //
	DeletedAt        string //
}

// goodsInfoColumns holds the columns for table goods_info.
var goodsInfoColumns = GoodsInfoColumns{
	Id:               "id",
	PicUrl:           "pic_url",
	Name:             "name",
	Price:            "price",
	Level1CategoryId: "level1_category_id",
	Level2CategoryId: "level2_category_id",
	Level3CategoryId: "level3_category_id",
	Brand:            "brand",
	CouponId:         "coupon_id",
	Stock:            "stock",
	Sale:             "sale",
	Tags:             "tags",
	DetailInfo:       "detail_info",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewGoodsInfoDao creates and returns a new DAO object for table data access.
func NewGoodsInfoDao() *GoodsInfoDao {
	return &GoodsInfoDao{
		group:   "default",
		table:   "goods_info",
		columns: goodsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsInfoDao) Columns() GoodsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
