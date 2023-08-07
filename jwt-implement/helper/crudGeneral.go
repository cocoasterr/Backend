package helper

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/cocoaster/golang-jwt/models"
	"github.com/gorilla/mux"
)


func GeneralCreate(w http.ResponseWriter,r *http.Request, data interface{}){
	if err := models.DB.Create(data).Error; err != nil{
		response := map[string]interface{}{"message": string(err.Error())}
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "product created!"}
	ResponseJSON(w, http.StatusOK, response)
}

func GeneralIndex(w http.ResponseWriter, data interface{}){
	models.DB.Find(&data)
	total := reflect.ValueOf(data).Len()
	response := map[string]interface{} {"status": "success","data": data, "total": total}
	ResponseJSON(w, http.StatusOK, response)
}

func GeneralFindById(w http.ResponseWriter, r *http.Request, data interface{}){
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := models.DB.First(&data, id).Error; err !=nil{
		response := map[string]interface{} {"status": "data not found"}
		ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	response := map[string]interface{} {"status": "success","data": data, "total": 1}
	ResponseJSON(w, http.StatusOK, response)
}

func GeneralDelete(w http.ResponseWriter, r *http.Request, data interface{}){
	id ,_ := strconv.Atoi(mux.Vars(r)["id"])

	if err := models.DB.Delete(&data, id).Error; err != nil{
		response := map[string]interface{} {"status": "data not found"}
		ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	response := map[string]interface{} {"status": "success"}
	ResponseJSON(w, http.StatusOK, response)
}