package config

import (
	"Golang_API/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	// Thay đổi thông tin kết nối của bạn tại đây
	connectionInfo := "host=localhost port=5432 user=kuro dbname=gin-gorm-rest sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	// Đảm bảo rằng bảng User đã được tạo trong cơ sở dữ liệu
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Could not create table:", err)
	}

}
