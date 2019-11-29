package bootstrap

import "github.com/gin-gonic/gin"

func LoadGin() func(*Bootstrapper) {
	return func(bootstrap *Bootstrapper) {
		Route := gin.New()
		//boot.Router = gin.Default()
		bootstrap.Router = Route
	}
}

