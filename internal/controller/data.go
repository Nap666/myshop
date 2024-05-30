package controller

import (
	"context"
	"myshop/api/backend"
	"myshop/internal/service"
)

type sData struct {
}

var Data = sData{}

func (c *sData) DataHead(ctx context.Context, req *backend.DateHeadReq) (res *backend.DateHeadRes, err error) {
	out, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DateHeadRes{
		out.TodayOrderCount,
		out.DAU,
		out.ConversionRate,
	}, nil
}
