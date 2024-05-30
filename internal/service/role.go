// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"myshop/internal/model"
)

type (
	IRole interface {
		Add(ctx context.Context, in *model.RoleAddInput) (out *model.RoleAddOutput, err error)
		Delete(ctx context.Context, in *model.RoleDeleteInput) (out *model.RoleDeleteOutput, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
