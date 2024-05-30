package controller

import (
	"context"
	"myshop/api/backend"
	"myshop/internal/model"
	"myshop/internal/service"
)

type sRole struct{}

var Role = sRole{}

func (s *sRole) AddController(ctx context.Context, req *backend.RoleAddReq) (res *backend.RoleAddRes, err error) {
	out, err := service.Role().Add(ctx,
		&model.RoleAddInput{
			Name: req.Name,
			Desc: req.Desc,
		},
	)
	if err != nil {
		return res, err
	}
	return &backend.RoleAddRes{
		RoleId: out.RoleId,
	}, nil
}
func (s *sRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	out, err := service.Role().Delete(ctx,
		&model.RoleDeleteInput{
			req.Id,
		})
	if err != nil {
		return res, err
	}
	res = &backend.RoleDeleteRes{
		Id: out.Id,
	}
	return res, nil
}
