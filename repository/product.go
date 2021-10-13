package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mall/model"
	"mall/query"
	"mall/utils"
)

type ProductRepository struct {
	DB *gorm.DB
}



type ProductRepoInterface interface {
	List(req *query.ListQuery) (products []*model.Product,err error)
	GetTotal(req *query.ListQuery) (total int,err error)
	Get(product model.Product) (*model.Product,error)
	Exist(product model.Product) *model.Product
	ExistByProductId(id string) *model.Product
	Add(product model.Product) (*model.Product,error)
	Edit(product model.Product) (bool,error)
	Delete(product model.Product) (bool,error)
}

func (repo *ProductRepository) List (req *query.ListQuery) (products []*model.Product,err error) {
	fmt.Println(req)
	db := repo.DB
	//算出页码
	limit,offset := utils.Page(req.PageSize,req.Page)

	if err := db.Limit(limit).Offset(offset).Find(&products).Error;err != nil{
		return nil,err
	}
	return products,err
}

func (repo *ProductRepository) GetTotal  (req *query.ListQuery) (total int ,err error){
	var products []*model.Product
	db := repo.DB
	if err := db.Find(&products).Count(&total).Error;err != nil{
		return total,err
	}
	return total,nil
}

func (repo *ProductRepository)  Get (product model.Product) (*model.Product,error){
	if err := repo.DB.Where(&product).Find(&product).Error;err != nil{
		return nil,err
	}
	return &product,nil
}

func (repo *ProductRepository) Exist(product model.Product) *model.Product {
	if product.ProductName != ""{
		var temp model.Product
		repo.DB.Where("product_name = ?",temp.ProductName).First(&temp)
		return &temp
	}
	return nil
}

func (repo *ProductRepository) ExistByProductId(id string) *model.Product {
	var p model.Product
	repo.DB.Where("id = ?", id).First(&p)
	return &p

}

func (repo *ProductRepository) Add(product model.Product)(*model.Product,error) {
	exist := repo.Exist(product)
	if exist != nil && exist.ProductName != ""{
		return &product,fmt.Errorf("商品已存在")
	}
	err := repo.DB.Create(product).Error
	if err != nil{
		return nil, fmt.Errorf("商品添加成功")
	}
	return &product,nil
}
func (repo *ProductRepository) Edit(product model.Product) (bool, error) {
	if product.ProductID == ""{
		return false, fmt.Errorf("请传入更新 ID")
	}
	p := &model.Product{}
	err := repo.DB.Model(p).Where("product_id = ?",product.ProductID).Updates(map[string]interface{}{
		"product_name":product.ProductName,"product_intro":product.ProductIntro,"category_id":product.CategoryID,
		"product_cover_imp":product.ProductCoverImg,"product_banner":product.ProductBanner,"original_price":product.OriginalPrice,
		"selling_price":product.SellingPrice,"stock_num":product.StockNum,"tag":product.Tag,"sell_status":product.SellStatus,
		"product_detail_content":product.ProductDetailContent,
	}).Error
	if err != nil{
		return false, err
	}
	return true,nil
}

func (repo *ProductRepository) Delete(product model.Product) (bool, error) {
	err := repo.DB.Model(&product).Where("product_id = ?",product.ProductID).Updates("is_deleted",product.IsDeleted).Error
	if err != nil{
		return false, err
	}
	return true,nil
}
