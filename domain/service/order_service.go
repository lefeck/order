package service

import (
	"github.com/asveg/order/domain/model"
	"github.com/asveg/order/domain/repository"
)

type IOrderService interface {
	AddOrder(order *model.Order)(orderID int64, err error)
	DeleteOrder(orderID int64)error
	UpdateOrder(order *model.Order)error
	FindOrderByID(orderID int64)(order *model.Order,err error)
	FindAll()(orderAll []model.Order, err error)
	UpdateShippStatus(orderID int64, shippStatus int32)error
	UpdatePayStatus(orderID int64,Paystatus int32) error
}

type OrderService struct {
	OrderRepostory repository.IOrderRepostory
}

//一个结构体实现了一个接口， 那么函数中返回值类型为接口时，就应该返回这个结构体
func NewOrderService(orderRepostory repository.IOrderRepostory)IOrderService {
	return &OrderService{OrderRepostory: orderRepostory}
}

func (o *OrderService) AddOrder(order *model.Order) (orderID int64, err error) {
	return o.OrderRepostory.CreateOrder(order)
}

func (o *OrderService) DeleteOrder(orderID int64) error {
	return o.OrderRepostory.DeleteOrder(orderID)
}

func (o *OrderService) UpdateOrder(order *model.Order) error {
	return o.OrderRepostory.UpdateOrder(order)
}

func (o *OrderService) FindOrderByID(orderID int64) (order *model.Order, err error) {
	return o.OrderRepostory.FindOrderByID(orderID)
}

func (o *OrderService) FindAll() (orderAll []model.Order, err error) {
	return o.OrderRepostory.FindAll()
}

func (o *OrderService) UpdateShippStatus(orderID int64, shippStatus int32) error {
	return o.OrderRepostory.UpdateShipStatus(orderID,shippStatus)
}

func (o *OrderService) UpdatePayStatus(orderID int64, Paystatus int32) error {
	return o.OrderRepostory.UpdatePayStatus(orderID,Paystatus)
}

