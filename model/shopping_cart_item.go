package model

type ShoppingCartItem struct {

	CartItemId string `json:"cartItemId"`
	UserId string `json:"userId"`
	ProductId string `json:"productId"`
	ProductCount string `json:"productCount"`
	IsDelete string `json:"isDelete"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
