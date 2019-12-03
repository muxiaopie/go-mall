package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/enum"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"

	"github.com/asaskevich/govalidator"

	"github.com/muxiaopie/go-mall/service"
	"github.com/muxiaopie/go-mall/util"
	"github.com/spf13/viper"
)

type User struct {
	Sev service.UserService
}

type LoginForm struct {
	UserName string `valid:"required" json:"username" form:"username"`
	PassWord string `valid:"required" json:"password" form:"password"`
}

func (u *User) User (c *gin.Context) error {
	id := c.Param("id")
	user,_ := u.Sev.Find(enum.ID,id)
	c.JSON(200,user)
	return nil
}


func (u *User) Login(c *gin.Context) error {
	var loginForm LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(loginForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}

	user,err := u.Sev.Find(enum.USERNAME,loginForm.UserName)
	if err!= nil {
		return err
	}
	if util.CheckPasswordHash(loginForm.PassWord,user.Password) {
		token,_ := jwt.GenerateToken(user.Id)
		c.JSON(200, struct {
			Token 	   string `json:"token"`
			ExpireTime string `json:"expire_time"`
		}{
			Token:token,
			ExpireTime:viper.GetString("expireTime"),
		})
		return nil
	}else{
		return errors.New("密码不正确,请重试")
	}

}

