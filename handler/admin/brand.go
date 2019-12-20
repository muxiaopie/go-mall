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
	Brand struct {
		Srv service.BrandService
	}
)


// 删除
func (h *Brand) Delete (c *gin.Context)  {
	id := c.Param("id")
	brandId, err := strconv.Atoi(id)
	brand,err := h.Srv.Find(brandId)

	if err != nil {
		panic(err)
	}
	err = h.Srv.Delete(brand)
	if err != nil {
		panic(err)
	}
	panic(errno.Success)
}

// 列表
func (h *Brand) List (c *gin.Context)  {
	var page Page
	if err := c.ShouldBindJSON(&page); err != nil {
		panic(err)
	}

	// 验证
	_, err := govalidator.ValidateStruct(page)
	if err != nil {
		panic(errno.ParameterError(err.Error()))
	}
	// where条件
	var maps map[string]interface{}

	// 搜索
	name := c.DefaultQuery("name", "")
	if name != "" {
		maps["name"] = name
	}

	pagination,err := h.Srv.Pagination(page.Page, page.Limit,maps)
	if err != nil {
		panic(err)
	}
	c.JSON(statusOk,pagination)

}

// 修改
func (h *Brand) Update (c *gin.Context) {
	var brandForm model.Brand
	if err := c.ShouldBindJSON(&brandForm); err != nil {
		panic(err)
	}

	id := c.Param("id")

	brandId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		brandInfo,err := h.Srv.ByName(value)

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
		err := errno.ParameterError(err.Error())
		panic(err)
	}
	brand,err := h.Srv.Find(brandId)
	if err != nil {
		panic(err)
	}
	brand.Name = brandForm.Name
	brand.Desc = brandForm.Desc
	brand.Logo = brandForm.Logo
	brand.Sort = brandForm.Sort
	err = h.Srv.Update(brand)
	if err != nil {
		panic(err)
	}
	panic(errno.Success)
}

// 新增
func (h *Brand) Create (c *gin.Context) {

	var brandForm model.Brand
	if err := c.ShouldBindJSON(&brandForm); err != nil {
		panic(err)
	}

	// 验证唯一性
	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(value string, params ...string) bool {
		brandInfo,err := h.Srv.ByName(value)

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
		err := errno.ParameterError(err.Error())
		panic(err)
	}

	brand,err := h.Srv.Create(model.Brand{
		Name:brandForm.Name,
		Desc:brandForm.Desc,
		Logo:brandForm.Logo,
		Sort:brandForm.Sort,
	})

	if err != nil {
		panic(err)
	}
	c.JSON(statusOk,brand)
}