package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/repository"
)

type (
	// service 服务
	CategoryService interface {
		Create(category model.Category)(model.Category,error)
		Find(id int)(model.Category,error)
		Update(category model.Category) error
		Delete(category model.Category) error
		ByName(name string) (model.Category,error)
		Pagination(page, limit int,maps map[string]interface{}) (*repository.Pagination,error)
	}

	Category struct {
		Repo repository.CategoryRepository
	}
)

// 获取服务
func NewCategoryService() (categoryService CategoryService) {
	return &Category{
		Repo:repository.NewCategoryRepository(),
	}
}

func (srv *Category) Create(category model.Category) (model.Category, error) {
	return srv.Repo.Create(category)
}

func (srv *Category) ByName(name string) (model.Category, error) {
	return srv.Repo.ByName(name)
}

func (srv *Category) Find(id int) (model.Category, error) {
	return srv.Repo.Find(id)
}

func (srv *Category) Update(category model.Category) error {
	return srv.Repo.Update(category)
}

func (srv *Category) Delete(category model.Category) error {
	return srv.Repo.Delete(category)
}

func (srv *Category) Pagination(page, limit int,maps map[string]interface{}) (*repository.Pagination,error) {
	return srv.Repo.Pagination(page, limit, maps)
}


