package handler

import "github.com/gin-gonic/gin"

type (
	Category struct {
		Sev
	}
	CategoryForm struct {
		Name string `valid:"required,unique(name)"`
		Desc string `valid:"email,required"`
		Logo string `valid:"required"`
		Sort string `valid:"required"`
	}

)

func (category Category) Create (c *gin.Context) error {

	return nil
}
