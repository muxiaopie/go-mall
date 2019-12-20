package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
)

type SpuRepository interface {
	Create(spu model.Spu)(model.Spu,error)
	Find(id int)(model.Spu,error)
	Update(spu model.Spu) error
	Delete(spu model.Spu) error
	Pagination(page,limit int,maps map[string]interface{}) (*Pagination,error)
	ByName(name string) (model.Spu,error)
}

type Spu struct {
	DB *gorm.DB
}

func NewSpuRepository() (spuRepository SpuRepository) {
	return &Spu{
		DB:bootstrap.Bootstrap.DB,
	}
}

func (repo *Spu) Create(spu model.Spu) (model.Spu, error) {
	return spu,repo.DB.Create(&spu).Error
}

func (repo *Spu) Find(id int) (spu model.Spu, err error) {
	err = repo.DB.First(id,&spu).Error
	return
}

func (repo *Spu) Update(spu model.Spu) error {
	return repo.DB.Save(&spu).Error
}

func (repo *Spu) Delete(spu model.Spu) error {
	return repo.DB.Delete(&spu).Error
}

func (repo *Spu) Pagination(page, limit int, maps map[string]interface{}) (*Pagination, error) {
	var spus []*model.Spu
	p := NewPagination(page,limit)
	p.Pagination(repo.DB).Where(maps).Find(&spus)
	var count int
	repo.DB.Find(&model.Category{}).Where(maps).Count(&count)
	p.Total = count
	p.Items = spus
	return p,nil
}

func (repo *Spu) ByName(name string) (spu model.Spu, err error) {
	err = repo.DB.Where("name = ?",name).First(&spu).Error
	return
}







