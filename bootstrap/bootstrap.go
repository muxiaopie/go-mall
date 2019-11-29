package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"sync"
)

var once sync.Once
var Bootstrap *Bootstrapper

// 引导程序
type Bootstrapper struct {
	Router *gin.Engine
	DB     *gorm.DB
}

// 程序初始化
func init() {
	once.Do(func() {
		Bootstrap = &Bootstrapper{}
	})
}

// Use : 驱动中间件
func (bootstrap *Bootstrapper) Use(options ...func(*Bootstrapper)) {
	for _, option := range options {
		option(bootstrap)
	}
}


func (bootstrap *Bootstrapper) Run(addr ...string)  {
	bootstrap.Router.Run(addr ...)
}
