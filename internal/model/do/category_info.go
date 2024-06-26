// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure of table category_info for DAO operations like Where/Data.
type CategoryInfo struct {
	g.Meta    `orm:"table:category_info, do:true"`
	Id        interface{} //
	ParentId  interface{} // 父级id
	Name      interface{} //
	PicUrl    interface{} // icon
	DeletedAt *gtime.Time //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Level     interface{} // 等级 默认1级分类
	Sort      interface{} //
}
