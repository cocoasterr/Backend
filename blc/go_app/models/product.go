package models

type Product struct {
	Id          string `json:"id"`
	ProductName string `json:"product_name"`
	Qty         string `json:"qty"`
	ExtTime
	CreatedBy
}
