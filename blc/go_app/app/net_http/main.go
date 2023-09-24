package main

import (
	"log"
	"net/http"

	"github.com/cocoasterr/Backend/go_app/app/net_http/controllers"
	PG "github.com/cocoasterr/Backend/go_app/arch/db/pg/gorm/config"
	PGRepo "github.com/cocoasterr/Backend/go_app/arch/db/pg/gorm/repository"
	service "github.com/cocoasterr/Backend/go_app/service"
	"gorm.io/gorm"
)

func lala(*gorm.DB) {

}
func main() {
	db, err := PG.PGConfigGorm()
	if err != nil {
		return
	}

	lala(db)
	productRepo := PGRepo.NewProductRepository(db)
	productService := service.NewProductService(productRepo.Repository)
	productController := controllers.NewProductController(*productService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
	})

	log.Println("Connected to port :8080")
	http.ListenAndServe(":8080", nil)
}
