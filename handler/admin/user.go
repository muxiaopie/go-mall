package admin

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/enum"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"
	"github.com/muxiaopie/go-mall/service"
	"github.com/muxiaopie/go-mall/util"
	"github.com/spf13/viper"
	"regexp"
)

type (
	User struct {
		Srv service.UserService
	}
	LoginForm struct {
		UserName string `valid:"required" json:"username" form:"username"`
		PassWord string `valid:"required" json:"password" form:"password"`
	}
	RegisterForm struct {
		UserName string `valid:"required,unique(username)"`
		Email    string `valid:"email,required,unique(email)"`
		PassWord string `valid:"required"`
	}
)

// 用户信息
func (u *User) Users (c *gin.Context) {
	P := u.Srv.Pagination(1,1)

	c.JSON(statusOk,P)
}

// 用户信息
func (u *User) User (c *gin.Context) {
	uid, err := userId(c)
	if err != nil {
		panic(err)
	}
	id := fmt.Sprintf("%d", uid)
	user, err := u.Srv.Find(enum.ID, id)
	if err != nil {
		panic(errno.NotFound)
	}
	c.JSON(statusOk, user)
}

// 登陆接口
func (u *User) Login (c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		panic(err)
	}
	_, err := govalidator.ValidateStruct(loginForm)
	if err != nil {
		err := errno.ParameterError(err.Error())
		panic(err)
	}
	user, err := u.Srv.Find(enum.USERNAME, loginForm.UserName)
	if err != nil {
		panic(err)
	}

	if user.Admin != 1 {
		panic(errors.New("你没有登陆权限"))
	}

	if util.CheckPasswordHash(loginForm.PassWord, user.Password) {
		token, _ := jwt.GenerateToken(user.Id)
		c.JSON(statusOk, struct {
			Token      string `json:"token"`
			ExpireTime string `json:"expire_time"`
		}{
			Token:      token,
			ExpireTime: viper.GetString("expireTime"),
		})
	} else {
		panic(errors.New("密码不正确,请重试"))
	}
}

// 注册接口
func (u *User) Register (c *gin.Context) {
	var registerForm RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		panic(err)
	}
	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		field := params[0]
		user, err := u.Srv.FindFieldValue(field, value)
		if err != nil {
			return true
		}
		if user.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err := govalidator.ValidateStruct(registerForm)
	if err != nil {
		err := errno.ParameterError(err.Error())
		panic(err)
	}

	// 加密
	password, err := util.HashPassword(registerForm.PassWord)
	if err != nil {
		err := errno.OtherError("注册失败")
		panic(err)
	}
	// 入库
	user := model.User{
		Username: registerForm.UserName,
		Email:    registerForm.Email,
		Password: password,
	}
	userInfo, err := u.Srv.Create(user)
	if err != nil {
		panic(err)
	}
	c.JSON(statusOk, userInfo)
}


// 修改的接口
