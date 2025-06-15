package controllers

import (
	"net/http"

	"github.com/HarshalSankanna/doctor-portal/doctor-receptionist-portal/config"
	"github.com/HarshalSankanna/doctor-portal/doctor-receptionist-portal/models"
	"github.com/HarshalSankanna/doctor-portal/doctor-receptionist-portal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input models.User
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	token, _ := utils.GenerateJWT(user.Email, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
