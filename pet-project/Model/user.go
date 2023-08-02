package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Tạo ID tự tăng "identity"
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Hoten      string    `json:"name"`
	NgaySinh   time.Time `json:"ngaysinh"`
	Email      string    `json:"email"`
	SDT        string    `json:"sdt"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
