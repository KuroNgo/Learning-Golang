package controller

import (
	"github.com/gin-gonic/gin"
	"kuro.com/pragnaticreviews/golang-gin-poc/entity"
	"kuro.com/pragnaticreviews/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	// Các thuộc tính của cấu trúc controller
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

// Phương thức FindAll với receiver là con trỏ (*controller)
func (c *controller) FindAll() []entity.Video {
	// THực hiện logic để lấy ra danh sách video
	// và trả về danhh sách đối tượng entity
	return c.service.FindAll()
}

// Phương thức Save với receiver là con trỏ (*controller)
func (c *controller) Save(ctx *gin.Context) entity.Video {

	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
