package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
)

// 获取repository
func NewBrandRepository() BrandRepository {
	return &Brand{
		DB:bootstrap.Bootstrap.DB,
	}
}


type BrandRepository interface {
	Create(brand model.Brand)(model.Brand,error)
	Find(id int)(model.Brand,error)
	Update(brand model.Brand) error
	Delete(brand model.Brand) error
	Pagination(page,limit int,maps map[string]interface{}) (*Pagination,error)
	ByName(name string) (model.Brand,error)
}

type Brand struct {
	DB *gorm.DB
}

// 增
func (repo *Brand) Create(brand model.Brand) (model.Brand,error) {
	return brand,repo.DB.Create(&brand).Error
}

func (repo *Brand) ByName(name string) (brand model.Brand,err error) {
	return brand,repo.DB.Where("name = ?",name).First(&brand).Error
}

// 查找
func (repo *Brand) Find(id int) (brand model.Brand,err error) {
	return brand,repo.DB.First(&brand,id).Error
}

// 修改
func (repo *Brand) Update(brand model.Brand) error {
	return repo.DB.Save(&brand).Error
}

// 删除
func (repo *Brand) Delete (brand model.Brand) error {
	return repo.DB.Delete(&brand).Error
}

// 分页
func (repo *Brand) Pagination (page,limit int,maps map[string]interface{}) (*Pagination,error)  {
	var brands []*model.Brand
	p := NewPagination(page,limit)
	p.Pagination(repo.DB).Where(maps).Find(&brands)
	var count int
	repo.DB.Find(&model.Brand{}).Where(maps).Count(&count)
	p.Total = count
	p.Items = brands
	return p,nil
}


