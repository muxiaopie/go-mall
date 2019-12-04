package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"

)



// 获取repository
func NewCategoryRepository() (categoryRepository CategoryRepository) {
	once.Do(func() {
		categoryRepository = &Category{
			DB:bootstrap.Bootstrap.DB,
		}
	})
	return categoryRepository
}


type CategoryRepository interface {
	Create(category model.Category)(model.Category,error)
	Find(id int)(model.Category,error)
	Update(category model.Category) error
	Delete(category model.Category) error
	Pagination(page,limit int) (*Pagination,error)
}

type Category struct {
	DB *gorm.DB
}

// 增
func (repo *Category) Create(category model.Category) (model.Category,error) {
	return category,repo.DB.Create(&category).Error
}

// 查找
func (repo *Category) Find(id int) (category model.Category,err error) {
	return category,repo.DB.First(&category,id).Error
}

// 修改
func (repo *Category) Update(category model.Category) error {
	return repo.DB.Save(&category).Error
}

// 删除
func (repo *Category) Delete (category model.Category) error {
	return repo.DB.Delete(&category).Error
}

// 分页
func (repo *Category) Pagination (page,limit int) (*Pagination,error)  {
	var categoryList []*model.Category
	p := NewPagination(page,limit)
	err := repo.DB.Limit(p.Page).Offset((p.Page - 1) * p.Limit).Find(&categoryList).Error
	if err != nil {
		return p,err
	}
	var count int
	repo.DB.Find(&categoryList).Count(&count)
	p.Total = count
	p.Items = categoryList
	return p,nil
}

