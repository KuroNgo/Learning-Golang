package main

import (
	"Golang_API/gin-gorm-rest/models"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connectionInfo := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"

	// Mở kết nối đến CSDL
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	// Đảm bảo bảng User đã được tạo trong cơ sở dữ liệu
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		// Nếu có lỗi xảy ra, dừng chương trình
		log.Fatal("Could not migrate the database: ", err)
	}

	// Khởi tạo router Gin
	router := gin.Default()

	// ĐỊnh nghĩa endpoint cho việc tạo người dùng mới
	router.POST("/users", func(c *gin.Context) {
		// Đọc dữ liệu từ yêu cầu người dùng
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}
		result := db.Create(&newUser)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Unable to create user"})
			return
		}
		c.JSON(201, newUser)
	})

	router.GET("/users", func(c *gin.Context) {
		var user []models.User

		// Truy vấn tất cả người dùng từ cơ sở dữ liệu
		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Unable to find users"})
			return
		}

		// Trả về danh sách người dùng
		c.JSON(200, user)
	})
	// Khởi động server
	router.Run(":8080")
}
