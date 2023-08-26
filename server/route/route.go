package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"server/conf"
	"server/handler"
	"server/middleware"
)

type Service struct {
	*conf.App
}

type IRoute interface {
	NewService() *Service
}

var server *gin.Engine

func NewService() *Service {
	s := &Service{
		conf.NewApp(),
	}
	corsConfig := cors.DefaultConfig()

	// AllowOrigins là danh sách các nguồn gốc mà yêu cầu tên miền chéo có thể được thực thi từ đó.
	// Nếu giá trị "*" đặc biệt có trong danh sách, tất cả nguồn gốc sẽ được cho phép.
	// Giá trị mặc định là []
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}

	// AllowOriginFunc là một hàm tùy chỉnh để xác thực nguồn gốc. Nó lấy nguồn gốc
	// làm đối số và trả về true nếu được phép hoặc sai nếu không. Nếu tùy chọn này là
	// được đặt, nội dung của AllowOrigins sẽ bị bỏ qua.

	// AllowMethods là danh sách các phương thức client được phép sử dụng
	// yêu cầu tên miền chéo. Giá trị mặc định là các phương thức đơn giản (GET, POST, PUT, PATCH, DELETE, HEAD và OPTIONS)
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// AllowHeaders là danh sách các tiêu đề không đơn giản mà khách hàng được phép sử dụng với
	// yêu cầu tên miền chéo.
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	// AllowCredentials cho biết liệu yêu cầu có thể bao gồm thông tin xác thực của người dùng như
	// cookie, xác thực HTTP hoặc chứng chỉ SSL phía máy khách.
	corsConfig.AllowCredentials = true

	corsConfig.ExposeHeaders = []string{"Content-Length", "Content-Type"}

	// MaxAge cho biết kết quả của yêu cầu preflight trong bao lâu (với độ chính xác thứ hai)
	// có thể được lưu vào bộ nhớ đệm
	corsConfig.MaxAge = 86400

	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	corsConfig.AllowWildcard = true

	// Cho phép sử dụng các lược đồ tiện ích mở rộng trình duyệt phổ biến
	corsConfig.AllowBrowserExtensions = true
	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success", "message": "Implement Google OAuth2 in GoLang"})
	})

	auth_router := router.Group("/auth")
	auth_router.POST("/google", handler.SignUpUser)
	auth_router.POST("/login", handler.SignInUser)
	auth_router.GET("/logout", middleware.DeserializerUser())

	router.GET("/sessions/oauth/google", handler.GoogleOAuth)
	router.GET("/users/me", middleware.DeserializerUser(), handler.GetMe)

	router.StaticFS("/images", http.Dir("./images"))
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route not found"})
	})
	return s
}
