package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/enum"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"
	"github.com/asaskevich/govalidator"
	"github.com/muxiaopie/go-mall/service"
	"github.com/muxiaopie/go-mall/util"
	"github.com/spf13/viper"
	"regexp"
)

type (

	User struct {
		Sev service.UserService
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



func (u *User) User (c *gin.Context) error {
	uid,err := userId(c)
	if err != nil {
		return err
	}
	id := fmt.Sprintf("%d",uid)
	user,_ := u.Sev.Find(enum.ID,id)
	c.JSON(200,user)
	return nil
}

// 登陆接口
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

func (u *User) Register(c *gin.Context) error  {


	var registerForm RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		return err
	}

	/*govalidator.TagMap["unique"] = govalidator.Validator(func(username string) bool {
		user,err := u.Sev.Find(enum.USERNAME,username)
		if err != nil {
			return true
		}
		if user.Id > 0 {
			return false
		}
		return trueunique
	})*/

	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		fmt.Println(str)
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	_, err := govalidator.ValidateStruct(registerForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}

	return errno.Success
	return nil


}

