package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSupplier godoc
// @Summary Get all suppliers
// @Tags suppliers
// @Accept json
// @Produce json
// @Router /supplier/ [get]
func GetSupplier(c *gin.Context) {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

// @tags suppliers
// @Summary Create supplier
// @Param supplier body models.Supplier true "Supplier data"
// @Router /supplier/ [post]
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

// @tags suppliers
// @Summary Get supplier by id
// @Param supplier_id path int true "Supplier ID"
// @Router /supplier/{supplier_id} [get]
func GetSupplierById(c *gin.Context) {
	var supplier models.Supplier
	if err := config.DB.Where("supplier_id = ?", c.Param("supplier_id")).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

// @tags suppliers
// @Summary Update supplier
// @Param supplier_id path int true "Supplier ID"
// @Param supplier body models.Supplier true "Supplier info"
// @Router /supplier/{supplier_id} [put]
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

// @tags suppliers
// @Summary Delete supplier
// @Param supplier_id path int true "Supplier ID"
// @Router /supplier/{supplier_id} [delete]
func DeleteSupplierById(c *gin.Context) {
	var supplier models.Supplier
	if err := config.DB.Where("supplier_id = ?", c.Param("supplier_id")).First(&supplier).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	config.DB.Delete(&supplier)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
