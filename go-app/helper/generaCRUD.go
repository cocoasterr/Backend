package helper

import (
	"fmt"
	"strconv"

	models "github.com/cocoasterr/go-app/model"
	"github.com/gin-gonic/gin"
)
func generalResponse(c *gin.Context, data []interface{},info ...int){
	page := info[0]
	total := info[1]
	if data != nil{
		c.JSON(200, gin.H{
			"message": "success",
			"data":    data,
			"page":    page,
			"total":   total,
		})
	}else{
		c.JSON(200, gin.H{
			"message": "success!",
		})
	}
}

func GetAllData(c *gin.Context, tableName string) {
	page,_ := strconv.Atoi(c.Request.Header.Get("page"))
	limit,_ := strconv.Atoi(c.Request.Header.Get("limit"))
	query := fmt.Sprintf("SELECT * FROM %s limit %d offset %d", tableName, limit, page)
	rows, err := models.DB.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get data"})
		return
	}

	var total int
	query_total := fmt.Sprintf("SELECT COUNT(id) from %s", tableName)
	err = models.DB.QueryRow(query_total).Scan(&total)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get data"})
		return
    }
	defer rows.Close()
	data := []interface{}{}
	columnNames, _ := rows.Columns()

	for rows.Next() {
		dest := make([]interface{}, len(columnNames))
		for i := range columnNames {
			dest[i] = new(interface{})
		}
		err := rows.Scan(dest...)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to scan data"})
			return
		}

		// itemMap := make(map[string]interface{})
		// for i, colName := range columnNames {
		// 	itemMap[colName] = *(dest[i].(*interface{}))
		// }

		data = append(data, dest)
	}
	generalResponse(c, data, page, total)
}
// func GetAllData(c *gin.Context, tableName string) {
// 	page, _ := strconv.Atoi(c.Request.Header.Get("page"))
// 	limit, _ := strconv.Atoi(c.Request.Header.Get("limit"))

// 	// Persiapkan prepared statement untuk mendapatkan data
// 	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", tableName)
// 	stmt, err := models.DB.Prepare(query)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to prepare statement"})
// 		return
// 	}
// 	defer stmt.Close()

// 	// Eksekusi prepared statement dengan parameter
// 	rows, err := stmt.Query(limit, (page-1)*limit)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to get data"})
// 		return
// 	}
// 	defer rows.Close()

// 	var total int
// 	queryTotal := fmt.Sprintf("SELECT COUNT(id) FROM %s", tableName)
// 	err = models.DB.QueryRow(queryTotal).Scan(&total)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to get data"})
// 		return
// 	}

// 	data := []interface{}{}
// 	columnNames, _ := rows.Columns()

// 	for rows.Next() {
// 		dest := make([]interface{}, len(columnNames))
// 		for i := range columnNames {
// 			dest[i] = new(interface{})
// 		}
// 		err := rows.Scan(dest...)
// 		if err != nil {
// 			c.JSON(500, gin.H{"error": "Failed to scan data"})
// 			return
// 		}

// 		itemMap := make(map[string]interface{})
// 		for i, colName := range columnNames {
// 			itemMap[colName] = *(dest[i].(*interface{}))
// 		}

// 		data = append(data, itemMap)
// 	}

// 	generalResponse(c, data, page, total)
// }
