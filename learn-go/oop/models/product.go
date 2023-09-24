package models

type Product struct {
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	UpdatedBy   GetNamePerson
	setTime
}

func NewProduct(productName, updatedBy string, qty int) *Product {
	return &Product{
		ProductName: productName,
		Qty:         qty,
		UpdatedBy:   *SetName(updatedBy),
		setTime:     *NewSetTime(),
	}
}
