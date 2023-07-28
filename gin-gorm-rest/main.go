package main

import (
	"Golang_API/gin-gorm-rest/models"
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var router = gin.New() // Thay đổi thông tin kết nối của bạn tại đây
	connectionInfo := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	defer db.Close()

	// Tạo câu lệnh SQL để tạo CSDL
	//createDBQuery := `
	//	CREATE DATABASE gin_gom_rest
	//`

	//_, err = db.Exec(createDBQuery)
	if err != nil {
		log.Fatal("Failed to create the database: ", err)
	}

	fmt.Println("Database gin_gom_rest created successfully")
	//Khởi tạo router Gin
	router = gin.Default()

	// ĐỊnh nghĩa endpoint cho việc tạo người dùng mới
	router.POST("/users", func(c *gin.Context) {
		// Đọc dữ liệu từ yêu cầu người dùng
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(201, newUser)
	})
	// Khởi động server
	router.Run(":8080")
}
