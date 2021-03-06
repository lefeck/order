// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	//通过订单id，获取订单信息
	GetOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*OrderInfo, error)
	//获取所有商品
	GetAllOrder(ctx context.Context, in *AllOrderRequest, opts ...client.CallOption) (*AllOrder, error)
	//创建订单
	CreateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*OrderID, error)
	//删除订单
	DeleteOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*Response, error)
	//更新订单支付状态
	UpdateOrderPayStatus(ctx context.Context, in *PayStatus, opts ...client.CallOption) (*Response, error)
	//更新订单发货状态
	UpdateOrderShippStatus(ctx context.Context, in *ShippStatus, opts ...client.CallOption) (*Response, error)
	//更新订单
	UpdateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*Response, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*OrderInfo, error) {
	req := c.c.NewRequest(c.name, "Order.GetOrderByID", in)
	out := new(OrderInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetAllOrder(ctx context.Context, in *AllOrderRequest, opts ...client.CallOption) (*AllOrder, error) {
	req := c.c.NewRequest(c.name, "Order.GetAllOrder", in)
	out := new(AllOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CreateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*OrderID, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(OrderID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) DeleteOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.DeleteOrderByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderPayStatus(ctx context.Context, in *PayStatus, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrderPayStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderShippStatus(ctx context.Context, in *ShippStatus, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrderShippStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

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

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		GetOrderByID(ctx context.Context, in *OrderID, out *OrderInfo) error
		GetAllOrder(ctx context.Context, in *AllOrderRequest, out *AllOrder) error
		CreateOrder(ctx context.Context, in *OrderInfo, out *OrderID) error
		DeleteOrderByID(ctx context.Context, in *OrderID, out *Response) error
		UpdateOrderPayStatus(ctx context.Context, in *PayStatus, out *Response) error
		UpdateOrderShippStatus(ctx context.Context, in *ShippStatus, out *Response) error
		UpdateOrder(ctx context.Context, in *OrderInfo, out *Response) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) GetOrderByID(ctx context.Context, in *OrderID, out *OrderInfo) error {
	return h.OrderHandler.GetOrderByID(ctx, in, out)
}

func (h *orderHandler) GetAllOrder(ctx context.Context, in *AllOrderRequest, out *AllOrder) error {
	return h.OrderHandler.GetAllOrder(ctx, in, out)
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *OrderInfo, out *OrderID) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}

func (h *orderHandler) DeleteOrderByID(ctx context.Context, in *OrderID, out *Response) error {
	return h.OrderHandler.DeleteOrderByID(ctx, in, out)
}

func (h *orderHandler) UpdateOrderPayStatus(ctx context.Context, in *PayStatus, out *Response) error {
	return h.OrderHandler.UpdateOrderPayStatus(ctx, in, out)
}

func (h *orderHandler) UpdateOrderShippStatus(ctx context.Context, in *ShippStatus, out *Response) error {
	return h.OrderHandler.UpdateOrderShippStatus(ctx, in, out)
}

func (h *orderHandler) UpdateOrder(ctx context.Context, in *OrderInfo, out *Response) error {
	return h.OrderHandler.UpdateOrder(ctx, in, out)
}
