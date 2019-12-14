package service

import (
	"errors"
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/enum"
	"github.com/muxiaopie/go-mall/repository"
)

// 获取服务
func NewUserService() (userService UserService) {
	return &User{
		Repo: repository.NewUserRepository(),
	}
}

// service 服务
type UserService interface {
	Find(action int, value string) (model.User, error)
	FindFieldValue(field string, value string) (model.User, error)
	Create(user model.User) (model.User, error)
	Pagination(page, limit int) *repository.Pagination
}

//
type User struct {
	Repo repository.UserRepository
}

func (srv *User) Find(action int, value string) (user model.User, err error) {
	if field, ok := enum.FieldMap[action]; ok {
		return srv.Repo.Find(field, value)
	}
	return user,errors.New("参数有误")
}

func (srv *User) FindFieldValue(field string, value string) (user model.User, err error) {
	return srv.Repo.Find(field, value)
}

func (srv *User) Create(user model.User) (model.User, error) {
	return srv.Repo.Create(user)
}

func (srv *User) Pagination(page, limit int) *repository.Pagination {
	var maps map[string]interface{}
	return srv.Repo.Pagination(page, limit, maps)
}
