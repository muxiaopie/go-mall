package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/repository"
	"sync"
)


var	once sync.Once

// 获取服务
func NewUserService() (userService UserService) {
	once.Do(func() {
		userService = &User{
			Repo:repository.NewUserRepository(),
		}
	})
	return userService
}

// service 服务
type UserService interface {
	Find (action int, value string) (model.User, error)
}

//
type User struct {
	Repo repository.UserRepository
}

func (ser *User) Find (action int, value string) (model.User, error) {
	return ser.Repo.Find(action,value)
}




