package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"myshop/internal/dao"
	"myshop/internal/model"
	"myshop/internal/service"
)

type sData struct{}

func init() {
	service.RegisterData(New())
}
func New() *sData {
	return &sData{}
}

func (c *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             100,
		ConversionRate:  100,
	}, nil
}

func todayOrderCount(ctx context.Context) int {
	count, _ := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count()
	return count
}
