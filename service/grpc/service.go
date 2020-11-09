package grpc

import (
	"context"
	"prott/protoorder"
)

// OrderService 订单服务
type OrderService struct{}

// CreateOrder 创建
func (o *OrderService) CreateOrder(ctx context.Context, req *protoorder.QOrderInfo, rsp *protoorder.ROrderResult) error {
	rsp.OrderId = req.GoodId
	return nil
}
