// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardExtend is the golang structure for table evente_e_card_extend.
type EventeECardExtend struct {
	Id           uint        `json:"id"           orm:"id"            description:""`
	OrgId        int         `json:"orgId"        orm:"org_id"        description:"主办id"`
	CardId       int         `json:"cardId"       orm:"card_id"       description:"E通卡id"`
	EventeId     int         `json:"eventeId"     orm:"evente_id"     description:"活动id"`
	ScreeningsId int         `json:"screeningsId" orm:"screenings_id" description:"场次id"`
	PriceId      string      `json:"priceId"      orm:"price_id"      description:"票品ID集合 如 :9988,2235,6667"`
	CreateDate   *gtime.Time `json:"createDate"   orm:"create_date"   description:"创建时间"`
	UpdateDate   *gtime.Time `json:"updateDate"   orm:"update_date"   description:"修改时间"`
}
