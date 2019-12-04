package model

type Property struct {
	Model
	Name 	   string 	`json:"username"`
	Desc 	   string 	`json:"nickname"`
	Logo 	   string 	`json:"logo"`
	Sort 	   int 		`json:"sort"`
	Status     uint 	`json:"-" gorm:"default:1"`
}


type PropertyValue struct {
	Model
	PropertyId int
	Value 	   string 	`json:"username"`
	Desc 	   string 	`json:"nickname"`
	Status     uint 	`json:"-" gorm:"default:1"`
}


type SpuSkuPropertyMap struct {
	Model
	SpuId 	int 			 `json:"spu_id"`
	SkuId 	int 			 `json:"sku_id"`
	PropertyId int 			 `json:"property_id"`
	PropertyName string 	 `json:"property_name"`
	PropertyValueId int 	 `json:"property_value_id"`
	PropertyValueName string `json:"property_value_name"`
	Status     uint 		 `json:"-" gorm:"default:1"`
}