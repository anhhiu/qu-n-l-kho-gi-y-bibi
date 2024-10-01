package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSupplier(c *gin.Context) {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

func CreateSupplier(c *gin.Context) {
	var input models.Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	supplier := models.Supplier{
		SupplierName: input.SupplierName,
		Address:      input.Address,
		Phone:        input.Phone,
		Email:        input.Email,
		Website:      input.Website,
	}
	config.DB.Create(&supplier)
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

func GetSupplierById(c *gin.Context) {
	var supplier models.Supplier
	if err := config.DB.Where("supplier_id = ?", c.Param("supplier_id")).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

func UpdateSuplierById(c *gin.Context) {
	var supplier models.Supplier

	if err := config.DB.Where("supplier_id = ?", c.Param("supplier_id")).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	var input models.Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Supplier bad request"})
		return
	}
	config.DB.Model(&supplier).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

func DeleteSupplierById(c *gin.Context) {
	var supplier models.Supplier
	if err := config.DB.Where("supplier_id = ?", c.Param("supplier_id")).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	config.DB.Delete(&supplier)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
