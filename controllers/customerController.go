package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomer godoc
// @Summary Get all customers
// @Description Retrieve a list of customers
// @Tags customers
// @Accept json
// @Produce json
// @Success 200 {array} models.Customer
// @Router /customer/ [get]
func GetCustomer(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// @tags customers
// @Summary Get customer by id
// @Param customer_id path int true "Customer ID"
// @Router /customer/{customer_id} [get]
func GetCustomerById(c *gin.Context) {
	var customer models.Customer
	if err := config.DB.Where("customer_id", c.Param("customer_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @tags customers
// @Summary Create customer
// @Param customer body models.Customer true "customer data"
// @Router /customer/ [post]
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

// @tags customers
// @Summary Update customer by id
// @Param customer_id path int true "Customer ID"
// @Param customer body models.Customer true "Customer info"
// @Router /customer/{customer_id} [put]
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

// @tags customers
// @Summary Delete customer by id
// @Param customer_id path int true "Customer ID"
// @Router /customer/{customer_id} [delete]
func DeleteCustomerById(c *gin.Context) {
	var customer models.Customer
	if err := config.DB.Where("customer_id = ?", c.Param("customer_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	config.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
