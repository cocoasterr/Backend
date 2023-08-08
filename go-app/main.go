package main

import (
	"fmt"
	"log"

	"github.com/cocoasterr/go-app/helper"
	models "github.com/cocoasterr/go-app/model"
	"github.com/cocoasterr/go-app/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func main() {
	// var err error
	db := models.ConnectDB()
    // if err != nil {
		// 	log.Fatal("Error connecting to database: ", err)
		// }
    defer db.Close()
	
	router := gin.Default()
	orm.ConnectDB()


	
	router.GET("/create10k", func(c *gin.Context) {
		totalLoops := 10000

		for i := 0; i < totalLoops; i++ {
			query := "INSERT INTO products (name, quantity) VALUES ('selimut', 300)"
			_, err := db.Exec(query)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	router.GET("/products", getAllProduct)
	router.GET("/productsorm", ormGetAllProduct)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")
}

func getAllProduct(c *gin.Context) {
	var product models.Product
	helper.GetAllData(c, product.TableName())
}

func ormGetAllProduct(c *gin.Context){
	var product models.Product
	helper.OrmGetAllData(c, product.TableName())
}


// func getTodos(c *gin.Context) {
//     rows, err := models.DB.Query("SELECT * FROM products")
//     if err != nil {
//         c.JSON(500, gin.H{"error": "Failed to get todos"})
//         return
//     }
//     defer rows.Close()

//     todos := []gin.H{}
//     for rows.Next() {
//         var id, quantity int
//         var name  string
//         err := rows.Scan(&id, &name, &quantity)
//         if err != nil {
//             c.JSON(500, gin.H{"error": "Failed to scan todos"})
//             return
//         }
//         todos = append(todos, gin.H{"id": id, "name": name, "quantity": quantity})
//     }

//     c.JSON(200, gin.H{
// 		"message": "success",
// 		"data": todos,
// 		"total": len(todos),
// 	})
// }

func getTodo(c *gin.Context) {
    id := c.Param("id")
    var name string
    var quantity int

    err := models.DB.QueryRow("SELECT name, quantity FROM products WHERE id=$1", id).Scan(&name, &quantity)
    if err != nil {
        c.JSON(404, gin.H{"error": "Todo not found"})
        return
    }

    c.JSON(200, gin.H{"id": id, "name": name, "quantity": quantity})
}


func createTodo(c *gin.Context) {
	var todo gin.H
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	stmt, err := models.DB.Prepare("INSERT INTO todos (title, content) VALUES ($1, $2) RETURNING id")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	var id int
	err = stmt.QueryRow(todo["title"], todo["content"]).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}
	todo["id"] = id

	c.JSON(201, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo gin.H
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	stmt, err := models.DB.Prepare("UPDATE todos SET title=$1, content=$2 WHERE id=$3")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update todo"})
		return
	}

	_, err = stmt.Exec(todo["title"], todo["content"], id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(200, gin.H{"id": id, "title": todo["title"], "content": todo["content"]})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	stmt, err := models.DB.Prepare("DELETE FROM todos WHERE id=$1")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s is deleted", id)})
}
