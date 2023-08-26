package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/model"
)

func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.User)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": model.FilteredResponse(&currentUser),
		}})
}
