package model

// 购物车表
type Cart struct {
	Model
	UserId 	   		   int 		`json:"userId"`
	SpuId 	   		   int 		`json:"spuId"`
	SkuId 	   		   int 		`json:"skuId"`
	Image 	   		   string 	`json:"image"`
	ProductName 	   string   `json:"productName"`
	Status     		   uint 	`json:"-" gorm:"default:1"`
}


type CartProperty struct {
	Model
}

func (c Cart) TableName() string {
	return "cart"
}