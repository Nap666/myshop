// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionInfo is the golang structure of table collection_info for DAO operations like Where/Data.
type CollectionInfo struct {
	g.Meta    `orm:"table:collection_info, do:true"`
	Id        interface{} //
	UserId    interface{} //
	ObjectId  interface{} //
	Type      interface{} // 收藏类型：1商品 2文章
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
