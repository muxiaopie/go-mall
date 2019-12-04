package model

type Category struct {
	Model
	Name 	   string 	`json:"username"`
	Desc 	   string 	`json:"nickname"`
	Logo 	   string 	`json:"logo"`
	Sort 	   int 		`json:"sort"`
	Status     uint 	`json:"-" gorm:"default:1"`
}
