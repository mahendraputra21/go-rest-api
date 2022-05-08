package model

type Product struct {
	ID          int16  `json:"id"`
	ProductName string `json:"product_name"`
	BrandId     int8   `json:"brand_id"`
	BrandName   string `json:"brand_name"`
	Price       int32  `json:"price"`
}
