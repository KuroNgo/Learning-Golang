package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kuro.com/pragnaticreviews/golang-gin-poc/service"
	"log"
	"net/http"
)

func AuthorizedJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]:", claims["name"])
			log.Println("Claims[Admin]", claims["admin"])
			log.Println("Claims[Issue]", claims["iss"])
			log.Println("Claims[IssueAt]", claims["iat"])
			log.Println("Claims[ExpireAt]", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}
