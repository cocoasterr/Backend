package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cocoasterr/net_http/app/controllers"
	"github.com/cocoasterr/net_http/app/services"
	PGConfig "github.com/cocoasterr/net_http/infra/sql/postgres"
	PGRepository "github.com/cocoasterr/net_http/infra/sql/postgres/repository"
)


func main(){
	db, err:= PGConfig.ConnectPG()
	if err != nil{
		log.Fatal("Failed to Connect DB")
	}

	ProductRepo:= PGRepository.NewProductRepository(db)
	ProductService:= services.NewProductService(*ProductRepo)
	Productcontroller := controllers.NewProductController(*ProductService)

	http.HandleFunc("/api/create-product", Productcontroller.CreateProductController)
	http.HandleFunc("/api/index-product", Productcontroller.IndexProdcuctController)
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		resp:= map[string]interface{}{"Message": "page not found!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Running on port :8080")
	http.ListenAndServe(":8080", nil)
}