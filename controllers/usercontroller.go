package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	user := models.Users{
		Role:     input.Role,
		UserName: input.UserName,
		PassWord: input.PassWord,
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	var users []models.Users
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
