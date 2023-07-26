package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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
	server.Static("/css", "/template/css")

	server.LoadHTMLGlob("templates/html/*.html")
	apiRoutes := server.Group("/api")
	{
		// để ghi log các yêu cầu HTTP và các thông tin liên quan
		// server.Use(gin.Logger())
		// Định nghĩa các route và xử lý request tại đây
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/postvideo", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!"})
			}
		})
	}
	// tự động xxuwr lý các panic lỗi runtime trong ứng dụng và
	// trả về một HTTP 500 Internal Server Error
	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.Auth(), gindump.Dump())
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	server.Run(":8080")
}
