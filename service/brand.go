package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/repository"
)

type (
	// service 服务
	BrandService interface {
		Create(brand model.Brand)(model.Brand,error)
		Find(id int)(model.Brand,error)
		Update(brand model.Brand) error
		Delete(brand model.Brand) error
		ByName(name string) (model.Brand,error)
		Pagination(page, limit int,maps map[string]interface{}) (*repository.Pagination,error)
	}
	Brand struct {
		Repo repository.BrandRepository
	}
)

func (srv Brand) Create(brand model.Brand) (model.Brand, error) {
	return  srv.Repo.Create(brand)
}

func (srv Brand) Find(id int) (model.Brand, error) {
	return srv.Repo.Find(id)
}

func (srv Brand) Update(brand model.Brand) error {
	return srv.Repo.Update(brand)
}

func (srv Brand) Delete(brand model.Brand) error {
	return srv.Repo.Delete(brand)
}

func (srv Brand) ByName(name string) (model.Brand, error) {
	return srv.Repo.ByName(name)
}

func (srv Brand) Pagination(page, limit int, maps map[string]interface{}) (*repository.Pagination, error) {
	return srv.Repo.Pagination(page,limit,maps)
}

// 获取服务
func NewBrandService() (brandService BrandService) {
	return &Brand{
		Repo:repository.NewBrandRepository(),
	}
}


