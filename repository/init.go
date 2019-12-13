package repository

import (
	"github.com/jinzhu/gorm"
)

// 分页
type Pagination struct {
	Items interface{} `json:"items"`
	Page 	int 	  `json:"page"`
	Total 	int 	  `json:"total"`
	Limit 	int 	  `json:"limit"`
}

// 构造一个分页数据
func NewPagination(page,limit int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit >= 100 || limit <= 0 {
		limit = 50
	}
	return &Pagination{
		Page:page,
		Limit:limit,
	}
}

// 分页
func (p *Pagination) Pagination (db *gorm.DB) *gorm.DB  {
	return db.Limit(p.Page).Offset((p.Page - 1) * p.Limit)
}
