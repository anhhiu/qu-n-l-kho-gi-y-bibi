package controllers

import (
	"bibi/config"
	"bibi/models"
	"fmt"

	//"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
// Get
// @tags auth
// @summary Get all users
// @router /user/ [get]
func GetUsers(c *gin.Context) {
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
		fmt.Println("Error retrieving users:", err) // Log lỗi
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// sủ dungjk header để phân quyền
func PhanQuyen(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Request.Header.Get("Role")
		authorized := false

		for _, v := range roles {
			if userRole == v {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "không có quyền truy cập ông cháu ơi"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthorizeRole(db *gorm.DB, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.Header.Get("Username")
		password := c.Request.Header.Get("Password")

		// Log để kiểm tra thông tin xác thực
		fmt.Printf("Username: %s, Password: %s\n", username, password)

		var user models.Users
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, sai tk"})
			c.Abort()
			return
		}

		// Kiểm tra mật khẩu
		/* if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, sai mk"})
			c.Abort()
			return
		} */
		if err := db.Where("password = ?", password).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, sai mk"})
			c.Abort()
			return
		}

		// Kiểm tra vai trò
		userRole := user.Role
		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()  
	}
}
