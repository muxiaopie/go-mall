package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/enum"
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
	Find(action int,value string) (model.User, error)
}

type User struct {
	DB *gorm.DB
}

// 查询服务
func (repo *User) Find(action int,value string) (user model.User, err error) {
	if field ,ok := enum.FieldMap[action];ok {
		err = repo.DB.Where(fmt.Sprintf("%s = ?",field), value).First(&user).Error
		return
	}
	return
}

