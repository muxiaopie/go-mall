package handler

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/service"
	"regexp"
)

type (
	Category struct {
		Sev service.CategoryService
	}
	CategoryForm struct {
		Name string `valid:"required,unique(name)"`
		Desc string `valid:"required"`
		Logo string `valid:"required"`
		Sort int `valid:"required"`
	}

	UpdateCategoryForm struct {
		Id int `valid:"required"`
		CategoryForm
	}

	Page struct {
		Page,Limit int
	}

)



func (category *Category) List(c *gin.Context) error {
	var page Page
	if err := c.ShouldBindJSON(&page); err != nil {
		return err
	}

	// 验证
	_, err := govalidator.ValidateStruct(page)
	if err != nil {
		return errno.ParameterError(err.Error())
	}
	var maps map[string]interface{}
	pagination,err := category.Sev.Pagination(page.Page,page.Limit,maps)
	if err != nil {
		return err
	}
	c.JSON(statusOk,pagination)

	return nil
}

// 修改
func (category *Category) Update(c *gin.Context) error {
	var categoryForm UpdateCategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := category.Sev.ByName(value)
		id := uint(categoryForm.Id)

		if err != nil {
			return true
		}

		if categoryInfo.Id == id  {
			return true
		}

		if categoryInfo.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err := govalidator.ValidateStruct(categoryForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}

	categoryModel,err := category.Sev.Find(categoryForm.Id)
	fmt.Println(categoryModel)

	if err != nil {
		return err
	}
	categoryModel.Name = categoryForm.Name
	categoryModel.Desc = categoryForm.Desc
	categoryModel.Logo = categoryForm.Logo
	categoryModel.Sort = categoryForm.Sort

	err = category.Sev.Update(categoryModel)

	if err != nil {
		return err
	}

	return errno.Success
}

// 新增
func (category *Category) Create (c *gin.Context) error {

	var categoryForm CategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := category.Sev.ByName(value)

		if err != nil {
			return true
		}
		if categoryInfo.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err := govalidator.ValidateStruct(categoryForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}

	categoryModel,err := category.Sev.Create(model.Category{
		Name:categoryForm.Name,
		Desc:categoryForm.Desc,
		Logo:categoryForm.Logo,
		Sort:categoryForm.Sort,
	})

	if err != nil {
		return err
	}

	c.JSON(statusOk,categoryModel)
	return nil
}


