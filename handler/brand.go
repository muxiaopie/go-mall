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
	Brand struct {
		Sev service.BrandService
	}
	BrandForm struct {
		Name string `valid:"required,unique(name)"`
		Desc string `valid:"required"`
		Logo string `valid:"required"`
		Sort int `valid:"required"`
	}
)


// 删除
func (h *Brand) Delete (c *gin.Context) error  {
	id := c.Param("id")
	brandId, err := strconv.Atoi(id)
	brand,err := h.Sev.Find(brandId)

	if err != nil {
		return err
	}
	err = h.Sev.Delete(brand)
	if err != nil {
		return err
	}
	return errno.Success
}

// 列表
func (h *Brand) List (c *gin.Context) error {
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
func (h *Brand) Update (c *gin.Context) error {
	var brandForm BrandForm
	if err := c.ShouldBindJSON(&brandForm); err != nil {
		return err
	}

	id := c.Param("id")

	brandId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		brandInfo,err := h.Sev.ByName(value)

		if err != nil {
			return true
		}

		if brandInfo.Id == uint(brandId)  {
			return true
		}

		if brandInfo.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err = govalidator.ValidateStruct(brandForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}
	brand,err := h.Sev.Find(brandId)
	if err != nil {
		return err
	}
	brand.Name = brandForm.Name
	brand.Desc = brandForm.Desc
	brand.Logo = brandForm.Logo
	brand.Sort = brandForm.Sort
	err = h.Sev.Update(brand)
	if err != nil {
		return err
	}
	return errno.Success
}

// 新增
func (h *Brand) Create (c *gin.Context) error {

	var brandForm BrandForm
	if err := c.ShouldBindJSON(&brandForm); err != nil {
		return err
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		brandInfo,err := h.Sev.ByName(value)

		if err != nil {
			return true
		}
		if brandInfo.Id > 0 {
			return false
		}
		return true
	})
	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\)$")

	// 验证
	_, err := govalidator.ValidateStruct(brandForm)
	if err != nil {
		return errno.ParameterError(err.Error())
	}

	brand,err := h.Sev.Create(model.Brand{
		Name:brandForm.Name,
		Desc:brandForm.Desc,
		Logo:brandForm.Logo,
		Sort:brandForm.Sort,
	})

	if err != nil {
		return err
	}

	c.JSON(statusOk,brand)
	return nil
}