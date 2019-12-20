package admin

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
		Srv service.CategoryService
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
func (h *Category) Delete (c *gin.Context)   {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	category,err := h.Srv.Find(categoryId)

	if err != nil {
		panic(err)
	}
	err = h.Srv.Delete(category)
	if err != nil {
		panic(err)
	}
	panic(errno.Success)
}

// 列表
func (h *Category) List (c *gin.Context) {
	var page Page
	if err := c.ShouldBindJSON(&page); err != nil {
		panic(err)
	}

	// 验证
	_, err := govalidator.ValidateStruct(page)
	if err != nil {
		err := errno.ParameterError(err.Error())
		panic(err)
	}
	// where条件
	var maps map[string]interface{}

	// 搜索
	name := c.DefaultQuery("name", "")
	if name != "" {
		maps["name"] = name
	}

	pagination,err := h.Srv.Pagination(page.Page,page.Limit,maps)
	if err != nil {
		panic(err)
	}
	c.JSON(statusOk,pagination)
}

// 修改
func (h *Category) Update (c *gin.Context) {
	var categoryForm CategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		panic(err)
	}

	id := c.Param("id")

	categoryId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := h.Srv.ByName(value)

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
		err := errno.ParameterError(err.Error())
		panic(err)
	}
	category,err := h.Srv.Find(categoryId)
	if err != nil {
		panic(err)
	}
	category.Name = categoryForm.Name
	category.Desc = categoryForm.Desc
	category.Logo = categoryForm.Logo
	category.Sort = categoryForm.Sort
	err = h.Srv.Update(category)
	if err != nil {
		panic(err)
	}
	panic(errno.Success)
}

// 新增
func (h *Category) Create (c *gin.Context)  {

	var categoryForm CategoryForm
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		panic(err)
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		categoryInfo,err := h.Srv.ByName(value)

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
		err := errno.ParameterError(err.Error())
		panic(err)
	}

	category,err := h.Srv.Create(model.Category{
		Name:categoryForm.Name,
		Desc:categoryForm.Desc,
		Logo:categoryForm.Logo,
		Sort:categoryForm.Sort,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(statusOk,category)
}


