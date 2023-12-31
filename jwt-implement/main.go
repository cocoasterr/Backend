package main

import (
	"log"
	"net/http"

	"github.com/cocoaster/golang-jwt/controllers/authcontroller"
	productcontroller "github.com/cocoaster/golang-jwt/controllers/productController"
	"github.com/cocoaster/golang-jwt/middlewares"
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
	api.HandleFunc("/product", productcontroller.CreateProduct).Methods("POST")
	api.HandleFunc("/products", productcontroller.IndexProduct).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.ProductById).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.UpdateProduct).Methods("PUT")
	api.HandleFunc("/product/{id}", productcontroller.DeleteProductById).Methods("DELETE")

	api.Use(middlewares.JWTMiddleware)

	log.Println(":8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}