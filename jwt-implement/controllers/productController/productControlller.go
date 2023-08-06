package productcontroller

import (
	"net/http"

	"github.com/cocoaster/golang-jwt/helper"
)

type product struct{
	Id int
	Name string
	Quantity int
}


func GetAllProduct(w http.ResponseWriter, r *http.Request){
	products := []product{
		{
			Id: 1,
			Name: "pillow",
			Quantity: 100,
		},
		{
			Id: 2,
			Name: "Bad Cover",
			Quantity: 500,
		},
		{
			Id: 3,
			Name: "Spring Bed",
			Quantity: 50,
		},
	}

	response := map[string]interface{} {"status": "success","data": products, "total": len(products)}
	helper.ResponseJSON(w, http.StatusOK, response)
}