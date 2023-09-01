package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cocoasterr/gonofw/app/controllers"
	"github.com/cocoasterr/gonofw/app/services"
	"github.com/cocoasterr/gonofw/infra/postgre"
	PGRepositories "github.com/cocoasterr/gonofw/infra/postgre/repositories"
)
func main(){
	db,err:= postgre.PGConfig()
	if err !=nil{
		log.Fatal("error oi")
	}

	productRepository := PGRepositories.NewProductRepository(db)
	productService := services.NewProductSevice(productRepository)
	productController := controllers.NewProductController(productService)

	http.HandleFunc("/product", productController.CreateProduct)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})
	
	port:= 8080
	log.Printf("Running on localhost:%d \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}