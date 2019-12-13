package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
)

// var userOnce sync.Once
// 获取repository
func NewUserRepository() (userRepository UserRepository) {
	return &User{
		DB:bootstrap.Bootstrap.DB,
	}
	/*userOnce.Do(func() {
		userRepository = &User{
			DB:bootstrap.Bootstrap.DB,
		}
	})
	return userRepository*/
}

// user
type UserRepository interface {
	Find(field ,value string) (model.User, error)
	Create(model.User)(model.User,error)
	Pagination(page,limit int,maps map[string]interface{}) *Pagination
}

type User struct {
	DB *gorm.DB
}

// 查找
func (repo *User) Find(field string,value string) (user model.User, err error) {
	err = repo.DB.Where(fmt.Sprintf("%s = ?",field), value).First(&user).Error
	return
}

// 创建
func (repo *User) Create(user model.User) (model.User,error) {
	return user,repo.DB.Create(&user).Error
}

// 分页
func (repo *User) Pagination (page,limit int,maps map[string]interface{}) *Pagination  {
	var users []*model.User
	var count int

	p := NewPagination(page,limit)
	p.Pagination(repo.DB).Where(maps).Find(&users)
	repo.DB.Model(&model.User{}).Where(maps).Count(&count)
	p.Items = users
	p.Total = count
	return p
}

