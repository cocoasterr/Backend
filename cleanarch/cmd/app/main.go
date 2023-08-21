package main

import (
	"net/http"

	controller "github.com/cocoasterr/cleanarch/cmd/app/controllers"
	configPostgres "github.com/cocoasterr/cleanarch/infra/database/postgres"
	"github.com/cocoasterr/cleanarch/infra/database/postgres/repository"
	"github.com/cocoasterr/cleanarch/service"
	"github.com/gin-gonic/gin"
)

func main(){
	db := configPostgres.ConnectDB()
    productRepo := repository.NewProductRepository(db)
    productUC := service.NewProductUseCase(*productRepo)
	controller := controller.NewProductController(*productUC)


	router := gin.Default()
	api := router.Group("/api")
	api.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success!",
		})
	})

	api.POST("/create-product", controller.CreateProductHandler)
	api.GET("/getall-product", controller.GetAllProductHandler)
	api.GET("/getbyid-product/:id", controller.GetProductById)

	router.Run(":8081")

}