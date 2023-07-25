package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"kuro.com/pragnaticreviews/golang-gin-poc/controller"
	"kuro.com/pragnaticreviews/golang-gin-poc/middleware"
	"kuro.com/pragnaticreviews/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	// tạo đối tượng *gin.Engine đã được cấu hình sẵn
	// server := gin.Default()
	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!!",
	// 	})
	// })
	// TODO: middleware
	// tạo mới một instance của framework Gin
	server := gin.New()

	// tự động xxuwr lý các panic lỗi runtime trong ứng dụng và
	// trả về một HTTP 500 Internal Server Error
	server.Use(gin.Recovery(), middleware.Logger())

	// để ghi log các yêu cầu HTTP và các thông tin liên quan
	// server.Use(gin.Logger())
	// Định nghĩa các route và xử lý request tại đây
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/postvideo", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
