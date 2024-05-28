package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleAddReq struct {
	g.Meta `path:"/backend/role/add" method:"post" desc:"添加角色" tags:"role" summary:"增加角色接口"`
	Name   string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleAddRes struct {
	RoleId uint `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"role" summary:"删除角色接口"`
	Id     int `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct {
	Id int `json:"id" dc:"删除时返回被删除的id"`
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"角色" summary:"角色列表接口"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
