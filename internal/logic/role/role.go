package role

import (
	"context"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"myshop/internal/dao"
	"myshop/internal/model"
	"myshop/internal/service"
)

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (c *sRole) Add(ctx context.Context, in *model.RoleAddInput) (out *model.RoleAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	id, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.RoleAddOutput{RoleId: uint(id)}, nil
}

func (c *sRole) Delete(ctx context.Context, in *model.RoleDeleteInput) (out *model.RoleDeleteOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	dao.RoleInfo.Ctx(ctx).Where(
		g.Map{
			dao.RoleInfo.Columns().Id: in.Id,
		}).Unscoped().Delete()

	return &model.RoleDeleteOutput{
		Id: in.Id,
	}, err

}
