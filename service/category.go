package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/repository"
)

// var categoryOnce sync.Once


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
	/*categoryOnce.Do(func() {
		categoryService = &Category{
			Repo:repository.NewCategoryRepository(),
		}
	})
	return categoryService*/
}

func (sev *Category) Create(category model.Category) (model.Category, error) {
	return sev.Repo.Create(category)
}

func (sev *Category) ByName(name string) (model.Category, error) {
	return sev.Repo.ByName(name)
}


func (sev *Category) Find(id int) (model.Category, error) {
	return sev.Repo.Find(id)
}

func (sev *Category) Update(category model.Category) error {
	return sev.Repo.Update(category)
}

func (sev *Category) Delete(category model.Category) error {
	return sev.Repo.Delete(category)
}

func (srv *Category) Pagination(page, limit int,maps map[string]interface{}) (*repository.Pagination,error) {
	return srv.Repo.Pagination(page, limit, maps)
}


