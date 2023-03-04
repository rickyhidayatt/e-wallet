package middleware

import (
	"e-wallet/config"
	"e-wallet/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		token, err := jwt.ParseWithClaims(cookie, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWT_KEY), nil
		})
		if err != nil {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*config.JWTClaim); ok && token.Valid {
			c.Set("userId", claims.Email)
			c.Next()
		} else {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", "Invalid Token")
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
	}
}
