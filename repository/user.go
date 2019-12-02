package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
	"sync"
)

var once sync.Once


// 获取repository
func NewUserRepository() (userRepository UserRepository) {
	once.Do(func() {
		userRepository = &User{
			DB:bootstrap.Bootstrap.DB,
		}
	})
	return userRepository
}


type UserRepository interface {
	Find(id uint) (*model.User, error)
}

type User struct {
	DB *gorm.DB
}

// 查询服务
func (repo *User) Find(id uint) (*model.User, error) {
	user :=  &model.User{}
	user.Id = uint(id)
	if err := repo.DB.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

