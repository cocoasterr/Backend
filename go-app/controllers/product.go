package controllers

import (
	"github.com/cocoasterr/go-app/helper"
	models "github.com/cocoasterr/go-app/model"
	"github.com/gin-gonic/gin"
)

//==================================================
//databasae/sql controllers

func GetAllProduct(c *gin.Context) {
	var product models.Product
	helper.GetAllData(c, product.TableName())
}


//==================================================
//orm controllers

func OrmGetAllProduct(c *gin.Context){
	var product models.Product
	helper.OrmGetAllData(c, product.TableName())
}