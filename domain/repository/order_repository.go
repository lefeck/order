package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/asveg/order/domain/model"
)

type IOrderRepostory interface {
	InitTable()error
	FindOrderByID(orderID int64)(*model.Order,error)
	CreateOrder(order *model.Order)(OrderID int64, err error)
	DeleteOrder(orderID int64)error
	UpdateOrder(order *model.Order)error
	FindAll()(orderAll []model.Order,err error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64,int32) error
}
/*
操作数据，通过database 去调用实现的方法，实现对order 的增删改查。
 */

// mysql包装
type OrderRepostory struct {
	mysql *gorm.DB
}
//实例化对象，返回到接口，结构体只要实现了接口，这个结构体就可以返回到这个接口
func NewOrderRepostory(db *gorm.DB) IOrderRepostory  {
	return &OrderRepostory{mysql: db}
}

//inittiable tables
func (o *OrderRepostory) InitTable()error  {
	return o.mysql.CreateTable(&model.Order{},&model.OrderDetail{}).Error
}

//find order id message
func (o *OrderRepostory) FindOrderByID(orderID int64)( order *model.Order, err error) {
	order = &model.Order{}
	return order, o.mysql.Preload("OrderDetail").First(order,orderID).Error
}
//create order detail
func (o *OrderRepostory) CreateOrder(order *model.Order) (OrderID int64, err error) {
	return order.ID, o.mysql.Create(order).Error
}
// delete order by order id
func (o *OrderRepostory) DeleteOrder(orderID int64) error {
	tx := o.mysql.Begin()
	//error handle
	defer func() {
		if r:=recover();r!=nil {
			tx.Rollback()
		}
	}()

	if tx.Error !=nil {
		return tx.Error
	}

	//the delete of Order
	if err:= tx.Unscoped().Where("id = ?",orderID).Delete(&model.Order{}).Error;err!=nil{
		tx.Rollback()
		return err
	}

	//the delete of all OrderDetail
	if err:=tx.Unscoped().Where("order_id = ?",orderID).Delete(&model.OrderDetail{}).Error;err!=nil{
		tx.Rollback()
		return err

	}
	return tx.Commit().Error
}
// update order message
func (o *OrderRepostory) UpdateOrder(order *model.Order) error {
	return o.mysql.Model(order).Update(order).Error
}
// find all order message
func (o *OrderRepostory) FindAll() (orderAll []model.Order, err error) {
	return orderAll, o.mysql.Preload("OrderDetail").Find(&orderAll).Error
}
// update ship status message
func (o *OrderRepostory) UpdateShipStatus(orderID int64, shippStatus int32) error {
	db := o.mysql.Model(&model.Order{}).Where("id=?",orderID).UpdateColumn("ship_status",shippStatus)
	if db.Error !=nil {
		return db.Error
	}
	if db.RowsAffected ==0 {
		return errors.New("更新失败")
	}
	return nil
}
//update pay status message
func (o *OrderRepostory) UpdatePayStatus(orderID int64, payStatus int32) error {
	db := o.mysql.Model(&model.Order{}).Where("id=?",orderID).UpdateColumn("pay_status",payStatus)
	if db.Error !=nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}


