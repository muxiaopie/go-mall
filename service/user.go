package service

import (
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/enum"
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
	FindFieldValue(field string,value string) (model.User,error)
	Create(user model.User) (model.User,error)

}

//
type User struct {
	Repo repository.UserRepository
}

func (ser *User) Find (action int, value string) (user model.User, err error) {
	if field ,ok := enum.FieldMap[action];ok {
		return ser.Repo.Find(field,value)
	}
	return
}

func (ser *User) FindFieldValue (field string, value string) (user model.User, err error) {
	return ser.Repo.Find(field,value)
}

func (ser *User) Create (user model.User) (model.User, error) {
	return ser.Repo.Create(user)
}




