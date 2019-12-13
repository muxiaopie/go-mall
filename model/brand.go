package model

type Brand struct {
	Model
	Name 	   string 	`json:"username"`
	Desc 	   string 	`json:"nickname"`
	Logo 	   string 	`json:"logo"`
	Sort 	   int 		`json:"sort"`
	Status     uint 	`json:"-" gorm:"default:1"`
	CategoryId int  	`json:"category_id"`
}

func (b Brand) TableName() string {
	return "brand"
}