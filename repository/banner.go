package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mall/model"
	"mall/query"
	"mall/utils"
)

type BannerRepository struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	List(req *query.ListQuery) (banners []*model.Banner,err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(banner model.Banner) (*model.Banner,error)
	Exist(banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(banner model.Banner) (*model.Banner,error)
	Edit(banner model.Banner) (bool,error)
	Delete(id string) (bool,error)
}

func (repo *BannerRepository) List(req *query.ListQuery) (banners []*model.Banner, err error) {
	fmt.Println(req)
	limit,offset := utils.Page(req.PageSize,req.Page)

	err = repo.DB.Limit(limit).Offset(offset).Find(&banners).Error;
	if err != nil {
		return nil, err
	}
	return banners,nil
}

func (repo *BannerRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var banners []model.Banner
	err = repo.DB.Find(&banners).Count(&total).Error
	if err != nil {
		return total,err
	}
	return total,err
}

func (repo *BannerRepository) Get(banner model.Banner) (*model.Banner, error) {
	if err := repo.DB.Where(&banner).Find(&banner).Error;err != nil{
		return nil, err
	}
	return &banner,nil
}

func (repo *BannerRepository) Exist(banner model.Banner) *model.Banner {
	if banner.Url != ""&& banner.RedirectUrl != ""{
		var b model.Banner
		repo.DB.Model(&banner).Where("url = ? and redirect_url",banner.Url,banner.RedirectUrl).First(&b)
		return &b
	}
	return nil
}

func (repo *BannerRepository) ExistByBannerID(id string) *model.Banner {
	var b model.Banner
	repo.DB.Where("id = ?").First(&b)
	return &b
}

func (repo *BannerRepository) Add(banner model.Banner) (*model.Banner, error) {
	exist := repo.Exist(banner)
	if exist != nil && exist.Url ==banner.Url && exist.RedirectUrl == banner.RedirectUrl{
		return nil,fmt.Errorf("轮播图已存在")
	}
	err := repo.DB.Create(banner).Error
	if err != nil{
		return nil,fmt.Errorf("轮播图添加失败")
	}
	return &banner,nil
}

func (repo *BannerRepository) Edit(banner model.Banner) (bool, error) {
	if banner.BannerID == ""{
		return false,fmt.Errorf("请传入更新ID")
	}
	b := &model.Banner{}
	err := repo.DB.Model(b).Where("banner_id = ?",banner.BannerID).Updates(map[string]interface{}{
		"Url": banner.Url,
		"RedirectUrl": banner.RedirectUrl,
		"OrderBy":     banner.Order,
	}).Error
	if err != nil{
		return false, err
	}
	return true,nil
}

func (repo *BannerRepository) Delete(id string) (bool, error) {
	err := repo.DB.Where("banner_id = ?", id).Delete(&model.Banner{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

