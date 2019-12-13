package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/model"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/service"
	"regexp"
	"strconv"
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
	// 分页
	Page struct {
		Page,Limit int
	}
)

// 删除
func (h *Category)Delete(c *gin.Context) error  {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	category,err := h.Sev.Find(categoryId)

	if err != nil {
		return err
	}
	err = h.Sev.Delete(category)
	if err != nil {
		return err
	}
	return errno.Success
}

// 列表
func (h *Category) List(c *gin.Context) error {
	var page Page
	if err := c.ShouldBindJSON(&page); err != nil {
		return err
	}

	// 验证
	_, err := govalidator.ValidateStruct(page)
	if err != nil {
		return errno.ParameterError(err.Error())
	}
	// where条件
	var maps map[string]interface{}

	// 搜索
	name := c.DefaultQuery("name", "")
	if name != "" {
		maps["name"] = name
	}

	pagination,err := h.Sev.Pagination(page.Page,page.Limit,maps)
	if err != nil {
		return err
	}
	c.JSON(statusOk,pagination)

	return nil
}

// 修改
func (h *Category) Update(c *gin.Context) error {
	var categoryForm CategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		return err
	}

	id := c.Param("id")

	categoryId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := h.Sev.ByName(value)

		if err != nil {
			return true
		}

		if categoryInfo.Id == uint(categoryId)  {
			return true
		}

		if categoryInfo.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err = govalidator.ValidateStruct(categoryForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}
	category,err := h.Sev.Find(categoryId)
	if err != nil {
		return err
	}
	category.Name = categoryForm.Name
	category.Desc = categoryForm.Desc
	category.Logo = categoryForm.Logo
	category.Sort = categoryForm.Sort
	err = h.Sev.Update(category)
	if err != nil {
		return err
	}
	return errno.Success
}

// 新增
func (h *Category) Create (c *gin.Context) error {

	var categoryForm CategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := h.Sev.ByName(value)

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

	category,err := h.Sev.Create(model.Category{
		Name:categoryForm.Name,
		Desc:categoryForm.Desc,
		Logo:categoryForm.Logo,
		Sort:categoryForm.Sort,
	})

	if err != nil {
		return err
	}

	c.JSON(statusOk,category)
	return nil
}


