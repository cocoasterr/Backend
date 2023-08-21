package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"qty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (r *Product) TableName() string {
	return "product"
}

func (r *Product) GetNameField() []string {
	res := []string{"bapa", "Name", "Quantity", "CreatedAt", "UpdatedAt"}
	fmt.Println(reflect.TypeOf(res))
	return res
}

func main() {
	product := &Product{
		Id:        1,
		Name:      "Product Name",
		Quantity:  10,
		CreatedAt: "2023-08-12T15:30:00Z",
		UpdatedAt: "2023-08-12T15:45:00Z",
	}

	d := product.GetNameField()
	fmt.Println(d)
}
