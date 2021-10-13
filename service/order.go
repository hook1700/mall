package service

import (
	"errors"
	"mall/repository"

	"mall/model"
	"mall/query"
)

type OrderSrv interface {
	List(req *query.ListQuery) (orders []*model.Order,err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(order model.Order) (*model.Order,error)
	Exist(order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(order model.Order) (*model.Order,error)
	Edit(order model.Order) (bool,error)
	Delete(o model.Order) (bool,error)
}

type OrderService struct {
	Repo repository.OrderRepoInterface
}

func (srv *OrderService) List(req *query.ListQuery) (orders []*model.Order, err error) {
	return srv.Repo.List(req)
}

func (srv *OrderService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *OrderService) Get(order model.Order) (*model.Order, error) {
	return srv.Repo.Get(order)
}

func (srv *OrderService) Exist(order model.Order) *model.Order {
	return srv.Repo.Exist(order)
}

func (srv *OrderService) ExistByOrderID(id string) *model.Order {
	return srv.Repo.ExistByOrderID(id)
}

func (srv *OrderService) Add(order model.Order) (*model.Order, error) {
	return srv.Repo.Add(order)
}

func (srv *OrderService) Edit(order model.Order) (bool, error) {
	o := srv.ExistByOrderID(order.OrderID)
	if o == nil || o.Mobile == ""{
		return false,errors.New("订单号不存在")
	}
	return srv.Repo.Edit(order)
}

func (srv *OrderService) Delete(o model.Order) (bool, error) {
	o.IsDeleted = !o.IsDeleted
	return srv.Repo.Delete(o)
}

