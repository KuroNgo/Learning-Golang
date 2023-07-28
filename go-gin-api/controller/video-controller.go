package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"kuro.com/pragnaticreviews/golang-gin-poc/entity"
	"kuro.com/pragnaticreviews/golang-gin-poc/service"
	"kuro.com/pragnaticreviews/golang-gin-poc/validators"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	//TODO Các thuộc tính của cấu trúc controller
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

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
// ĐƯợc sử dụng để xử lý yêu cầu POST và lưu trữ
// một đối tượng video mới được gửi bởi người dùng
func (c *controller) Save(ctx *gin.Context) error {

	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
