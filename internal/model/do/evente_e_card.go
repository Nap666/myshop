// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECard is the golang structure of table evente_e_card for DAO operations like Where/Data.
type EventeECard struct {
	g.Meta        `orm:"table:evente_e_card, do:true"`
	Id            interface{} //
	OrgId         interface{} // 主办id
	CardName      interface{} // E通卡名称
	SaleStartDate *gtime.Time // 售卖开始时间
	SaleStopDate  *gtime.Time // 售卖结束时间
	Usage         interface{} // 使用方式 1、预约 2、免预约
}
