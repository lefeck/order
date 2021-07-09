package handler

import (
	"context"
	"github.com/asveg/order/domain/model"
	"github.com/asveg/order/domain/service"
	. "github.com/asveg/order/proto/order"
	"github.com/asveg/common"
)

type Order struct{
	OrderService service.IOrderService
}

/*
type OrderHandler interface {
	//通过订单id，获取订单信息
	GetOrderByID(context.Context, *OrderID, *OrderInfo) error
	//获取所有商品
	GetAllOrder(context.Context, *AllOrderRequest, *AllOrder) error
	//创建订单
	CreateOrder(context.Context, *OrderInfo, *OrderID) error
	//删除订单
	DeleteOrderByID(context.Context, *OrderID, *Response) error
	//更新订单支付状态
	UpdateOrderPayStatus(context.Context, *PayStatus, *Response) error
	//更新订单发货状态
	UpdateOrderShippStatus(context.Context, *ShippStatus, *Response) error
	//更新订单
	UpdateOrder(context.Context, *OrderInfo, *Response) error
}
 */

func (o *Order)GetOrderByID(ctx context.Context, request  *OrderID, response *OrderInfo) error  {
	order, err := o.OrderService.FindOrderByID(request.OrderId)
	if err !=nil {
		return err
	}
	if err := common.SwapTo(order,response); err !=nil {
		return err
	}
	return nil
}

func (o *Order)GetAllOrder(ctx context.Context, request *AllOrderRequest, response *AllOrder) error  {
	orderAll, err := o.OrderService.FindAll()
	if err !=nil {
		return err
	}
	for _,v :=range orderAll {
		order := &OrderInfo{}
		if err := common.SwapTo(v,request); err !=nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo,order)
	}
	return nil
}

func (o *Order) CreateOrder(ctx context.Context, request *OrderInfo, response *OrderID) error  {
	order := &model.Order{}

	if err := common.SwapTo(request,order); err !=nil {
		return err
	}

	orderID, err := o.OrderService.AddOrder(order)
	if err !=nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

func (o *Order)DeleteOrderByID(ctx context.Context, request  *OrderID, response  *Response) error  {
	if err :=o.OrderService.DeleteOrder(request.OrderId); err !=nil {
		return err
	}
	response.Msg = "delete order is success"
	return nil
}

func (o *Order) UpdateOrderPayStatus(ctx context.Context, request *PayStatus,  response *Response) error  {
	if err := o.OrderService.UpdatePayStatus(request.OrderId,request.PayStatus); err !=nil {
		return err
	}
	response.Msg="update order pay success"
	return nil
}

func (o *Order)UpdateOrderShippStatus(ctx context.Context,  request *ShippStatus, response *Response) error  {
	if err :=o.OrderService.UpdateShippStatus(request.OrderId,request.ShippStatus); err !=nil {
		return err
	}
	response.Msg = "update order shipp success"
	return nil
}

func (o *Order) UpdateOrder(ctx context.Context, request *OrderInfo, response *Response) error  {
	order := &model.Order{}
	if err :=common.SwapTo(request,order); err !=nil {
		return err
	}
	if err := o.OrderService.UpdateOrder(order); err !=nil {
		return err
	}
	response.Msg ="update order success"
	return nil
}