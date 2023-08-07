package productcontroller

import (
	"net/http"
	"strconv"

	"github.com/cocoaster/golang-jwt/helper"
	"github.com/cocoaster/golang-jwt/models"
	"github.com/gorilla/mux"
)


func CreateProduct(w http.ResponseWriter, r *http.Request){
	var userInput models.Product
	payload := helper.GetJson(&userInput, w, r).(*models.Product)
	helper.GeneralCreate(w,r, &payload)
}


func IndexProduct(w http.ResponseWriter, r *http.Request){
	var products []models.Product
	 
	helper.GeneralIndex(w, products)
}


func ProductById(w http.ResponseWriter, r *http.Request){
	var product []models.Product
	
	helper.GeneralFindById(w, r, product)
}


func DeleteProductById(w http.ResponseWriter, r *http.Request){
	var product models.Product

	helper.GeneralDelete(w,r, product)
}


func UpdateProduct(w http.ResponseWriter, r *http.Request){
	var product models.Product
	payload := helper.GetJson(&product, w, r).(*models.Product)

	id,_ := strconv.Atoi(mux.Vars(r)["id"])
	err := models.DB.First(&product, id)
	if err.Error != nil{
		response := map[string]interface{} {"status": "data not found"}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	//read json from body
	product.Name = payload.Name
	product.Quantity = payload.Quantity
	
	result:= models.DB.Save(&product)
	if result.Error != nil{
		response := map[string]interface{} {"status": "data update failed!"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{} {"status": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
