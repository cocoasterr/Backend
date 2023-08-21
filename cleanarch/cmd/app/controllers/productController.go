// controller.go

package controller

import (
	"net/http"
	"strconv"

	"github.com/cocoasterr/cleanarch/domain"
	"github.com/cocoasterr/cleanarch/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
    productUC service.ProductUseCase
}

func NewProductController(productUC service.ProductUseCase) *Controller {
    return &Controller{
        productUC: productUC,
    }
}

func (con *Controller) CreateProductHandler(c *gin.Context) {
    var product domain.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := con.productUC.CreateProduct(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

func (con *Controller) GetAllProductHandler(c *gin.Context) {
    limit, _:= strconv.Atoi(c.Request.Header.Get("limit"))
    page, _ := strconv.Atoi(c.Request.Header.Get("page"))

    var product domain.Product
    res, err := con.productUC.GetAllProduct(&product, page,limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Index failed!"})
        return
    }
    c.JSON(http.StatusOK, res)
}

func (con *Controller) GetProductById(c *gin.Context){
    id, _ := strconv.Atoi(c.Param("id"))
    
    var product domain.Product

    res,err := con.productUC.GetProductById(&product, id)
    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "get by id failed!"})
        return
    }
    c.JSON(http.StatusOK, res)
}