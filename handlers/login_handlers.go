package handlers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var model models.AuthClaim
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var result bool
	result, err := services.Auth(model)
	if err != nil || false == result {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var token string
	token, err = config.GenerateJWT(model.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token ": token})
}
func Signup(c *gin.Context) {
	var model models.AuthClaim
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var result bool
	result, err := services.SignUp(model)
	msg := "User already exists"
	if err != nil {
		msg = err.Error()
	}
	if err != nil || result == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
