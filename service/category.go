package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/repository"
	"sync"
)

var categoryOnce sync.Once


type (
	// service 服务
	CategoryService interface {
		Create(category model.Category)(model.Category,error)
		Find(id int)(model.Category,error)
		Update(category model.Category) error
		Delete(category model.Category) error
	}

	Category struct {
		Repo repository.CategoryRepository
	}
)

// 获取服务
func NewCategoryService() (categoryService CategoryService) {
	categoryOnce.Do(func() {
		categoryService = &Category{
			Repo:repository.NewCategoryRepository(),
		}
	})
	return categoryService
}

func (sev *Category) Create(category model.Category) (model.Category, error) {
	return sev.Repo.Create(category)
}

func (sev *Category) Find(id int) (model.Category, error) {
	return sev.Repo.Find(id)
}

func (sev *Category) Update(category model.Category) error {
	return sev.Update(category)
}

func (sev *Category) Delete(category model.Category) error {
	return sev.Repo.Delete(category)
}




