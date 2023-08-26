package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"server/conf"
	"server/model"
	"server/utils"
)

func SignUpUser(ctx *gin.Context) {
	var payload *model.RegisterUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	newUser := model.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  payload.Password,
		Role:      "user",
		Verified:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	result := conf.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "UNIQUE constraint failed: user.email") {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": model.FilteredResponse(&newUser)}})
}

func LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged out"})
}

func GoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")
	var pathUrl string = "/"

	if ctx.Query("state") != "" {
		pathUrl = ctx.Query("state")
	}

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tokenRes, err := utils.GetGoogleOAuthToken(code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	google_user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	email := strings.ToLower(google_user.Email)

	user_data := model.User{
		Name:      google_user.Name,
		Email:     email,
		Password:  "",
		Role:      "user",
		Provider:  "google",
		Photo:     google_user.Picture,
		Verified:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if conf.DB.Model(&user_data).Where("email = ?", email).Updates(&user_data).RowsAffected == 0 {
		conf.DB.Create(&user_data)
	}

	var user model.User
	conf.DB.First(&user, "email = ?", email)

	config, _ := conf.LoadConfig(".")

	token, err := utils.GenerateToken(config.TokenExpireIn, user.ID.String(), config.JWTTokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(config.FrontEndOrigin, pathUrl))
}

func SignInUser(ctx *gin.Context) {
	var payload *model.LoginUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	result := conf.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Provider == "Google" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	config, _ := conf.LoadConfig(".")

	token, err := utils.GenerateToken(config.TokenExpireIn, user.ID.String(), config.JWTTokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": model.FilteredResponse(&user)}})
}
