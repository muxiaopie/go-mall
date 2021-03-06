package model

type Sku struct {
	Model
	SpuId 		int 		 `json:"spu_id"`
	PropertyIds string 		 `json:"property_ids"`
	BrandUrl 	string 		 `json:"brand_url"`
	MainUrl     string 		 `json:"main_url"`
	SalesPrice 		int    	 `json:"sales_price"`
	ReferencePrice 	int    	 `json:"reference_price"`
	Sort 	   		int 	 `json:"sort"`
	Status     		uint 	 `json:"-" gorm:"default:1"`
}


type SkuImage struct {
	Model
	SpuId    int 		`json:"spuId"`
	ImageUrl string 	`json:"imageUrl"`
	Type     int 		`json:"type"`
	Status   int 		`json:"-" gorm:"default:1"`
}