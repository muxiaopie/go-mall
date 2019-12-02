package model

import "github.com/muxiaopie/go-mall/util"

type Model struct {
	Id        uint 		 	  `gorm:"primary_key" json:"id"`  // 主键
	CreatedAt util.LocalTime  `json:"created_at"`			  // 创建时间
	UpdatedAt util.LocalTime  `json:"updated_at"`			  // 更新时间
	DeletedAt util.LocalTime  `json:"deleted_at"`			  // 删除时间
}