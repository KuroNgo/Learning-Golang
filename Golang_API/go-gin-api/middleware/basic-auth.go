package middleware

import (
	"github.com/gin-gonic/gin"
)

// ĐIn nghĩa một middleware xác thực sử dụng Basic Authentication cho ứng dụng web sử dụng framework Gin
// Middleware này được sử dụng thông tin xác thực cụ thể
// Được cung cấp bởi gin.Account
func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"kuro": "12345",
	})
}
