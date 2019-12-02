package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/service"
	"strconv"
)

type User struct {
	Sev service.UserService
}

func (u *User) User (c *gin.Context)  {
	id := c.Param("id")
	idInt ,_ := strconv.Atoi(id)
	user,_ := u.Sev.Find(uint(idInt))
	c.JSON(200,user)
}