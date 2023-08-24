package domain

type DataModel interface {
    TableName() string
	NewProduct(p Product) Product
}