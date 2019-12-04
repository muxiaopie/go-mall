package model


type Spu struct {
	Model
	BrandId 		int
	CategoryId 		int
	Name 			string 	 `json:"name"`
	Desc 			string 	 `json:"desc"`
	SellingPoint 	string 	 `json:"selling_point"`
	Unit 			string 	 `json:"unit"`
	BannerUrl 		string 	 `json:"banner_url"`
	MainUrl 		string 	 `json:"main_url"`
	Specification   string 	 `json:"specification"`
	SalesPrice 		int    	 `json:"sales_price"`
	ReferencePrice 	int    	 `json:"reference_price"`
	Sort 	   		int 	 `json:"sort"`
	Status     		uint 	 `json:"-" gorm:"default:1"`
	Brand 			Brand 	 `json:"brand"`
	Category        Category `json:"category"`
}
