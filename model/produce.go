package model

type Product struct {

	ProductID 				string 		`json:"productId" gorm:"column:product_id"`
	ProductName 			string 		`json:"productName" gorm:"column:product_name"`
	ProductIntro 			string 		`json:"productIntro" gorm:"column:product_intro"`
	CategoryID 				string 		`json:"categoryId" gorm:"column:category_id"`
	ProductCoverImg	 		string 		`json:"productCoverImg" gorm:"column:product_cover_img"`
	ProductBanner 			string 		`json:"productBanner" gorm:"column:product_banner"`
	OriginalPrice 			int 		`json:"originalPrice" gorm:"column:original_price"`
	SellingPrice 			int 		`json:"sellingPrice" gorm:"column:selling_price"`
	StockNum 				int 		`json:"stockNum" gorm:"column:stock_num"`
	Tag 					string 		`json:"tag" gorm:"column:tag"`
	SellStatus 				int 		`json:"sellStatus" gorm:"column:sell_status"`
	CreateUser 				string 		`json:"createUser" gorm:"column:create_user"`
	UpdateUser 				string 		`json:"updateUser" gorm:"update_user"`
	ProductDetailContent 	string 		`json:"productDetailContent" gorm:"column:product_detail_content"`
	IsDeleted 				bool 		`json:"isDeleted" gorm:"column:is_deleted"`

}
