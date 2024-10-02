package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get all orders
// @Tags orders
// @Accept json
// @Produce json
// @Router /order/ [get]
func GetAllOrder(c *gin.Context) {
	var orders models.Order
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// @Summary create order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Orders data"
// @Router /order/ [post]
func CreateOrder(c *gin.Context) {
	// Khai báo cấu trúc đầu vào
	var input struct {
		CustomerID int       `json:"customer_id"`
		OrderDate  time.Time `json:"order_date"`
		Products   []struct {
			ProductID int     `json:"product_id"`
			Quantity  int     `json:"quantity"`
			UnitPrice float64 `json:"unit_price"`
		} `json:"products"`
	}

	// Kiểm tra lỗi JSON đầu vào
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Tạo đơn hàng mới
	order := models.Order{
		CustomerID: input.CustomerID,
		OrderDate:  input.OrderDate,
		Status:     "Pending", // Trạng thái mặc định
	}

	// Lưu đơn hàng vào cơ sở dữ liệu
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Tính toán tổng tiền
	var totalAmount float64
	for _, p := range input.Products {
		// Tính toán tổng tiền từ sản phẩm
		totalAmount += float64(p.Quantity) * p.UnitPrice

		// Tạo chi tiết đơn hàng
		orderDetail := models.OrderDetail{
			OrderID:   order.OrderID,
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			UnitPrice: p.UnitPrice,
		}

		// Lưu chi tiết đơn hàng vào cơ sở dữ liệu
		if err := config.DB.Create(&orderDetail).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
			return
		}
	}
	
	// Cập nhật tổng tiền cho đơn hàng
	order.TotalAmount = totalAmount
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order total amount"})
		return
	}

	// Trả về phản hồi
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}
// @Summary update order
	// @Tags orders
	// @Accept json
	// @Produce json
	//@Param order_id  path int true "OrderID"
	//@Param order body models.Order true "Orders data"
	// @Router /order/ [put]
func UpdateOrderById(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error not found"})
		return
	}
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error bad request"})
		return
	}

	config.DB.Model(&order).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": order})
}
//@tags orders
//@summary delele order by id
//@param orders_id path int true "OrderID"
//@router /order/{order_id} [delete]
func DeleteOrderById(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error not found"})
		return
	}

	config.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
//@tags orders
//@summary get order by id
//@param orders_id path int true "OrderID"
//@router /order/{order_id} [get]
func GetOrderById(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
