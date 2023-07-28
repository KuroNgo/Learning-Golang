package models

import "gorm.io/gorm"

type User struct {
	// gorm.Model sẽ tự động tạo các trươn ID, CreateAt, UpdateAt, DeleteAt trong bảng của mình
	// Trong trường hợp này, ID sẽ được PostgreSQL tự động tăng khi thêm dữ liệu mới
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
