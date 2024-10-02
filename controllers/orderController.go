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

// @Summary Create Order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body CreateOrderInput true "Order data"
// @Router /order/ [post]
func CreateOrder(c *gin.Context) {
	// Khai báo cấu trúc đầu vào
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Tạo đơn hàng mới
	order := models.Order{
		CustomerID: input.CustomerID,
		OrderDate:  time.Now(), // Bạn có thể lấy ngày giờ hiện tại
		Status:     "Pending",  // Trạng thái mặc định
	}

	// Lưu đơn hàng vào cơ sở dữ liệu
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Tính toán tổng tiền và lưu chi tiết đơn hàng
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

// CreateOrderInput là cấu trúc dùng để nhận dữ liệu đầu vào
type CreateOrderInput struct {
	CustomerID int `json:"customer_id"`
	Products   []struct {
		ProductID int     `json:"product_id"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
	} `json:"products"`
}

// @Summary Update Order by ID
// @Tags orders
// @Param order_id path int true "Order ID"
// @Param order body UpdateOrderInput true "Order data"
// @Router /order/{order_id} [put]
func UpdateOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")
	var order models.Order

	// Kiểm tra xem hóa đơn có tồn tại hay không
	if err := config.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cập nhật thông tin hóa đơn
	order.CustomerID = input.CustomerID
	order.TotalAmount = input.TotalAmount
	order.Status = input.Status

	// Lưu hóa đơn đã sửa đổi vào cơ sở dữ liệu
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// UpdateOrderInput là cấu trúc dùng để nhận dữ liệu đầu vào cho việc cập nhật
type UpdateOrderInput struct {
	CustomerID  int     `json:"customer_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}

// @Summary Delete Order by ID
// @Tags orders
// @Param order_id path int true "Order ID"
// @Router /order/{order_id} [delete]
func DeleteOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")
	var order models.Order

	// Kiểm tra xem hóa đơn có tồn tại hay không
	if err := config.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Xóa hóa đơn khỏi cơ sở dữ liệu
	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// @tags orders
// @summary delele order by id
// @param orders_id path int true "OrderID"
// @router /order/{order_id} [delete]
func GetOrderById(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
