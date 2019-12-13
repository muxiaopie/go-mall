package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
	"sync"
)


var categoryOnce sync.Once

// 获取repository
func NewCategoryRepository() (categoryRepository CategoryRepository) {
	categoryOnce.Do(func() {
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
	Pagination(page,limit int,maps map[string]interface{}) (*Pagination,error)
	ByName(name string) (model.Category,error)
}

type Category struct {
	DB *gorm.DB
}

// 增
func (repo *Category) Create(category model.Category) (model.Category,error) {
	return category,repo.DB.Create(&category).Error
}

func (repo *Category) ByName(name string) (model.Category,error) {
	var category model.Category
	return category,repo.DB.Where("name = ?",name).First(&category).Error
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
func (repo *Category) Pagination (page,limit int,maps map[string]interface{}) (*Pagination,error)  {
	var categorys []*model.Category
	/*p := NewPagination(page,limit)
	err := repo.DB.Limit(p.Page).Offset((p.Page - 1) * p.Limit).Find(&categoryList).Error
	if err != nil {
		return p,err
	}*/

	p := NewPagination(page,limit)
	p.Pagination(repo.DB).Where(maps).Find(&categorys)
	var count int
	repo.DB.Find(&model.Category{}).Where(maps).Count(&count)
	p.Total = count
	p.Items = categorys
	return p,nil
}


