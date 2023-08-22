package controller

import (
	"Golang_API/gin-gorm-rest/models"

	"github.com/gin-gonic/gin"
)

type controller struct {
	user models.User
}

func addUserData(ctx *gin.Context) {
	var newUser models.User

	// Có lỗi xảy ra và chúng ta cần xử lý nó
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
}

// Hiển thị thông tin người dùng
func (c *controller) FindAll() []models.User {
	// Thực hiện login để trả về danh sách user được tạo
	return c.user.FindAll()
}
