package model

type Category struct {
	Model
	Name 	   string 	`json:"username" valid:"required,unique(name)"`
	Desc 	   string 	`json:"nickname" valid:"required"`
	Logo 	   string 	`json:"logo" valid:"required"`
	Sort 	   int 		`json:"sort" valid:"required"`
	Status     uint 	`json:"-" gorm:"default:1"`
}

func (c Category) TableName() string {
	return "category"
}