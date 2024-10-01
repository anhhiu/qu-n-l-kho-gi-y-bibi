package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func GetCustomerById(c *gin.Context) {
	var customer models.Customer
	if err := config.DB.Where("customer_id", c.Param("customer_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer bad request"})
		return
	}
	customer := models.Customer{
		CustomerName: input.CustomerName,
		Email:        input.Email,
		Phone:        input.Phone,
		Address:      input.Address,
		Password:     input.Password,
	}
	config.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func UpdateCustomerById(c *gin.Context) {
	var customer models.Customer
	if err := config.DB.Where("customer_id = ?", c.Param("customer_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer bad request"})
		return
	}

	config.DB.Model(&customer).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomerById(c *gin.Context) {
	var customer models.Customer
	if err := config.DB.Where("customer_id = ?", c.Param("customer_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	config.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
