package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"myshop/internal/dao"
	"myshop/internal/model"
	"myshop/internal/model/entity"
	"myshop/internal/service"
	"myshop/utility"
)

type sLogin struct {
}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if adminInfo.Password != utility.EncryptPassword(in.Password, adminInfo.UserSalt) {
		return gerror.New("密码不正确")
	}
	//设置用户Session.将用户信息保存到服务器
	err = service.Session().SetUser(ctx, &adminInfo)
	if err != nil {
		return err
	}

	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})

	return nil

}
