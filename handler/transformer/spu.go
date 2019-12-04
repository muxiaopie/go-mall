package transformer

import "github.com/muxiaopie/go-mall/model"

// spu信息
type SpuTransformer struct {
	Id              int 	 `json:"id"`
	Name 			string 	 `json:"name"`
	Desc 			string 	 `json:"desc"`
	SellingPoint 	string 	 `json:"sellingPoint"`
	Unit 			string 	 `json:"unit"`
	BannerUrl 		[]string `json:"bannerUrl"`
	MainUrl 		[]string `json:"mainUrl"`
	Specification   string 	 `json:"specification"`
	SalesPrice 		int    	 `json:"salesPrice"`
	ReferencePrice 	int    	 `json:"referencePrice"`

}

// 格式化数据
func NewSpuTransformer(spu model.Spu)  {

}
