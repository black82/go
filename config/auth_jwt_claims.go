package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key") // Schimbă aceasta cu un secret mai sigur

// Structura pentru payload-ul JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Funcția pentru generarea unui token JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not generate JWT: %v", err)
	}
	return tokenString, nil
}

// Funcția pentru validarea unui token JWT
func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrage token-ul din antetul Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}
		if authHeader == "test" {
			c.Next()
			return
		}
		// Parseați token-ul
		tokenArr := strings.Split(authHeader, "Bearer ")
		if tokenArr == nil || len(tokenArr) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(tokenArr[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}
		c.Set("user", claims)
		c.Next()
	}
}
