package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"mall/model"
	"mall/query"
	"mall/repository"
)

type ProductSrv interface {
	List(req *query.ListQuery) (Products []*model.Product,err error)
	GetTotal(req *query.ListQuery)(total int,err error)
	Get(product model.Product) (*model.Product,error)
	Exist(product model.Product) *model.Product
	ExistByproductID(id string) *model.Product
	Add(product model.Product) (*model.Product, error)
	Edit(product model.Product) (bool,error)
	Delete(id string)(bool,error)
}

type ProductService struct {
	Repo repository.ProductRepoInterface
}

func (srv ProductService) List(req *query.ListQuery) (Products []*model.Product, err error) {
	return srv.Repo.List(req)
}

func (srv ProductService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv ProductService) Get(product model.Product) (*model.Product, error) {
	return srv.Repo.Get(product)
}

func (srv ProductService) Exist(product model.Product) *model.Product {
	return srv.Repo.Exist(product)
}

func (srv ProductService) ExistByproductID(id string) *model.Product {
	return srv.Repo.ExistByProductId(id)
}

func (srv ProductService) Add(product model.Product) (*model.Product, error) {
	if product.ProductID == ""{
		product.ProductID = uuid.NewV4().String()
	}
	product.IsDeleted = false
	return srv.Repo.Add(product)
}

func (srv ProductService) Edit(product model.Product) (bool, error) {
	p := srv.ExistByproductID(product.ProductID)
	if p == nil || p.ProductName == "" {
		return false, errors.New("参数传递错误")
	}
	return srv.Repo.Edit(product)
}

func (srv ProductService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}

	p := srv.ExistByproductID(id)
	if p == nil {
		return false, errors.New("参数错误")
	}
	p.IsDeleted = !p.IsDeleted
	return srv.Repo.Delete(*p)
}
