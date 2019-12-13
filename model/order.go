package model

import "github.com/muxiaopie/go-mall/util"

type Order struct {
	Model
	OrderNum   string 	 	   `json:"orderNum"`
	AddressId  int 	  	 	   `json:"addressId"`
	UserId     int 	  	 	   `json:"userId"`
	ItemsTotal int 	  	 	   `json:"items_total"`
	AdjustmentsTotal int 	   `json:"adjustments_total"`
	PayTotal int 		 	   `json:"pay_total"`
	Status int 			 	   `json:"status"`
	PaymentStatus int 	 	   `json:"payment_status"`
	ShipmentStatus int 	 	   `json:"shipment_status"`
	UserIp int 		     	   `json:"user_ip"`
	PayedAt 	util.LocalTime `json:"payed_at"`
	ConfirmedAt util.LocalTime `json:"confirmed_at"`
	ReviewedAt  util.LocalTime `json:"reviewed_at"`
	FulfilledAt util.LocalTime `json:"fulfilled_at"`
	Rest		string 		   `json:"rest"`
	Type        int 		   `json:"type"`

}

// 订单sku
type OrderItem struct {
	Model
	OrderId int
	SkuId int
	SpuId int
	Num int
	UnitsTotal int
	AdjustmentsTotal int
	Total int
	UnitPrice int
	Rest string
	Snapshot string
	Type int
}


type OrderItemUnit struct {
	Model
	ItemId 	   int
	ShipmentId int
	AdjustmentsTotal int
}

type Adjustment struct {
	OrderId int
	OrderItemId int
	OrderItemUnitId int
	Type int
	Label string
	OriginCode string
	Included int
	Amount int
}

type Shipments struct {

}

type Payments struct {

}