package model

//订单详情

type OrderItem struct {
	OrderItemId string `json:"orderItemId"`
	OrderId string `json:"orderId"`
	ProductId string `json:"productId"`
	ProductName string `json:"productName"`
	ProductCoverImg string `json:"productCoverImg"`



}
