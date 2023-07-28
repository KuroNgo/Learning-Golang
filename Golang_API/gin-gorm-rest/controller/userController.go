package controller

import (
	"Golang_API/gin-gorm-rest/models"
	"github.com/gin-gonic/gin"
)

func addUserData(ctx *gin.Context) {
	var newUser models.User

	// Có lỗi xảy ra và chúng ta cần xử lý nó
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
}
