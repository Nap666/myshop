// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardExtend is the golang structure of table evente_e_card_extend for DAO operations like Where/Data.
type EventeECardExtend struct {
	g.Meta       `orm:"table:evente_e_card_extend, do:true"`
	Id           interface{} //
	OrgId        interface{} // 主办id
	CardId       interface{} // E通卡id
	EventeId     interface{} // 活动id
	ScreeningsId interface{} // 场次id
	PriceId      interface{} // 票品ID集合 如 :9988,2235,6667
	CreateDate   *gtime.Time // 创建时间
	UpdateDate   *gtime.Time // 修改时间
}
