package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
)

type SkuRepository interface {
	Create(sku model.Sku)(model.Sku,error)
	Find(id int)(model.Sku,error)
	Update(sku model.Sku) error
	Delete(sku model.Sku) error
	Pagination(page,limit int,maps map[string]interface{}) (*Pagination,error)
	ByName(name string) (model.Sku,error)
}

type Sku struct {
	DB *gorm.DB
}

func NewSkuRepository() (skuRepository SkuRepository) {
	return &Sku{
		DB:bootstrap.Bootstrap.DB,
	}
}

func (repo *Sku) Create(sku model.Sku) (model.Sku, error) {
	return sku,repo.DB.Create(&sku).Error
}

func (repo *Sku) Find(id int) (sku model.Sku, err error) {
	err = repo.DB.First(id,&sku).Error
	return
}

func (repo *Sku) Update(sku model.Sku) error {
	return repo.DB.Save(&sku).Error

}

func (repo *Sku) Delete(sku model.Sku) error {
	return repo.DB.Delete(&sku).Error
}

func (repo *Sku) Pagination(page, limit int, maps map[string]interface{}) (*Pagination, error) {
	panic("implement me")
}

func (repo *Sku) ByName(name string) (model.Sku, error) {
	panic("implement me")
}







