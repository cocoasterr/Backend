package helper

import (
	"strconv"

	"github.com/cocoasterr/go-app/orm"
	"github.com/gin-gonic/gin"
)


func ormGeneralResponse(c *gin.Context, data []map[string]interface{}, page int, total int) {
	if data != nil {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    data,
			"page":    page,
			"total":   total,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success!",
		})
	}
}

func OrmGetAllData(c *gin.Context, tableName string) {
	page,_ := strconv.Atoi(c.Request.Header.Get("page"))
	limit,_ := strconv.Atoi(c.Request.Header.Get("limit"))
	offset := (page - 1) * limit

	var total int64
	if err := orm.DB.Table(tableName).Count(&total).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get data"})
		return
	}

	var data []map[string]interface{}
	if err := orm.DB.Table(tableName).Offset(offset).Limit(limit).Find(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get data"})
		return
	}

	ormGeneralResponse(c, data, page, int(total))
}