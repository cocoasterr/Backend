package main

import (
	"log"
	"net/http"

	"github.com/cocoaster/golang-jwt/controllers/authcontroller"
	productcontroller "github.com/cocoaster/golang-jwt/controllers/productController"
	"github.com/cocoaster/golang-jwt/models"
	"github.com/gorilla/mux"
)


func main()  {
	models.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/health-check", authcontroller.HealthCheck).Methods("GET")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("POST")
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.GetAllProduct).Methods("GET")


	log.Println(":8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}