package controllers

import (
	"bibi/config"
	"bibi/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	//"google.golang.org/protobuf/internal/order"
	"gorm.io/gorm"
)

// @tags orderdetails
// @summary Get all orderdetails
// @router /orderdetail/ [get]
func GetALLOrderDetail(c *gin.Context) {
	var orderdetails []models.OrderDetail
	if err := config.DB.Preload("Order").Preload("Product").Find(&orderdetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order details"})
		return
	}
	if len(orderdetails) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No order details found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderdetails})
}

// @Summary Get order_detail by ID
// @Description Retrieve a specific order detail by ID
// @Tags orderdetails
// @Accept json
// @Produce json
// @Param order_detail_id path int true "OrderDetailID"
// @Router /orderdetail/{order_detail_id} [get]
func GetOrderDetailById(c *gin.Context) {
	var orderdetail models.OrderDetail
	if err := config.DB.Preload("Order").Preload("Product").Where("order_detail_id = ?", c.Param("order_detail_id")).First(&orderdetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order detail not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order detail"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderdetail})
}
