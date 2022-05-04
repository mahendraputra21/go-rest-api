package model

type Order struct {
	ID          int16  `json:"id"`
	ProductId   int16  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int32  `json:"price"`
	Qty         int8   `json:"qty"`
	GrandTotal  int32  `json:"grand_total"`
}
